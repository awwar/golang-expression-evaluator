package virtual_machine

import (
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
	return v.calculate(rv, func(lValue float64, rValue float64) float64 { return lValue + rValue })
}

func (v *Value) Subtract(rv *Value) (*Value, error) {
	return v.calculate(rv, func(lValue float64, rValue float64) float64 { return lValue - rValue })
}

func (v *Value) Multiply(rv *Value) (*Value, error) {
	return v.calculate(rv, func(lValue float64, rValue float64) float64 { return lValue * rValue })
}

func (v *Value) Divide(rv *Value) (*Value, error) {
	return v.calculate(rv, func(lValue float64, rValue float64) float64 { return lValue / rValue })
}

func (v *Value) Power(rv *Value) (*Value, error) {
	return v.calculate(rv, func(lValue float64, rValue float64) float64 { return math.Pow(lValue, rValue) })
}

func (v *Value) calculate(rv *Value, callback func(float64, float64) float64) (*Value, error) {
	result := v.buildResult(rv)

	lValue, rValue, err := v.parseFloats(rv)

	if err != nil {
		return nil, err
	}

	value := callback(lValue, rValue)

	if result.Type == Float {
		result.Value = strconv.FormatFloat(value, 'f', 6, 64)
	} else {
		result.Value = strconv.Itoa(int(value))
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
	lValue, err := strconv.ParseFloat(v.Value, 64)

	if err != nil {
		return 0, 0, err
	}

	rValue, err := strconv.ParseFloat(rv.Value, 64)

	if err != nil {
		return 0, 0, err
	}

	return lValue, rValue, nil
}
