package datatype

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type JSON json.RawMessage

var (
	errScan = errors.New("failed to unmarshal JSON")

	errInvalidType = errors.New("arg invalid datatype")
)

// Scan scan value into Jsonb, implements sql.Scanner interface.
func (j *JSON) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("%w %s", errScan, value)
	}

	result := json.RawMessage{}
	err := json.Unmarshal(bytes, &result)
	*j = JSON(result)
	return err
}

// Value return json value, implement driver.Valuer interface.
func (j JSON) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return json.RawMessage(j).MarshalJSON()
}

func (j JSON) MarshalJSON() ([]byte, error) {
	return json.RawMessage(j).MarshalJSON()
}

func (j *JSON) UnmarshalJSON(data []byte) error {
	result := json.RawMessage{}
	err := json.Unmarshal(data, &result)
	*j = JSON(result)
	return err
}

func (j JSON) ToMap() (map[string]interface{}, error) {
	b, err := j.MarshalJSON()
	if err != nil {
		return nil, err
	}
	m := make(map[string]interface{})
	err = json.Unmarshal(b, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (j JSON) ToStringMap() (map[string]string, error) {
	s := make(map[string]string)
	m, err := j.ToMap()
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		r, ok := v.(string)
		if !ok {
			return nil, fmt.Errorf("%v: %w", v, errInvalidType)
		}
		s[k] = r
	}
	return s, nil
}
