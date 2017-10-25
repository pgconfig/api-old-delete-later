package rules

// CheckpointSegments : Computes a 'checkpoint_segments' GUC of postgresql.conf
func CheckpointSegments(args ParameterArgs) (int, DatabaseParameter, error) {
	return computeParameter(args, setCheckpointSegments)
}

func setCheckpointSegments(args ParameterArgs) DatabaseParameter {

	var newParam = DatabaseParameter{
		Name:         "checkpoint_segments",
		MaxValue:     -1,
		DefaultValue: 3,
		MaxVersion:   9.4,
		Type:         NumericParameter,
		Category:     ChekPointRelatedCategory,
		DocURLSuffix: "runtime-config-wal.html#GUC-CHECKPOINT-SEGMENTS",
		Abstract:     "This parameter defines how much WAL files can be stored before a automatic CHECKPOINT. All files are stored in the pg_xlog directory.",
		Articles: []ArticleRecommendation{
			ArticleRecommendation{
				Title: "WRITE AHEAD LOG + UNDERSTANDING POSTGRESQL.CONF: CHECKPOINT_SEGMENTS, CHECKPOINT_TIMEOUT and CHECKPOINT_WARNING",
				URL:   "https://www.depesz.com/2011/07/14/write-ahead-log-understanding-postgresql-conf-checkpoint_segments-checkpoint_timeout-checkpoint_warning/",
			},
		},
	}

	if args.Env == WebEnvironment || args.Env == MixedEnvironment {
		newParam.Rule = "32"
	} else if args.Env == OLTPEnvironment {
		newParam.Rule = "96"
	} else if args.Env == DWEnvironment {
		newParam.Rule = "256"
	} else {
		newParam.Rule = "16"
	}

	return newParam
}
