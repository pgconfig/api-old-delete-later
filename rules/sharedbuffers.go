package rules

// SharedBuffers : Computes a 'shared_buffers' GUC of postgresql.conf
func SharedBuffers(args ParameterArgs) (int, DatabaseParameter, error) {
	return computeParameter(args, setSharedBuffers)
}

func setSharedBuffers(args ParameterArgs) DatabaseParameter {

	var sharedBuffers = DatabaseParameter{
		Name:     "shared_buffers",
		MaxValue: 8 * GIGABYTE,
		Type:     BytesParameter}

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
