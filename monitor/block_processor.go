package monitor

type BlockProcessor interface {
	Name() string                              // processor name
	GetDatabaseBlockHeight() (uint64, error)   // get database max block height
	GetBlockchainBlockHeight() (uint64, error) // get blockchain max finalized block height
	Process(blockHeight uint64) error          // process a block
}
