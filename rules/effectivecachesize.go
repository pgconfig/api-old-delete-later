package rules

// EffectiveCacheSize : Computes a 'effective_cache_size' GUC of postgresql.conf
func EffectiveCacheSize(args ParameterArgs) (int, DatabaseParameter, error) {
	return computeParameter(args, effectiveCacheSizeRules)
}

func effectiveCacheSizeRules(args ParameterArgs) DatabaseParameter {

	var newParam = DatabaseParameter{
		Name:     "effective_cache_size",
		MaxValue: -1,
		Type:     BytesParameter}

	if args.PGVersion <= 9.2 {
		newParam.DefaultValue = 128 * MEGABYTE
	} else {
		newParam.DefaultValue = 4 * GIGABYTE
	}

	if args.Env == DesktopEnvironment {
		newParam.Rule = "TOTAL_RAM / 4"
	} else {
		newParam.Rule = "TOTAL_RAM / 4 * 3"
	}

	return newParam
}
