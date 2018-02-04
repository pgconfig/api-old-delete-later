package params

import (
	"fmt"
)

// Input contains a argument to compute a parameter
type Input struct {
	PGVersion float32
	Env       EnvironmentName
	TotalRAM  int
	MaxConn   int
	OSFamily  OSFamily
	OSArch    OSArch
	HideDoc   bool
}

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

// OSArch represents the OS arch
type OSArch string

const (
	// Arch32Bits represents the 'i686' arch
	Arch32Bits OSArch = "i686"

	// Arch64Bits represents the 'x86-64' arch
	Arch64Bits OSArch = "x86-64"
)

type paramCategory string

const (

	// MemoryRelatedCategory : Defines a memory category
	MemoryRelatedCategory paramCategory = "Memory Configuration"

	// ChekPointRelatedCategory : Defines a checkpoint category
	ChekPointRelatedCategory paramCategory = "Checkpoint Related Configuration"

	// NetworkRelatedCategory : Defines a network category
	NetworkRelatedCategory paramCategory = "Network Related Configuration"
)

type paramType string

const (
	// BytesParameter : This value can be formated as bytes
	BytesParameter paramType = "Bytes"

	// NumericParameter : This value should be displayed as number
	NumericParameter paramType = "Decimal"

	// StringParameter : This value should be displayed as string
	StringParameter paramType = "String"
)

type recommendation struct {
	Title string
	URL   string
}

// Doc contains documentation related to a parameter
type Doc struct {
	Abstract     string            `json:"abstract"`
	Articles     map[string]string `json:"recomendations"`
	DocURLSuffix string            `json:"url"`
	DefaultValue int               `json:"default_value"`
}

const (
	minimumVer float32 = 8.0
	defaultVer float32 = 10.0
)

type paramCompute func(*Parameter, *Input) (interface{}, error)

// Parameter contains data about a PostgreSQL parameter
type Parameter struct {
	Name        string `json:"name"`
	maxValue    int
	maxVersion  float32
	minVersion  float32
	Value       interface{} `json:"config_value"`
	Type        paramType   `json:"format"`
	Doc         *Doc        `json:"documentation,omitempty"`
	computeFunc paramCompute
}

func validateArgs(p *Parameter, args *Input) (err error) {
	setDefaults(p, args)

	if args.PGVersion <= minimumVer || args.PGVersion > p.maxVersion || args.PGVersion < p.minVersion {
		err = fmt.Errorf("Version %.1f unsupported for %s", args.PGVersion, p.Name)
		return
	}

	return
}

func setDefaults(p *Parameter, args *Input) {

	if args == nil {
		args = &Input{}
	}

	if p.maxVersion == 0.0 {
		p.maxVersion = defaultVer
	}
	if args.PGVersion == 0.0 {
		args.PGVersion = defaultVer
	}
}

// Compute calculates a parameter
func (p *Parameter) Compute(args Input) (err error) {

	p.Value, err = p.computeFunc(p, &args)

	if args.HideDoc {
		p.Doc = nil
	}

	return
}
