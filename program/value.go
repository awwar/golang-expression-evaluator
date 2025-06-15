package program

import (
	"fmt"
	"math"
	"strconv"

	"expression_parser/utility"
)

type NodeValueType int

const (
	Integer NodeValueType = iota
	Float   NodeValueType = iota
	Atom    NodeValueType = iota
	Boolean NodeValueType = iota
	String  NodeValueType = iota
)

var MapTypeToTypeName = map[NodeValueType]string{Integer: "int", Float: "float", Atom: "atom", String: "string", Boolean: "bool"}

func NewString(val string) *Value {
	return &Value{ValueType: String, StringVal: &val}
}

func NewFloat(val float64) *Value {
	return &Value{ValueType: Float, FloatVal: &val}
}

type Value struct {
	ValueType NodeValueType
	BoolVal   *bool
	StringVal *string
	FloatVal  *float64
	IntVal    *int64
}

// FOR REPR PURPOSE ONLY!!!
func (v *Value) String() string {
	asStringValue, _ := v.ToString()

	if v.ValueType == String {
		return fmt.Sprintf(`"%s"`, *asStringValue.StringVal)
	}

	if v.ValueType == String {
		return fmt.Sprintf(`"%s"`, *asStringValue.StringVal)
	}

	return *asStringValue.StringVal
}

func (v *Value) GoString() string {
	return v.String()
}

func (v *Value) TypeAsString() string {
	return MapTypeToTypeName[v.ValueType]
}

func (v *Value) Add(rv *Value) (*Value, error) {
	if v.ValueType == String || rv.ValueType == String {
		leftValue, err := v.ToString()
		if err != nil {
			return nil, err
		}

		rightValue, err := rv.ToString()
		if err != nil {
			return nil, err
		}

		stVal := fmt.Sprintf("%s%s", *leftValue.StringVal, *rightValue.StringVal)

		return &Value{StringVal: &stVal, ValueType: String}, nil
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

func (v *Value) More(rv *Value) (*Value, error) {
	if !v.IsNumber() || !rv.IsNumber() {
		return nil, fmt.Errorf("unable to > with %s to %s", v.TypeAsString(), rv.TypeAsString())
	}

	return v.cmp(
		rv,
		func(v1 float64, v2 float64) bool { return v1 > v2 },
		func(v1 int64, v2 int64) bool { return v1 > v2 },
	)
}

func (v *Value) Less(rv *Value) (*Value, error) {
	if !v.IsNumber() || !rv.IsNumber() {
		return nil, fmt.Errorf("unable to < with %s to %s", v.TypeAsString(), rv.TypeAsString())
	}

	return v.cmp(
		rv,
		func(v1 float64, v2 float64) bool { return v1 < v2 },
		func(v1 int64, v2 int64) bool { return v1 < v2 },
	)
}

func (v *Value) Eq(rv *Value) (*Value, error) {
	if !v.IsNumber() || !rv.IsNumber() {
		return nil, fmt.Errorf("unable to = with %s to %s", v.TypeAsString(), rv.TypeAsString())
	}

	return v.cmp(
		rv,
		func(v1 float64, v2 float64) bool { return v1 == v2 },
		func(v1 int64, v2 int64) bool { return v1 == v2 },
	)
}

func (v *Value) Subtract(rv *Value) (*Value, error) {
	if !v.IsNumber() || !rv.IsNumber() {
		return nil, fmt.Errorf("unable to subtract %s from %s", v.TypeAsString(), rv.TypeAsString())
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

	if (rv.ValueType == Integer && *rv.IntVal == 0) || (rv.ValueType == Float && *rv.FloatVal == 0) {
		return nil, fmt.Errorf("unable to divide on 0")
	}

	newValue := Value{ValueType: Float}

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
	return v.ValueType == Float || v.ValueType == Integer
}

func (v *Value) IsBoolean() bool {
	return v.ValueType == Boolean
}

func (v *Value) IsAtom() bool {
	return v.ValueType == Atom
}

func (v *Value) IsMinusOrPlus() bool {
	if !v.IsAtom() {
		return false
	}

	return *v.StringVal == "-" || *v.StringVal == "+"
}

func (v *Value) ToFloat() (*Value, error) {
	newValue := *v

	if v.ValueType == Float {
		return &newValue, nil
	} else if v.ValueType == Integer {
		intVal := *v.IntVal
		floatVal := float64(intVal)

		newValue.FloatVal = &floatVal
		newValue.IntVal = nil

		return &newValue, nil
	}

	return nil, fmt.Errorf("unable to convert %s to float", v.TypeAsString())
}

func (v *Value) ToBoolean() (*Value, error) {
	newValue := *v

	if v.ValueType == Boolean {
		return &newValue, nil
	}
	if v.ValueType == Float {
		newValue.ValueType = Boolean
		newValue.BoolVal = utility.AsPtr(*v.FloatVal != 0.0)
		newValue.FloatVal = nil
	} else if v.ValueType == Integer {
		newValue.ValueType = Boolean
		newValue.BoolVal = utility.AsPtr(*v.IntVal != 0)
		newValue.IntVal = nil
	} else if v.ValueType == String {
		newValue.ValueType = Boolean
		newValue.BoolVal = utility.AsPtr(*v.StringVal != "")
		newValue.StringVal = nil
	}

	return nil, fmt.Errorf("unable to convert %s to bool", v.TypeAsString())
}

func (v *Value) ToString() (*Value, error) {
	newValue := *v
	newValue.ValueType = String

	if v.ValueType == Integer {
		newString := fmt.Sprintf("%d", *v.IntVal)
		newValue.StringVal = &newString
		newValue.IntVal = nil
	} else if v.ValueType == Float {
		newString := strconv.FormatFloat(*v.FloatVal, 'f', -1, 64)
		newValue.StringVal = &newString
		newValue.FloatVal = nil
	} else if v.ValueType == Boolean {
		newString := fmt.Sprint(*newValue.BoolVal)
		newValue.StringVal = &newString
		newValue.BoolVal = nil
	}

	return &newValue, nil
}

func (v *Value) calculate(rv *Value, fCb func(float64, float64) float64, iCb func(int64, int64) int64) (*Value, error) {
	newValue := Value{ValueType: Float}

	if v.ValueType == Integer && rv.ValueType == Integer {
		newValue.ValueType = Integer
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

func (v *Value) cmp(rv *Value, fCb func(float64, float64) bool, iCb func(int64, int64) bool) (*Value, error) {
	newValue := Value{ValueType: Boolean}

	if v.ValueType == Integer && rv.ValueType == Integer {
		newValue.ValueType = Integer
		lVal := *v.IntVal
		rVal := *rv.IntVal
		newValue.BoolVal = utility.AsPtr(iCb(lVal, rVal))

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

	newValue.BoolVal = utility.AsPtr(fCb(*lVal.FloatVal, *rVal.FloatVal))

	return &newValue, nil
}
