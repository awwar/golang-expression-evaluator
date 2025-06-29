package program

import (
	"fmt"
	"math"
	"strings"

	"expression_parser/utility"
)

const (
	StringType  = "string"
	IntegerType = "integer"
	FloatType   = "float"
	BooleanType = "boolean"
)

type NodeValueType int

var emptyFloat = NewFloat(0.0).(*FloatValue)
var emptyInteger = NewInteger(0).(*IntValue)

func IsVariable(v Value) bool {
	return v.TypeName() == StringType && strings.HasPrefix(v.String(), "$")
}

func IsNumber(v Value) bool {
	return v.TypeName() == IntegerType || v.TypeName() == FloatType
}

func IsMinusOrPlus(v Value) bool {
	return v.TypeName() == StringType && (v.String() == "-" || v.String() == "+")
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
	return StringType
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
	return v, nil
}

type FloatValue float64

func (v *FloatValue) String() string {
	return fmt.Sprintf("%f", *v)
}

func (v *FloatValue) TypeName() string {
	return FloatType
}
func (v *FloatValue) ToInteger() (*IntValue, error) {
	return NewInteger(int64(*v)).(*IntValue), nil
}
func (v *FloatValue) ToFloat() (*FloatValue, error) {
	return v, nil
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
	return IntegerType
}
func (v *IntValue) ToInteger() (*IntValue, error) {
	return v, nil
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
	if *v {
		return "true"
	}
	return "false"
}

func (v *BooleanValue) TypeName() string {
	return BooleanType
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
	return v, nil
}
func (v *BooleanValue) ToString() (*StringValue, error) {
	return NewString(v.String()).(*StringValue), nil
}

type Value interface {
	fmt.Stringer
	TypeName() string
	ToInteger() (*IntValue, error)
	ToFloat() (*FloatValue, error)
	ToBoolean() (*BooleanValue, error)
	ToString() (*StringValue, error)
}

func Add(lv Value, rv Value) (Value, error) {
	nlv, nrv, err := convertionPoisining(lv, rv)
	if err != nil {
		return lv, err
	}
	switch nlv.TypeName() {
	case IntegerType:
		return NewInteger(int64(*nlv.(*IntValue)) + int64(*nrv.(*IntValue))), nil
	case FloatType:
		return NewFloat(float64(*nlv.(*FloatValue)) + float64(*nrv.(*FloatValue))), nil
	case StringType:
		return NewString(string(*nlv.(*StringValue)) + string(*nrv.(*StringValue))), nil
	case BooleanType:
		return NewBoolean(bool(*nlv.(*BooleanValue)) || bool(*nrv.(*BooleanValue))), nil
	}
	return lv, fmt.Errorf("cannot add %s to %s", nlv.TypeName(), nrv.TypeName())
}
func More(lv Value, rv Value) (Value, error) {
	nlv, nrv, err := convertionPoisining(lv, rv)
	if err != nil {
		return lv, err
	}
	switch nlv.TypeName() {
	case IntegerType:
		return NewBoolean(int64(*nlv.(*IntValue)) > int64(*nrv.(*IntValue))), nil
	case FloatType:
		return NewBoolean(float64(*nlv.(*FloatValue)) > float64(*nrv.(*FloatValue))), nil
	case StringType:
		return NewBoolean(string(*nlv.(*StringValue)) > string(*nrv.(*StringValue))), nil
	}
	return lv, fmt.Errorf("cannot assert then %s > %s", nlv.TypeName(), nrv.TypeName())
}
func Less(lv Value, rv Value) (Value, error) {
	nlv, nrv, err := convertionPoisining(lv, rv)
	if err != nil {
		return lv, err
	}
	switch nlv.TypeName() {
	case IntegerType:
		return NewBoolean(int64(*nlv.(*IntValue)) < int64(*nrv.(*IntValue))), nil
	case FloatType:
		return NewBoolean(float64(*nlv.(*FloatValue)) < float64(*nrv.(*FloatValue))), nil
	case StringType:
		return NewBoolean(string(*nlv.(*StringValue)) < string(*nrv.(*StringValue))), nil
	}
	return lv, fmt.Errorf("cannot assert then %s < %s", nlv.TypeName(), nrv.TypeName())
}
func Eq(lv Value, rv Value) (Value, error) {
	nlv, nrv, err := convertionPoisining(lv, rv)
	if err != nil {
		return lv, err
	}
	switch nlv.TypeName() {
	case IntegerType:
		return NewBoolean(int64(*nlv.(*IntValue)) == int64(*nrv.(*IntValue))), nil
	case FloatType:
		return NewBoolean(float64(*nlv.(*FloatValue)) == float64(*nrv.(*FloatValue))), nil
	case StringType:
		return NewBoolean(string(*nlv.(*StringValue)) == string(*nrv.(*StringValue))), nil
	case BooleanType:
		return NewBoolean(bool(*nlv.(*BooleanValue)) == bool(*nrv.(*BooleanValue))), nil
	}
	return lv, fmt.Errorf("cannot assert then %s == %s", nlv.TypeName(), nrv.TypeName())
}
func Subtract(lv Value, rv Value) (Value, error) {
	nlv, nrv, err := convertionPoisining(lv, rv)
	if err != nil {
		return lv, err
	}
	switch nlv.TypeName() {
	case IntegerType:
		return NewInteger(int64(*nlv.(*IntValue)) - int64(*nrv.(*IntValue))), nil
	case FloatType:
		return NewFloat(float64(*nlv.(*FloatValue)) - float64(*nrv.(*FloatValue))), nil
	}
	return lv, fmt.Errorf("cannot substruct %s from %s", nlv.TypeName(), nrv.TypeName())
}
func Multiply(lv Value, rv Value) (Value, error) {
	nlv, nrv, err := convertionPoisining(lv, rv)
	if err != nil {
		return lv, err
	}
	switch nlv.TypeName() {
	case IntegerType:
		return NewInteger(int64(*nlv.(*IntValue)) * int64(*nrv.(*IntValue))), nil
	case FloatType:
		return NewFloat(float64(*nlv.(*FloatValue)) * float64(*nrv.(*FloatValue))), nil
	case BooleanType:
		return NewBoolean(bool(*nlv.(*BooleanValue)) && bool(*nrv.(*BooleanValue))), nil
	}
	return lv, fmt.Errorf("cannot multiply %s on %s", nlv.TypeName(), nrv.TypeName())
}
func Divide(lv Value, rv Value) (Value, error) {
	nlv, nrv, err := convertionPoisining(lv, rv)
	if err != nil {
		return lv, err
	}
	switch nlv.TypeName() {
	case IntegerType:
		return NewInteger(int64(*nlv.(*IntValue)) / int64(*nrv.(*IntValue))), nil
	case FloatType:
		return NewFloat(float64(*nlv.(*FloatValue)) / float64(*nrv.(*FloatValue))), nil
	}
	return lv, fmt.Errorf("cannot devide %s by %s", nlv.TypeName(), nrv.TypeName())
}
func Power(lv Value, rv Value) (Value, error) {
	nlv, nrv, err := convertionPoisining(lv, rv)
	if err != nil {
		return lv, err
	}
	switch nlv.TypeName() {
	case IntegerType:
		v := int64(math.Pow(float64(*nlv.(*IntValue)), float64(*nrv.(*IntValue))))
		return NewInteger(v), nil
	case FloatType:
		return NewFloat(math.Pow(float64(*nlv.(*FloatValue)), float64(*nrv.(*FloatValue)))), nil
	}
	return lv, fmt.Errorf("cannot power %s with %s", nlv.TypeName(), nrv.TypeName())
}

func convertionPoisining(lv Value, rv Value) (Value, Value, error) {
	if lv.TypeName() == rv.TypeName() {
		return lv, rv, nil
	}

	if lv.TypeName() == StringType || rv.TypeName() == StringType {
		nlv, lerr := lv.ToString()
		nrv, rerr := rv.ToString()
		return nlv, nrv, chooseErr(lerr, rerr)
	}

	if lv.TypeName() == FloatType || rv.TypeName() == FloatType {
		nlv, lerr := lv.ToFloat()
		nrv, rerr := rv.ToFloat()
		return nlv, nrv, chooseErr(lerr, rerr)
	}

	if lv.TypeName() == IntegerType || rv.TypeName() == IntegerType {
		nlv, lerr := lv.ToInteger()
		nrv, rerr := rv.ToInteger()
		return nlv, nrv, chooseErr(lerr, rerr)
	}

	if lv.TypeName() == BooleanType || rv.TypeName() == BooleanType {
		nlv, lerr := lv.ToBoolean()
		nrv, rerr := rv.ToBoolean()
		return nlv, nrv, chooseErr(lerr, rerr)
	}

	return lv, rv, fmt.Errorf("cannot convert %s to %s", lv.TypeName(), rv.TypeName())
}

func chooseErr(lerr error, rerr error) error {
	if rerr != nil {
		return lerr
	}

	return rerr
}
