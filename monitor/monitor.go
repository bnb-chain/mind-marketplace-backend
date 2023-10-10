package monitor

import (
	"github.com/bnb-chain/greenfield-data-marketplace-backend/util"
	"gorm.io/gorm"
	"time"
)

type Monitor struct {
	processor BlockProcessor
}

func NewMonitor(p BlockProcessor) *Monitor {
	return &Monitor{processor: p}
}

func (m *Monitor) Start() {
	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		err := m.run()
		if err != nil {
			util.Logger.Errorf("fail to run with error: %s", err)
		}
	}
}

func (m *Monitor) run() error {
	blockchainHeight, err := m.processor.GetBlockchainBlockHeight()
	if err != nil {
		return err
	}
	dbHeight, err := m.processor.GetDatabaseBlockHeight()
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	for dbHeight < blockchainHeight {
		err = m.processor.Process(dbHeight + 1)
		if err != nil {
			return err
		}
		dbHeight++
	}

	return nil
}
