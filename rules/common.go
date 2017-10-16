package rules

import (
	"go/ast"
	"go/token"
	"reflect"
	"strconv"
)


// SupportedVersions : A list of PostgreSQL version supported by this software
// Old versions (even not currently supported by the community) are available too.
// The idea is support any version greater than 9.0.
var SupportedVersions = []float32{9.0, 9.1, 9.2, 9.3, 9.4, 9.5, 9.6, 10.0}

// ParameterType : Defines a type to contain a parameter
type ParameterType int

const (
	// BytesParameter : This value can be formated as bytes
	BytesParameter ParameterType = 1 + iota

	// NumericParameter : This value should be displayed as number
	NumericParameter

	// StringParameter : This value should be displayed as string
	StringParameter

	// TimeParameter : This value should be displayed as time
	TimeParameter
)

type EnvironmentName int

const (
	WebEnvironment EnvironmentName = 1 + iota
	OLTPEnvironment
	DWEnvironment
	MixedEnvironment
	DesktopEnvironment
)

// DatabaseParameter : Provides a parameter details
type DatabaseParameter struct {
	Name         string
	Value        int
	MaxValue     int
	DefaultValue int
	Type         ParameterType
	Rule         string
}

const (
	KILOBYTE = 1024
	MEGABYTE = KILOBYTE * 1024
	GIGABYTE = MEGABYTE * 1024
)

// important! https://thorstenball.com/blog/2016/11/16/putting-eval-in-go/
func eval(exp ast.Expr) int {
	switch exp := exp.(type) {
	case *ast.BinaryExpr:
		return evalBinaryExpr(exp)
	case *ast.BasicLit:
		switch exp.Kind {
		case token.INT:
			i, _ := strconv.Atoi(exp.Value)
			return i
		}
	}

	return 0
}

func evalBinaryExpr(exp *ast.BinaryExpr) int {
	left := eval(exp.X)
	right := eval(exp.Y)

	switch exp.Op {
	case token.ADD:
		return left + right
	case token.SUB:
		return left - right
	case token.MUL:
		return left * right
	case token.QUO:
		return left / right
	}

	return 0
}

// source: https://stackoverflow.com/questions/10485743/contains-method-for-a-slice
func hasElem(s interface{}, elem interface{}) bool {
	arrV := reflect.ValueOf(s)

	if arrV.Kind() == reflect.Slice {
		for i := 0; i < arrV.Len(); i++ {

			// XXX - panics if slice element points to an unexported struct field
			// see https://golang.org/pkg/reflect/#Value.Interface
			if arrV.Index(i).Interface() == elem {
				return true
			}
		}
	}

	return false
}

// fix the value respecting the minimum(default) and the maximium value (it avoids to check for each parameter)
func fixValue(p *DatabaseParameter, value int, pgVersion float32) int {
	result := value

	if p.DefaultValue > 0 && value < p.DefaultValue {
		result = p.DefaultValue
	} else if p.MaxValue > 0 && value > p.MaxValue {
		result = p.MaxValue
	}

	if !hasElem(SupportedVersions, pgVersion) {
		result = -1
	}

	return result
}
