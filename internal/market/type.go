package market

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type OverUnderMarket struct {
	ID        string    `json:"id"`
	EventID   uint64    `json:"eventId"`
	Name      string    `json:"name"`
	Side      string    `json:"side"`
	Exchange  string    `json:"exchange"`
	Over      PriceSize `json:"over"`
	Under     PriceSize `json:"under"`
	Timestamp int64     `json:"timestamp"`
}

func (m OverUnderMarket) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func (m *OverUnderMarket) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &m)
}

type PriceSize struct {
	Price  float32     `json:"price"`
	Size   float32     `json:"size"`
}
