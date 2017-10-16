package rules

import (
	"go/parser"
	"strconv"
	"strings"
)

var workMem = DatabaseParameter{
	Name:     "work_mem",
	MaxValue: -1,
	Type:     BytesParameter}

// WorkMem : Computes a 'work_mem' GUC of postgresql.conf
func WorkMem(pgVersion float32, env EnvironmentName, totalRAM int, maxConn int) (int, DatabaseParameter, error) {
	setWorkMem(pgVersion, env)

	strRule := ""
	strRule = strings.Replace(effectiveCacheSize.Rule, "TOTAL_RAM", strconv.Itoa(totalRAM), -1)
	strRule = strings.Replace(strRule, "MAX_CONNECTIONS", strconv.Itoa(maxConn), -1)

	exp, err := parser.ParseExpr(strRule)
	if err != nil {
		return 0, DatabaseParameter{}, err
	}

	effectiveCacheSize.Value = fixValue(&effectiveCacheSize, eval(exp), pgVersion)

	return effectiveCacheSize.Value, effectiveCacheSize, nil
}

func setWorkMem(pgVersion float32, env EnvironmentName) {

	if pgVersion <= 9.3 {
		effectiveCacheSize.DefaultValue = 1 * MEGABYTE
	} else {
		effectiveCacheSize.DefaultValue = 4 * MEGABYTE
	}

	if env == WebEnvironment || env == OLTPEnvironment {
		effectiveCacheSize.Rule = "TOTAL_RAM / MAX_CONNECTIONS"
	} else if env == DWEnvironment || env == MixedEnvironment {
		effectiveCacheSize.Rule = "TOTAL_RAM / 2 / MAX_CONNECTIONS"
	} else {
		effectiveCacheSize.Rule = "TOTAL_RAM / 6 / MAX_CONNECTIONS"
	}
}
