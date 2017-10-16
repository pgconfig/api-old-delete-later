package rules

import (
	"go/parser"
	"strconv"
	"strings"
)

var effectiveCacheSize = DatabaseParameter{
	Name:     "effective_cache_size",
	MaxValue: -1,
	Type:     BytesParameter}

// EffectiveCacheSize : Computes a 'effective_cache_size' GUC of postgresql.conf
func EffectiveCacheSize(pgVersion float32, env EnvironmentName, totalRAM int) (int, DatabaseParameter, error) {
	setEffectiveCacheSize(pgVersion, env)

	strRule := strings.Replace(effectiveCacheSize.Rule, "TOTAL_RAM", strconv.Itoa(totalRAM), -1)

	exp, err := parser.ParseExpr(strRule)
	if err != nil {
		return 0, DatabaseParameter{}, err
	}

	effectiveCacheSize.Value = fixValue(&effectiveCacheSize, eval(exp), pgVersion)

	return effectiveCacheSize.Value, effectiveCacheSize, nil
}

func setEffectiveCacheSize(pgVersion float32, env EnvironmentName) {
	effectiveCacheSize.MaxValue = -1

	if pgVersion <= 9.2 {
		effectiveCacheSize.DefaultValue = 128 * MEGABYTE
	} else {
		effectiveCacheSize.DefaultValue = 4 * GIGABYTE
	}

	if env == DesktopEnvironment {
		effectiveCacheSize.Rule = "TOTAL_RAM / 4"
	} else {
		effectiveCacheSize.Rule = "TOTAL_RAM / 4 * 3"
	}
}
