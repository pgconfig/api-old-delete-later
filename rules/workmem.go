package rules

// WorkMem : Computes a 'work_mem' GUC of postgresql.conf
func WorkMem(args ParameterArgs) (interface{}, DatabaseParameter, error) {
	return computeParameter(args, setWorkMem)
}

func setWorkMem(args ParameterArgs) DatabaseParameter {

	var workMem = DatabaseParameter{
		Name:         "work_mem",
		MaxValue:     -1,
		Type:         BytesParameter,
		Category:     MemoryRelatedCategory,
		DocURLSuffix: "runtime-config-resource.html#GUC-WORK-MEM",
		Abstract:     "This parameter defines how much a work_mem buffer can allocate. Each query can open many work_mem buffers when execute (normally one by subquery) if it uses any sort (or aggregate) operation. When work_mem its too small a temp file is created.",
		Articles: []ArticleRecommendation{
			{Title: "Understaning postgresql.conf: WORK_MEM", URL: "https://www.depesz.com/2011/07/03/understanding-postgresql-conf-work_mem/"},
		},
	}

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

// MaintenanceWorkMem : Computes a 'maintenance_work_mem' GUC of postgresql.conf
func MaintenanceWorkMem(args ParameterArgs) (interface{}, DatabaseParameter, error) {
	return computeParameter(args, setMaintenanceWorkMem)
}

func setMaintenanceWorkMem(args ParameterArgs) DatabaseParameter {

	newValue := DatabaseParameter{
		Name:     "maintenance_work_mem",
		MaxValue: 2 * GIGABYTE,
		Type:     BytesParameter,
		Category: MemoryRelatedCategory,
		Abstract: "This parameter defines how much a maintenance operation (ALTER TABLE, VACUUM, REINDEX, AutoVACUUM worker, etc) buffer can use.",
	}

	if args.PGVersion <= 9.3 {
		newValue.DefaultValue = 16 * MEGABYTE
	} else {
		newValue.DefaultValue = 64 * MEGABYTE
	}

	if args.Env == DWEnvironment {
		newValue.Rule = "TOTAL_RAM / 8"
	} else {
		newValue.Rule = "TOTAL_RAM / 16"
	}

	return newValue
}
