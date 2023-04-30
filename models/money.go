package models

import (
	"database/sql/driver"
	"fmt"
	"strings"
)

type Money float64

func (m Money) Value() (driver.Value, error) {
	return fmt.Sprintf("%.2f", m), nil
}

func (m *Money) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	switch value.(type) {
	case float64:
		*m = Money(value.(float64))
	case string:
		strValue := strings.Replace(value.(string), ",", ".", -1)
		floatValue, err := fmt.Sscan(strValue, m)
		if err != nil {
			return err
		}
		if floatValue != 1 {
			return fmt.Errorf("expected to scan 1 value, got %d", floatValue)
		}
	default:
		return fmt.Errorf("cannot convert %v to Money", value)
	}

	return nil
}
