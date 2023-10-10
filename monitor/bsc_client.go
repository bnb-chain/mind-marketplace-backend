package monitor

import (
	"context"
	"fmt"
	"github.com/avast/retry-go/v4"
	ethcommon "github.com/ethereum/go-ethereum/common"

	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/bnb-chain/greenfield-data-marketplace-backend/util"
)

type BscClient struct {
	rpcClient *rpc.Client // for apis eth_getFinalizedBlock and eth_getFinalizedHeader usage, supported by BSC
	ethClient *ethclient.Client
	provider  string
	height    uint64
	updatedAt time.Time
}

func newBscClients(rpcAddrs []string) []*BscClient {
	bscClients := make([]*BscClient, 0)
	for _, provider := range rpcAddrs {
		rpcClient, err := rpc.DialContext(context.Background(), provider)
		if err != nil {
			panic("new rpc client error")
		}
		ethClient, err := ethclient.Dial(provider)
		if err != nil {
			panic("new eth client error")
		}

		bscClients = append(bscClients, &BscClient{
			rpcClient: rpcClient,
			ethClient: ethClient,
			provider:  provider,
			updatedAt: time.Now(),
		})
	}
	return bscClients
}

type BscCompositeClients struct {
	mutex                     sync.RWMutex
	clientIdx                 int
	bscClients                []*BscClient
	numberOfBlocksForFinality int
}

func NewBscCompositeClients(rpcAddrs []string, numberOfBlocksForFinality int) *BscCompositeClients {
	return &BscCompositeClients{
		clientIdx:                 0,
		bscClients:                newBscClients(rpcAddrs),
		numberOfBlocksForFinality: numberOfBlocksForFinality,
	}
}

func (e *BscCompositeClients) GetRpcClient() *rpc.Client {
	e.mutex.RLock()
	defer e.mutex.RUnlock()
	return e.bscClients[e.clientIdx].rpcClient
}

func (e *BscCompositeClients) GetEthClient() *ethclient.Client {
	e.mutex.RLock()
	defer e.mutex.RUnlock()
	return e.bscClients[e.clientIdx].ethClient
}

func (e *BscCompositeClients) SwitchClient() {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	e.clientIdx++
	if e.clientIdx >= len(e.bscClients) {
		e.clientIdx = 0
	}
	util.Logger.Infof("switch to provider: %s", e.bscClients[e.clientIdx].provider)
}

func (e *BscCompositeClients) GetLatestFinalizedBlockHeightWithRetry() (latestHeight uint64, err error) {
	return e.getLatestBlockHeightWithRetry(e.GetEthClient(), e.GetRpcClient(), true)
}

func (e *BscCompositeClients) GetLatestBlockHeightWithRetry() (latestHeight uint64, err error) {
	return e.getLatestBlockHeightWithRetry(e.GetEthClient(), e.GetRpcClient(), false)
}

func (e *BscCompositeClients) getLatestBlockHeightWithRetry(ethClient *ethclient.Client, rpcClient *rpc.Client, finalized bool) (latestHeight uint64, err error) {
	return latestHeight, retry.Do(func() error {
		latestHeight, err = e.getLatestBlockHeight(ethClient, rpcClient, finalized)
		return err
	}, RtyAttem,
		RtyDelay,
		RtyErr,
		retry.OnRetry(func(n uint, err error) {
			util.Logger.Errorf("failed to query latest height, attempt: %d times, max_attempts: %d", n+1, RtyAttNum)
		}))
}

func (e *BscCompositeClients) getLatestBlockHeight(client *ethclient.Client, rpcClient *rpc.Client, finalized bool) (uint64, error) {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), RPCTimeout)
	defer cancel()
	if finalized {
		return e.getFinalizedBlockHeight(ctxWithTimeout, rpcClient)
	}
	header, err := client.HeaderByNumber(ctxWithTimeout, nil)
	if err != nil {
		return 0, err
	}
	return header.Number.Uint64(), nil
}

func (e *BscCompositeClients) UpdateClientLoop() {
	ticker := time.NewTicker(SleepSecondForUpdateClient)
	for range ticker.C {
		util.Logger.Infof("start to monitor bsc data-seeds healthy")
		for _, bscClient := range e.bscClients {
			if time.Since(bscClient.updatedAt).Seconds() > DataSeedDenyServiceThreshold {
				util.Logger.Error(fmt.Sprintf("data seed %s is not accessable", bscClient.provider))
			}
			height, err := e.getLatestBlockHeight(bscClient.ethClient, bscClient.rpcClient, true)
			if err != nil {
				util.Logger.Errorf("get latest block height error, err=%s", err.Error())
				continue
			}
			bscClient.height = height
			bscClient.updatedAt = time.Now()
		}

		highestHeight := uint64(0)
		highestIdx := 0
		for idx := 0; idx < len(e.bscClients); idx++ {
			if e.bscClients[idx].height > highestHeight {
				highestHeight = e.bscClients[idx].height
				highestIdx = idx
			}
		}
		// current client block sync is fall behind, switch to the client with the highest block height
		if e.bscClients[e.clientIdx].height+FallBehindThreshold < highestHeight {
			e.mutex.Lock()
			e.clientIdx = highestIdx
			e.mutex.Unlock()
		}
	}
}

func (e *BscCompositeClients) GetBlockHeader(height uint64) (*types.Header, error) {
	ctx, cancel := context.WithTimeout(context.Background(), RPCTimeout)
	defer cancel()
	header, err := e.GetEthClient().HeaderByNumber(ctx, big.NewInt(int64(height)))
	if err != nil {
		return nil, err
	}
	return header, nil
}

// getFinalizedBlockHeight gets the finalizedBlockHeight, which is the larger one between (fastFinalizedBlockHeight, BscBlocksForFinality from config).
func (e *BscCompositeClients) getFinalizedBlockHeight(ctx context.Context, rpcClient *rpc.Client) (uint64, error) {
	var head *types.Header
	if err := rpcClient.CallContext(ctx, &head, "eth_getFinalizedHeader", e.numberOfBlocksForFinality); err != nil {
		return 0, err
	}
	if head == nil || head.Number == nil {
		return 0, ethereum.NotFound
	}
	return head.Number.Uint64(), nil
}

func (e *BscCompositeClients) QueryChainLogs(fromBlock, toBlock uint64, topicHashes []ethcommon.Hash, contract ethcommon.Address) ([]types.Log, error) {
	client := e.GetEthClient()
	topics := [][]ethcommon.Hash{topicHashes}
	logs, err := client.FilterLogs(context.Background(), ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(fromBlock)),
		ToBlock:   big.NewInt(int64(toBlock)),
		Topics:    topics,
		Addresses: []ethcommon.Address{contract},
	})
	if err != nil {
		return nil, err
	}
	return logs, nil
}
