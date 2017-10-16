package rules

import (
	"go/parser"
	"strconv"
	"strings"
)

var sharedBuffers = DatabaseParameter{
	Name:     "shared_buffers",
	MaxValue: 8 * GIGABYTE,
	Type:     BytesParameter}

// SharedBuffers : Computes a 'shared_buffers' GUC of postgresql.conf
func SharedBuffers(pgVersion float32, osFamily string, env EnvironmentName, totalRAM int) (int, DatabaseParameter, error) {

	setSharedBuffersRules(pgVersion, osFamily, env)

	strRule := strings.Replace(sharedBuffers.Rule, "TOTAL_RAM", strconv.Itoa(totalRAM), -1)

	exp, err := parser.ParseExpr(strRule)
	if err != nil {
		return 0, DatabaseParameter{}, err
	}

	sharedBuffers.Value = fixValue(&sharedBuffers, eval(exp), pgVersion)

	return sharedBuffers.Value, sharedBuffers, nil
}

func setSharedBuffersRules(pgVersion float32, osFamily string, env EnvironmentName) {

	sharedBuffers.MaxValue = 8 * GIGABYTE

	if osFamily == "Windows" && pgVersion <= 9.6 {
		sharedBuffers.MaxValue = 512 * MEGABYTE
	}

	if pgVersion <= 9.2 {
		sharedBuffers.DefaultValue = 32 * MEGABYTE
	} else {
		sharedBuffers.DefaultValue = 128 * MEGABYTE
	}

	if env == DesktopEnvironment {
		sharedBuffers.Rule = "TOTAL_RAM / 16"
	} else {
		sharedBuffers.Rule = "TOTAL_RAM / 4"
	}
}
