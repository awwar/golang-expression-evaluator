package program

import (
	"fmt"
	"math"
	"strings"

	"expression_parser/utility"
)

type NodeValueType int

var emptyFloat = NewFloat(0.0).(*FloatValue)
var emptyString = NewString("").(*StringValue)
var emptyBoolean = NewBoolean(false).(*BooleanValue)
var emptyInteger = NewInteger(0).(*IntValue)

func IsVariable(v Value) bool {
	return v.TypeName() == "string" && strings.HasPrefix(v.String(), "$")
}

func IsNumber(v Value) bool {
	return v.TypeName() == "integer" || v.TypeName() == "float"
}

func IsMinusOrPlus(v Value) bool {
	return v.TypeName() == "string" && (v.String() == "-" || v.String() == "+")
}

func NewString(val string) Value {
	return utility.AsPtr(StringValue(val))
}
func NewFloat(val float64) Value {
	return utility.AsPtr(FloatValue(val))
}
func NewBoolean(val bool) Value {
	return utility.AsPtr(BooleanValue(val))
}
func NewInteger(val int64) Value {
	return utility.AsPtr(IntValue(val))
}

type StringValue string

func (v *StringValue) String() string {
	return string(*v)
}

func (v *StringValue) TypeName() string {
	return "string"
}
func (v *StringValue) Add(rv Value) (Value, error) {
	strval, err := rv.ToString()
	return NewString(string(*v) + string(*strval)), err
}
func (v *StringValue) More(rv Value) (Value, error) {
	strval, err := rv.ToString()
	return NewBoolean(len(*v) > len(*strval)), err
}
func (v *StringValue) Less(rv Value) (Value, error) {
	strval, err := rv.ToString()
	return NewBoolean(len(*v) < len(*strval)), err
}
func (v *StringValue) Eq(rv Value) (Value, error) {
	strval, err := rv.ToString()
	return NewBoolean(string(*v) == string(*strval)), err
}
func (v *StringValue) Subtract(rv Value) (Value, error) {
	return emptyString, fmt.Errorf("cannot subtract %s from %s", v.TypeName(), rv.TypeName())
}
func (v *StringValue) Multiply(rv Value) (Value, error) {
	return emptyString, fmt.Errorf("cannot multiply %s to %s", v.TypeName(), rv.TypeName())
}
func (v *StringValue) Divide(rv Value) (Value, error) {
	return emptyString, fmt.Errorf("cannot divide %s by %s", v.TypeName(), rv.TypeName())
}
func (v *StringValue) Power(rv Value) (Value, error) {
	return emptyString, fmt.Errorf("cannot power %s with %s", v.TypeName(), rv.TypeName())
}
func (v *StringValue) ToInteger() (*IntValue, error) {
	return emptyInteger, fmt.Errorf("unable to convert %s to integer", v.TypeName())
}
func (v *StringValue) ToFloat() (*FloatValue, error) {
	return emptyFloat, fmt.Errorf("unable to convert %s to float", v.TypeName())
}
func (v *StringValue) ToBoolean() (*BooleanValue, error) {
	return NewBoolean(len(*v) > 0).(*BooleanValue), nil
}
func (v *StringValue) ToString() (*StringValue, error) {
	return NewString(string(*v)).(*StringValue), nil
}

type FloatValue float64

func (v *FloatValue) String() string {
	return fmt.Sprintf("%f", *v)
}

func (v *FloatValue) TypeName() string {
	return "float"
}
func (v *FloatValue) Add(rv Value) (Value, error) {
	floatval, err := rv.ToFloat()
	return NewFloat(float64(*v) + float64(*floatval)), err
}
func (v *FloatValue) More(rv Value) (Value, error) {
	floatval, err := rv.ToFloat()
	return NewBoolean(float64(*v) > float64(*floatval)), err
}
func (v *FloatValue) Less(rv Value) (Value, error) {
	floatval, err := rv.ToFloat()
	return NewBoolean(float64(*v) < float64(*floatval)), err
}
func (v *FloatValue) Eq(rv Value) (Value, error) {
	floatval, err := rv.ToFloat()
	return NewBoolean(float64(*v) == float64(*floatval)), err
}
func (v *FloatValue) Subtract(rv Value) (Value, error) {
	floatval, err := rv.ToFloat()
	return NewFloat(float64(*v) - float64(*floatval)), err
}
func (v *FloatValue) Multiply(rv Value) (Value, error) {
	floatval, err := rv.ToFloat()
	return NewFloat(float64(*v) * float64(*floatval)), err
}
func (v *FloatValue) Divide(rv Value) (Value, error) {
	floatval, err := rv.ToFloat()
	return NewFloat(float64(*v) / float64(*floatval)), err
}
func (v *FloatValue) Power(rv Value) (Value, error) {
	floatval, err := rv.ToFloat()
	return NewFloat(math.Pow(float64(*v), float64(*floatval))), err
}
func (v *FloatValue) ToInteger() (*IntValue, error) {
	return NewInteger(int64(*v)).(*IntValue), nil
}
func (v *FloatValue) ToFloat() (*FloatValue, error) {
	return NewFloat(float64(*v)).(*FloatValue), nil
}
func (v *FloatValue) ToBoolean() (*BooleanValue, error) {
	return NewBoolean(*v > 0).(*BooleanValue), nil
}
func (v *FloatValue) ToString() (*StringValue, error) {
	return NewString(fmt.Sprintf("%f", *v)).(*StringValue), nil
}

type IntValue int64

func (v *IntValue) String() string {
	return fmt.Sprintf("%d", *v)
}

func (v *IntValue) TypeName() string {
	return "integer"
}
func (v *IntValue) Add(rv Value) (Value, error) {
	intVal, err := rv.ToInteger()
	return NewInteger(int64(*v) + int64(*intVal)), err
}
func (v *IntValue) More(rv Value) (Value, error) {
	intVal, err := rv.ToInteger()
	return NewBoolean(int64(*v) > int64(*intVal)), err
}
func (v *IntValue) Less(rv Value) (Value, error) {
	intVal, err := rv.ToInteger()
	return NewBoolean(int64(*v) < int64(*intVal)), err
}
func (v *IntValue) Eq(rv Value) (Value, error) {
	intVal, err := rv.ToInteger()
	return NewBoolean(int64(*v) == int64(*intVal)), err
}
func (v *IntValue) Subtract(rv Value) (Value, error) {
	intVal, err := rv.ToInteger()
	return NewInteger(int64(*v) - int64(*intVal)), err
}
func (v *IntValue) Multiply(rv Value) (Value, error) {
	intVal, err := rv.ToInteger()
	return NewInteger(int64(*v) * int64(*intVal)), err
}
func (v *IntValue) Divide(rv Value) (Value, error) {
	intVal, err := rv.ToInteger()
	return NewInteger(int64(*v) / int64(*intVal)), err
}
func (v *IntValue) Power(rv Value) (Value, error) {
	intVal, err := rv.ToInteger()
	return NewInteger(int64(math.Pow(float64(*v), float64(*intVal)))), err
}
func (v *IntValue) ToInteger() (*IntValue, error) {
	return NewInteger(int64(*v)).(*IntValue), nil
}
func (v *IntValue) ToFloat() (*FloatValue, error) {
	return NewFloat(float64(*v)).(*FloatValue), nil
}
func (v *IntValue) ToBoolean() (*BooleanValue, error) {
	return NewBoolean(*v > 0).(*BooleanValue), nil
}
func (v *IntValue) ToString() (*StringValue, error) {
	return NewString(fmt.Sprintf("%d", *v)).(*StringValue), nil
}

type BooleanValue bool

func (v *BooleanValue) String() string {
	if *v == true {
		return "true"
	}
	return "false"
}

func (v *BooleanValue) TypeName() string {
	return "boolean"
}
func (v *BooleanValue) Add(rv Value) (Value, error) {
	return emptyBoolean, fmt.Errorf("cannot subtract %s from %s", v.TypeName(), rv.TypeName())
}
func (v *BooleanValue) More(rv Value) (Value, error) {
	return emptyBoolean, fmt.Errorf("cannot subtract %s from %s", v.TypeName(), rv.TypeName())
}
func (v *BooleanValue) Less(rv Value) (Value, error) {
	return emptyBoolean, fmt.Errorf("cannot subtract %s from %s", v.TypeName(), rv.TypeName())
}
func (v *BooleanValue) Eq(rv Value) (Value, error) {
	boolVal, err := rv.ToBoolean()
	return NewBoolean(bool(*v) == bool(*boolVal)), err
}
func (v *BooleanValue) Subtract(rv Value) (Value, error) {
	return emptyBoolean, fmt.Errorf("cannot subtract %s from %s", v.TypeName(), rv.TypeName())
}
func (v *BooleanValue) Multiply(rv Value) (Value, error) {
	return emptyBoolean, fmt.Errorf("cannot subtract %s from %s", v.TypeName(), rv.TypeName())
}
func (v *BooleanValue) Divide(rv Value) (Value, error) {
	return emptyBoolean, fmt.Errorf("cannot subtract %s from %s", v.TypeName(), rv.TypeName())
}
func (v *BooleanValue) Power(rv Value) (Value, error) {
	return emptyBoolean, fmt.Errorf("cannot subtract %s from %s", v.TypeName(), rv.TypeName())
}
func (v *BooleanValue) ToInteger() (*IntValue, error) {
	val := 1
	if !*v {
		val = 0
	}
	return NewInteger(int64(val)).(*IntValue), nil
}
func (v *BooleanValue) ToFloat() (*FloatValue, error) {
	val := 1.0
	if !*v {
		val = 0.0
	}
	return NewFloat(val).(*FloatValue), nil
}
func (v *BooleanValue) ToBoolean() (*BooleanValue, error) {
	return NewBoolean(bool(*v)).(*BooleanValue), nil
}
func (v *BooleanValue) ToString() (*StringValue, error) {
	return NewString(v.String()).(*StringValue), nil
}

type Value interface {
	fmt.Stringer
	TypeName() string
	Add(rv Value) (Value, error)
	More(rv Value) (Value, error)
	Less(rv Value) (Value, error)
	Eq(rv Value) (Value, error)
	Subtract(rv Value) (Value, error)
	Multiply(rv Value) (Value, error)
	Divide(rv Value) (Value, error)
	Power(rv Value) (Value, error)
	ToInteger() (*IntValue, error)
	ToFloat() (*FloatValue, error)
	ToBoolean() (*BooleanValue, error)
	ToString() (*StringValue, error)
}

//// FOR REPR PURPOSE ONLY!!!
//func (v *Value) String() string {
//	asStringValue, _ := v.ToString()
//
//	if v.ValueType == String {
//		return fmt.Sprintf(`"%s"`, *asStringValue.StringVal)
//	}
//
//	return *asStringValue.StringVal
//}
//
//func (v *Value) TypeName() string {
//	return MapTypeToTypeName[v.ValueType]
//}
//
//func (v *Value) Add(rv *Value) (*Value, error) {
//	if v.ValueType == String || rv.ValueType == String {
//		leftValue, err := v.ToString()
//		if err != nil {
//			return nil, err
//		}
//
//		rightValue, err := rv.ToString()
//		if err != nil {
//			return nil, err
//		}
//
//		stVal := fmt.Sprintf("%s%s", *leftValue.StringVal, *rightValue.StringVal)
//
//		return &Value{StringVal: &stVal, ValueType: String}, nil
//	}
//
//	if !v.IsNumber() || !rv.IsNumber() {
//		return nil, fmt.Errorf("unable to add %s to %s", v.TypeName(), rv.TypeName())
//	}
//
//	return v.calculate(
//		rv,
//		func(v1 float64, v2 float64) float64 { return v1 + v2 },
//		func(v1 int64, v2 int64) int64 { return v1 + v2 },
//	)
//}
//
//func (v *Value) More(rv *Value) (*Value, error) {
//	if !v.IsNumber() || !rv.IsNumber() {
//		return nil, fmt.Errorf("unable to > with %s to %s", v.TypeName(), rv.TypeName())
//	}
//
//	return v.cmp(
//		rv,
//		func(v1 float64, v2 float64) bool { return v1 > v2 },
//		func(v1 int64, v2 int64) bool { return v1 > v2 },
//	)
//}
//
//func (v *Value) Less(rv *Value) (*Value, error) {
//	if !v.IsNumber() || !rv.IsNumber() {
//		return nil, fmt.Errorf("unable to < with %s to %s", v.TypeName(), rv.TypeName())
//	}
//
//	return v.cmp(
//		rv,
//		func(v1 float64, v2 float64) bool { return v1 < v2 },
//		func(v1 int64, v2 int64) bool { return v1 < v2 },
//	)
//}
//
//func (v *Value) Eq(rv *Value) (*Value, error) {
//	if !v.IsNumber() || !rv.IsNumber() {
//		return nil, fmt.Errorf("unable to = with %s to %s", v.TypeName(), rv.TypeName())
//	}
//
//	return v.cmp(
//		rv,
//		func(v1 float64, v2 float64) bool { return v1 == v2 },
//		func(v1 int64, v2 int64) bool { return v1 == v2 },
//	)
//}
//
//func (v *Value) Subtract(rv *Value) (*Value, error) {
//	if !v.IsNumber() || !rv.IsNumber() {
//		return nil, fmt.Errorf("unable to subtract %s from %s", v.TypeName(), rv.TypeName())
//	}
//
//	return v.calculate(
//		rv,
//		func(v1 float64, v2 float64) float64 { return v1 - v2 },
//		func(v1 int64, v2 int64) int64 { return v1 - v2 },
//	)
//}
//
//func (v *Value) Multiply(rv *Value) (*Value, error) {
//	if !v.IsNumber() || !rv.IsNumber() {
//		return nil, fmt.Errorf("unable to multiply %s to %s", v.TypeName(), rv.TypeName())
//	}
//
//	return v.calculate(
//		rv,
//		func(v1 float64, v2 float64) float64 { return v1 * v2 },
//		func(v1 int64, v2 int64) int64 { return v1 * v2 },
//	)
//}
//
//func (v *Value) Divide(rv *Value) (*Value, error) {
//	if !v.IsNumber() || !rv.IsNumber() {
//		return nil, fmt.Errorf("unable to divide %s to %s", v.TypeName(), rv.TypeName())
//	}
//
//	if (rv.ValueType == Integer && *rv.IntVal == 0) || (rv.ValueType == Float && *rv.FloatVal == 0) {
//		return nil, fmt.Errorf("unable to divide on 0")
//	}
//
//	newValue := Value{ValueType: Float}
//
//	lVal, err := v.ToFloat()
//	if err != nil {
//		return nil, err
//	}
//
//	rVal, err := rv.ToFloat()
//	if err != nil {
//		return nil, err
//	}
//
//	rez := (*lVal.FloatVal) / (*rVal.FloatVal)
//	newValue.FloatVal = &rez
//
//	return &newValue, nil
//}
//
//func (v *Value) Power(rv *Value) (*Value, error) {
//	if !v.IsNumber() || !rv.IsNumber() {
//		return nil, fmt.Errorf("unable to power %s to %s", v.TypeName(), rv.TypeName())
//	}
//
//	return v.calculate(
//		rv,
//		func(v1 float64, v2 float64) float64 { return math.Pow(v1, v2) },
//		func(v1 int64, v2 int64) int64 { return int64(math.Pow(float64(v1), float64(v2))) },
//	)
//}
//
//func (v *Value) IsNumber() bool {
//	return v.ValueType == Float || v.ValueType == Integer
//}
//
//func (v *Value) IsBoolean() bool {
//	return v.ValueType == Boolean
//}
//
//func (v *Value) IsAtom() bool {
//	return v.ValueType == Atom
//}
//
//func (v *Value) IsVariable() bool {
//	return v.IsAtom() && strings.HasPrefix(*v.StringVal, "$")
//}
//
//func (v *Value) IsMinusOrPlus() bool {
//	if !v.IsAtom() {
//		return false
//	}
//
//	return *v.StringVal == "-" || *v.StringVal == "+"
//}
//
//func (v *Value) ToFloat() (*Value, error) {
//	newValue := *v
//
//	if v.ValueType == Float {
//		return &newValue, nil
//	} else if v.ValueType == Integer {
//		intVal := *v.IntVal
//		floatVal := float64(intVal)
//
//		newValue.FloatVal = &floatVal
//		newValue.IntVal = nil
//
//		return &newValue, nil
//	}
//
//	return nil, fmt.Errorf("unable to convert %s to float", v.TypeName())
//}
//
//func (v *Value) ToBoolean() (*Value, error) {
//	newValue := *v
//
//	if v.ValueType == Boolean {
//		return &newValue, nil
//	}
//	if v.ValueType == Float {
//		newValue.ValueType = Boolean
//		newValue.BoolVal = utility.AsPtr(*v.FloatVal != 0.0)
//		newValue.FloatVal = nil
//	} else if v.ValueType == Integer {
//		newValue.ValueType = Boolean
//		newValue.BoolVal = utility.AsPtr(*v.IntVal != 0)
//		newValue.IntVal = nil
//	} else if v.ValueType == String {
//		newValue.ValueType = Boolean
//		newValue.BoolVal = utility.AsPtr(*v.StringVal != "")
//		newValue.StringVal = nil
//	}
//
//	return nil, fmt.Errorf("unable to convert %s to bool", v.TypeName())
//}
//
//func (v *Value) ToString() (*Value, error) {
//	newValue := *v
//	newValue.ValueType = String
//
//	if v.ValueType == Integer {
//		newString := fmt.Sprintf("%d", *v.IntVal)
//		newValue.StringVal = &newString
//		newValue.IntVal = nil
//	} else if v.ValueType == Float {
//		newString := strconv.FormatFloat(*v.FloatVal, 'f', -1, 64)
//		newValue.StringVal = &newString
//		newValue.FloatVal = nil
//	} else if v.ValueType == Boolean {
//		newString := fmt.Sprint(*newValue.BoolVal)
//		newValue.StringVal = &newString
//		newValue.BoolVal = nil
//	}
//
//	return &newValue, nil
//}
//
//func (v *Value) calculate(rv *Value, fCb func(float64, float64) float64, iCb func(int64, int64) int64) (*Value, error) {
//	newValue := Value{ValueType: Float}
//
//	if v.ValueType == Integer && rv.ValueType == Integer {
//		newValue.ValueType = Integer
//		lVal := *v.IntVal
//		rVal := *rv.IntVal
//		rez := iCb(lVal, rVal)
//		newValue.IntVal = &rez
//
//		return &newValue, nil
//	}
//
//	lVal, err := v.ToFloat()
//	if err != nil {
//		return nil, err
//	}
//
//	rVal, err := rv.ToFloat()
//	if err != nil {
//		return nil, err
//	}
//
//	rez := fCb(*lVal.FloatVal, *rVal.FloatVal)
//	newValue.FloatVal = &rez
//
//	return &newValue, nil
//}
//
//func (v *Value) cmp(rv *Value, fCb func(float64, float64) bool, iCb func(int64, int64) bool) (*Value, error) {
//	newValue := Value{ValueType: Boolean}
//
//	if v.ValueType == Integer && rv.ValueType == Integer {
//		newValue.ValueType = Integer
//		lVal := *v.IntVal
//		rVal := *rv.IntVal
//		newValue.BoolVal = utility.AsPtr(iCb(lVal, rVal))
//
//		return &newValue, nil
//	}
//
//	lVal, err := v.ToFloat()
//	if err != nil {
//		return nil, err
//	}
//
//	rVal, err := rv.ToFloat()
//	if err != nil {
//		return nil, err
//	}
//
//	newValue.BoolVal = utility.AsPtr(fCb(*lVal.FloatVal, *rVal.FloatVal))
//
//	return &newValue, nil
//}
