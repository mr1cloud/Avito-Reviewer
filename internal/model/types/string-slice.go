package types

import (
	"encoding/json"
	"fmt"
)

type StringSlice []string

func (s *StringSlice) Scan(src interface{}) error {
	if src == nil {
		*s = make(StringSlice, 0)
		return nil
	}

	switch v := src.(type) {
	case []byte:
		return json.Unmarshal(v, s)
	case string:
		return json.Unmarshal([]byte(v), s)
	default:
		return fmt.Errorf("cannot scan %T into StringSlice", src)
	}
}
