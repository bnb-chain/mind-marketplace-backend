package monitor

import (
	"encoding/json"
	"github.com/shopspring/decimal"
)

type Extra struct {
	Desc  string          `json:"desc"`
	Url   string          `json:"url"`
	Price decimal.Decimal `json:"price"`
}

func parseExtra(str string) (*Extra, error) {
	var extra Extra
	err := json.Unmarshal([]byte(str), &extra)
	if err != nil {
		return nil, err
	}

	return &extra, nil
}
