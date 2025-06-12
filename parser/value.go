package parser

import (
	"fmt"
	"math"
	"strconv"
)

type NodeValueType int

const (
	Integer NodeValueType = iota
	Float                 = iota
	Atom                  = iota
	String                = iota
)

var MapTypeToTypeName = map[NodeValueType]string{Integer: "int", Float: "float", Atom: "atom", String: "string"}

type Value struct {
	Type      NodeValueType
	StringVal *string
	FloatVal  *float64
	IntVal    *int64
}

// FOR REPR PURPOSE ONLY!!!
func (v *Value) String() string {
	asStringValue, _ := v.ToString()

	if v.Type == String {
		return fmt.Sprintf(`"%s"`, *asStringValue.StringVal)
	}

	return *asStringValue.StringVal
}

func (v *Value) GoString() string {
	return v.String()
}

func (v *Value) TypeAsString() string {
	return MapTypeToTypeName[v.Type]
}

func (v *Value) Add(rv *Value) (*Value, error) {
	if v.Type == String || rv.Type == String {
		leftValue, err := v.ToString()
		if err != nil {
			return nil, err
		}

		rightValue, err := rv.ToString()
		if err != nil {
			return nil, err
		}

		stVal := fmt.Sprintf("%s%s", *leftValue.StringVal, *rightValue.StringVal)

		return &Value{StringVal: &stVal, Type: String}, nil
	}

	if !v.IsNumber() || !rv.IsNumber() {
		return nil, fmt.Errorf("unable to add %s to %s", v.TypeAsString(), rv.TypeAsString())
	}

	return v.calculate(
		rv,
		func(v1 float64, v2 float64) float64 { return v1 + v2 },
		func(v1 int64, v2 int64) int64 { return v1 + v2 },
	)
}

func (v *Value) Subtract(rv *Value) (*Value, error) {
	if !v.IsNumber() || !rv.IsNumber() {
		return nil, fmt.Errorf("unable to substract %s from %s", v.TypeAsString(), rv.TypeAsString())
	}

	return v.calculate(
		rv,
		func(v1 float64, v2 float64) float64 { return v1 - v2 },
		func(v1 int64, v2 int64) int64 { return v1 - v2 },
	)
}

func (v *Value) Multiply(rv *Value) (*Value, error) {
	if !v.IsNumber() || !rv.IsNumber() {
		return nil, fmt.Errorf("unable to multiply %s to %s", v.TypeAsString(), rv.TypeAsString())
	}

	return v.calculate(
		rv,
		func(v1 float64, v2 float64) float64 { return v1 * v2 },
		func(v1 int64, v2 int64) int64 { return v1 * v2 },
	)
}

func (v *Value) Divide(rv *Value) (*Value, error) {
	if !v.IsNumber() || !rv.IsNumber() {
		return nil, fmt.Errorf("unable to divide %s to %s", v.TypeAsString(), rv.TypeAsString())
	}

	if (rv.Type == Integer && *rv.IntVal == 0) || (rv.Type == Float && *rv.FloatVal == 0) {
		return nil, fmt.Errorf("unable to divide on 0")
	}

	newValue := Value{Type: Float}

	lVal, err := v.ToFloat()
	if err != nil {
		return nil, err
	}

	rVal, err := rv.ToFloat()
	if err != nil {
		return nil, err
	}

	rez := (*lVal.FloatVal) / (*rVal.FloatVal)
	newValue.FloatVal = &rez

	return &newValue, nil
}

func (v *Value) Power(rv *Value) (*Value, error) {
	if !v.IsNumber() || !rv.IsNumber() {
		return nil, fmt.Errorf("unable to power %s to %s", v.TypeAsString(), rv.TypeAsString())
	}

	return v.calculate(
		rv,
		func(v1 float64, v2 float64) float64 { return math.Pow(v1, v2) },
		func(v1 int64, v2 int64) int64 { return int64(math.Pow(float64(v1), float64(v2))) },
	)
}

func (v *Value) IsNumber() bool {
	return v.Type == Float || v.Type == Integer
}

func (v *Value) IsMinusOrPlus() bool {
	if v.Type != Atom {
		return false
	}

	return *v.StringVal == "-" || *v.StringVal == "+"
}

func (v *Value) ToFloat() (*Value, error) {
	newValue := *v

	if v.Type == Float {
		return &newValue, nil
	} else if v.Type == Integer {
		intVal := *v.IntVal
		floatVal := float64(intVal)

		newValue.FloatVal = &floatVal
		newValue.IntVal = nil

		return &newValue, nil
	}

	return nil, fmt.Errorf("unable to convert %s to float", v.TypeAsString())
}

func (v *Value) ToString() (*Value, error) {
	newValue := *v
	newValue.Type = String

	if v.Type == Integer {
		newString := fmt.Sprintf("%d", *v.IntVal)
		newValue.StringVal = &newString
		newValue.IntVal = nil
	} else if v.Type == Float {
		newString := strconv.FormatFloat(*v.FloatVal, 'f', -1, 64)
		newValue.StringVal = &newString
		newValue.FloatVal = nil
	}

	return &newValue, nil
}

func (v *Value) calculate(rv *Value, fCb func(float64, float64) float64, iCb func(int64, int64) int64) (*Value, error) {
	newValue := Value{Type: Float}

	if v.Type == Integer && rv.Type == Integer {
		newValue.Type = Integer
		lVal := *v.IntVal
		rVal := *rv.IntVal
		rez := iCb(lVal, rVal)
		newValue.IntVal = &rez

		return &newValue, nil
	}

	lVal, err := v.ToFloat()
	if err != nil {
		return nil, err
	}

	rVal, err := rv.ToFloat()
	if err != nil {
		return nil, err
	}

	rez := fCb(*lVal.FloatVal, *rVal.FloatVal)
	newValue.FloatVal = &rez

	return &newValue, nil
}
