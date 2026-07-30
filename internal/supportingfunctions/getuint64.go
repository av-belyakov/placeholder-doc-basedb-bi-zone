package supportingfunctions

import "fmt"

func GetUint64(a any) (uint64, error) {
	switch v := a.(type) {
	case int:
		return uint64(v), nil

	case int8:
		return uint64(v), nil

	case int16:
		return uint64(v), nil

	case int32:
		return uint64(v), nil

	case int64:
		return uint64(v), nil

	case float32:
		return uint64(v), nil

	case float64:
		return uint64(v), nil

	case uint8:
		return uint64(v), nil

	case uint16:
		return uint64(v), nil

	case uint32:
		return uint64(v), nil

	case uint64:
		return v, nil

	}

	return 0, fmt.Errorf("the '%v' value is not a number", a)
}
