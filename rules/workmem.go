package rules

// WorkMem : Computes a 'work_mem' GUC of postgresql.conf
func WorkMem(args ParameterArgs) (int, DatabaseParameter, error) {
	return computeParameter(args, setWorkMem)
}

func setWorkMem(args ParameterArgs) DatabaseParameter {

	var workMem = DatabaseParameter{
		Name:     "work_mem",
		MaxValue: -1,
		Type:     BytesParameter}

	if args.PGVersion <= 9.3 {
		workMem.DefaultValue = 1 * MEGABYTE
	} else {
		workMem.DefaultValue = 4 * MEGABYTE
	}

	if args.Env == WebEnvironment || args.Env == OLTPEnvironment {
		workMem.Rule = "TOTAL_RAM / MAX_CONNECTIONS"
	} else if args.Env == DWEnvironment || args.Env == MixedEnvironment {
		workMem.Rule = "TOTAL_RAM / 2 / MAX_CONNECTIONS"
	} else {
		workMem.Rule = "TOTAL_RAM / 6 / MAX_CONNECTIONS"
	}

	return workMem
}
