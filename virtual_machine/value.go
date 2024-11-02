package virtual_machine

import (
	"fmt"
	"math"
	"strconv"
)

const (
	Integer = iota
	Float   = iota
)

type Value struct {
	Type  int
	Value string
}

func (v *Value) String() string {
	return v.Value
}

func (v *Value) Add(rv *Value) (*Value, error) {
	result := v.buildResult(rv)

	if result.Type == Float {
		lValue, rValue, err := v.parseFloats(rv)

		if err != nil {
			return nil, err
		}

		result.Value = fmt.Sprintf("%f", lValue+rValue)
	} else {
		lValue, rValue, err := v.parseInt(rv)

		if err != nil {
			return nil, err
		}

		result.Value = fmt.Sprintf("%d", lValue+rValue)
	}

	return result, nil
}

func (v *Value) Subtraction(rv *Value) (*Value, error) {
	result := v.buildResult(rv)

	if result.Type == Float {
		lValue, rValue, err := v.parseFloats(rv)

		if err != nil {
			return nil, err
		}

		result.Value = fmt.Sprintf("%f", lValue-rValue)
	} else {
		lValue, rValue, err := v.parseInt(rv)

		if err != nil {
			return nil, err
		}

		result.Value = fmt.Sprintf("%d", lValue-rValue)
	}

	return result, nil
}

func (v *Value) Multiplication(rv *Value) (*Value, error) {
	result := v.buildResult(rv)

	if result.Type == Float {
		lValue, rValue, err := v.parseFloats(rv)

		if err != nil {
			return nil, err
		}

		result.Value = fmt.Sprintf("%f", lValue*rValue)
	} else {
		lValue, rValue, err := v.parseInt(rv)

		if err != nil {
			return nil, err
		}

		result.Value = fmt.Sprintf("%d", lValue*rValue)
	}

	return result, nil
}

func (v *Value) Divide(rv *Value) (*Value, error) {
	result := v.buildResult(rv)

	if result.Type == Float {
		lValue, rValue, err := v.parseFloats(rv)

		if err != nil {
			return nil, err
		}

		result.Value = fmt.Sprintf("%f", lValue/rValue)
	} else {
		lValue, rValue, err := v.parseInt(rv)

		if err != nil {
			return nil, err
		}

		result.Value = fmt.Sprintf("%d", lValue/rValue)
	}

	return result, nil
}

func (v *Value) Power(rv *Value) (*Value, error) {
	result := v.buildResult(rv)

	if result.Type == Float {
		lValue, rValue, err := v.parseFloats(rv)

		if err != nil {
			return nil, err
		}

		result.Value = fmt.Sprintf("%f", math.Pow(lValue, rValue))
	} else {
		lValue, rValue, err := v.parseInt(rv)

		if err != nil {
			return nil, err
		}

		result.Value = fmt.Sprintf("%d", int(math.Pow(float64(lValue), float64(rValue))))
	}

	return result, nil
}

func (v *Value) buildResult(rv *Value) *Value {
	newType := Integer

	if v.Type == Float || rv.Type == Float {
		newType = Float
	}

	return &Value{
		Type:  newType,
		Value: "0",
	}
}

func (v *Value) parseFloats(rv *Value) (float64, float64, error) {
	lValue, err := strconv.ParseFloat(v.Value, 32)

	if err != nil {
		return 0, 0, err
	}

	rValue, err := strconv.ParseFloat(rv.Value, 32)

	if err != nil {
		return 0, 0, err
	}

	return lValue, rValue, nil
}

func (v *Value) parseInt(rv *Value) (int64, int64, error) {
	lValue, err := strconv.ParseInt(v.Value, 10, 64)

	if err != nil {
		return 0, 0, err
	}

	rValue, err := strconv.ParseInt(rv.Value, 10, 64)

	if err != nil {
		return 0, 0, err
	}

	return lValue, rValue, nil
}
