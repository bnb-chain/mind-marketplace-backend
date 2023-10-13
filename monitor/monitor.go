package monitor

import (
	"github.com/bnb-chain/greenfield-data-marketplace-backend/util"
	"gorm.io/gorm"
	"time"
)

type Monitor struct {
	processor   BlockProcessor
	startHeight uint64
}

func NewMonitor(p BlockProcessor, startHeight uint64) *Monitor {
	return &Monitor{processor: p, startHeight: startHeight}
}

func (m *Monitor) Start() {
	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		err := m.run()
		if err != nil {
			util.Logger.Errorf("processor: %s, fail to run with error: %s", m.processor.Name(), err)
		}
	}
}

func (m *Monitor) run() error {
	blockchainHeight, err := m.processor.GetBlockchainBlockHeight()
	util.Logger.Infof("processor: %s, current blockchain height: %d", m.processor.Name(), blockchainHeight)
	if err != nil {
		return err
	}

	dbHeight, err := m.processor.GetDatabaseBlockHeight()
	util.Logger.Infof("processor: %s, current database height: %d", m.processor.Name(), dbHeight)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if dbHeight < m.startHeight {
		dbHeight = m.startHeight - 1 // to include start height
	}

	for dbHeight < blockchainHeight {
		util.Logger.Infof("processor: %s, processing height: %d", m.processor.Name(), dbHeight+1)
		err = m.processor.Process(dbHeight + 1)
		if err != nil {
			return err
		}
		dbHeight++
		time.Sleep(50 * time.Millisecond)
	}

	return nil
}
