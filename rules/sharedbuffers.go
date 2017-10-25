package rules

// SharedBuffers : Computes a 'shared_buffers' GUC of postgresql.conf
func SharedBuffers(args ParameterArgs) (int, DatabaseParameter, error) {
	return computeParameter(args, setSharedBuffers)
}

func setSharedBuffers(args ParameterArgs) DatabaseParameter {

	var sharedBuffers = DatabaseParameter{
		Name:         "shared_buffers",
		MaxValue:     8 * GIGABYTE,
		Type:         BytesParameter,
		Category:     MemoryRelatedCategory,
		DocURLSuffix: "runtime-config-resource.html#GUC-SHARED-BUFFERS",
		Abstract:     "This parameter allocate memory slots, used by all process. Mainly works as the disk cache and its similar to oracle's SGA buffer.",
		Articles: []ArticleRecommendation{
			{Title: "Tuning Your PostgreSQL Server", URL: "https://wiki.postgresql.org/wiki/Tuning_Your_PostgreSQL_Server#shared_buffers"},
			{Title: "Tuning shared_buffers and wal_buffers", URL: "http://rhaas.blogspot.com.br/2012/03/tuning-sharedbuffers-and-walbuffers.html"},
		},
	}

	if args.OSFamily == WindowsOS && args.PGVersion <= 9.6 {
		sharedBuffers.MaxValue = 512 * MEGABYTE
	}

	if args.PGVersion <= 9.2 {
		sharedBuffers.DefaultValue = 32 * MEGABYTE
	} else {
		sharedBuffers.DefaultValue = 128 * MEGABYTE
	}

	if args.Env == DesktopEnvironment {
		sharedBuffers.Rule = "TOTAL_RAM / 16"
	} else {
		sharedBuffers.Rule = "TOTAL_RAM / 4"
	}

	return sharedBuffers
}
