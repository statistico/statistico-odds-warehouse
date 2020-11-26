package market

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Market struct {
	ID              string            `json:"id"`
	EventID         uint64            `json:"eventId"`
	Name            string            `json:"name"`
	Side            string            `json:"side"`
	Exchange        string            `json:"exchange"`
	Runners         []*Runner         `json:"runners"`
	Timestamp       int64             `json:"timestamp"`
}

func (m Market) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func (m *Market) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &m)
}

type Runner struct {
	ID     uint64      `json:"id"`
	Name   string      `json:"name"`
	Prices []PriceSize `json:"prices"`
}

type PriceSize struct {
	Price float32 `json:"price"`
	Size  float32 `json:"size"`
}
