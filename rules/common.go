package rules

import (
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"strconv"
	"strings"
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

// OSFamily : Defines a operational system familly
type OSFamily string

const (
	// WindowsOS : Defines Windows OS
	WindowsOS OSFamily = "Windows"

	// LinuxOS : Defines Linux OS
	LinuxOS OSFamily = "Linux"

	// UnixOS : Defines Unix OS
	UnixOS OSFamily = "Unix"
)

// EnvironmentName : Defines the environment name
type EnvironmentName string

const (
	// WebEnvironment : Defines the "WEB" environment
	WebEnvironment EnvironmentName = "WEB"

	// OLTPEnvironment : Defines the "OLTP" environment
	OLTPEnvironment EnvironmentName = "OLTP"

	// DWEnvironment : Defines the "DW" environment
	DWEnvironment EnvironmentName = "DW"

	// MixedEnvironment : Defines the "Mixed" environment
	MixedEnvironment EnvironmentName = "Mixed"

	// DesktopEnvironment : Defines the "Desktop" environment
	DesktopEnvironment EnvironmentName = "Desktop"
)

// ParameterCategory : Defines a category for a parameter
type ParameterCategory string

const (

	// MemoryRelatedCategory : Defines a memory category
	MemoryRelatedCategory ParameterCategory = "Memory Configuration"

	// ChekPointRelatedCategory : Defines a checkpoint category
	ChekPointRelatedCategory ParameterCategory = "Checkpoint Related Configuration"

	// NetworkRelatedCategory : Defines a network category
	NetworkRelatedCategory ParameterCategory = "Network Related Configuration"
)

// DatabaseParameter : Provides a parameter details
type DatabaseParameter struct {
	Name         string
	Value        int
	MaxValue     int
	DefaultValue int
	Type         ParameterType
	Rule         string
	Abstract     string
	Category     ParameterCategory
	Articles     []ArticleRecommendation
	DocURLSuffix string
}

// ArticleRecommendation : contains a article related with the paramater
type ArticleRecommendation struct {
	Title string
	URL   string
}

const (
	// KILOBYTE : Defined in bytes
	KILOBYTE = 1024

	// MEGABYTE : Defined in bytes
	MEGABYTE = KILOBYTE * 1024

	// GIGABYTE : Defined in bytes
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

// ParameterArgs : contains a argument to compute a parameter
type ParameterArgs struct {
	PGVersion float32
	Env       EnvironmentName
	TotalRAM  int
	MaxConn   int
	OSFamily  OSFamily
}

// ParameterRule : Defines a functions who compute a rule for the parameter
type ParameterRule func(ParameterArgs) DatabaseParameter

func computeParameter(args ParameterArgs, f ParameterRule) (int, DatabaseParameter, error) {
	param := f(args)

	strRule := param.Rule

	if strings.Contains(strRule, "TOTAL_RAM") && args.TotalRAM > 0 {
		strRule = strings.Replace(strRule, "TOTAL_RAM", strconv.Itoa(args.TotalRAM), -1)
	}

	if strings.Contains(strRule, "MAX_CONNECTIONS") && args.MaxConn > 0 {
		strRule = strings.Replace(strRule, "MAX_CONNECTIONS", strconv.Itoa(args.MaxConn), -1)
	}

	exp, err := parser.ParseExpr(strRule)
	if err != nil {
		return 0, DatabaseParameter{}, err
	}

	param.Value = fixValue(&param, eval(exp), args.PGVersion)

	return param.Value, param, nil
}
