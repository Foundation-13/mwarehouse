package types

import (
	"encoding/json"
	"fmt"
	"io"
)

const (
	FilterTypeMonochrome = "monochrome"
	FilterTypeScale      = "scale"
)

var AllFilters = []string{
	FilterTypeMonochrome,
	FilterTypeScale,
}

type Filter interface {
	GetName() string
	GetApplied() bool
	GetExecutionTime() int64
}

// NamedFilter
type NamedFilter struct {
	Name          string `json:"name"`
	Applied       bool   `json:"applied"`
	ExecutionTime int64  `json:"execution_time"`
}

func (f NamedFilter) GetName() string {
	return f.Name
}

func (f NamedFilter) GetApplied() bool {
	return f.Applied
}

func (f NamedFilter) GetExecutionTime() int64 {
	return f.ExecutionTime
}

// FilterWithPercentage
type FilterWithPercentage struct {
	NamedFilter
	Percentage int `json:"percentage"`
}

// Filters
type Filters struct {
	Filters    []Filter          `json:"-"`
	RawFilters []json.RawMessage `json:"filters"`
}

// Marshal/Unmarshal
func UnmarshalFilters(r io.Reader) (Filters, error) {
	var filters Filters

	err := json.NewDecoder(r).Decode(&filters)
	if err != nil {
		return Filters{}, fmt.Errorf("failed to unmarshal filters, %w", err)
	}

	for _, raw := range filters.RawFilters {
		var v NamedFilter
		err = json.Unmarshal(raw, &v)
		if err != nil {
			return Filters{}, err
		}

		var f Filter
		switch v.Name {
		case FilterTypeMonochrome:
			f = v
		case FilterTypeScale:
			var fp FilterWithPercentage
			err = json.Unmarshal(raw, &fp)
			if err != nil {
				return Filters{}, fmt.Errorf("failed to unmarshal filter with percentage, %w", err)
			}
			f = fp
		default:
			return Filters{}, fmt.Errorf("unknown filter type: %s", v.Name)
		}

		filters.Filters = append(filters.Filters, f)
	}

	return filters, nil
}
