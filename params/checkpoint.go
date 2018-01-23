package params

// CheckPointSegments contains a 'checkpoint_segments' GUC of postgresql.conf
var (
	maxVersion         float32 = 9.4
	CheckPointSegments         = Parameter{
		Name:       "checkpoint_segments",
		maxValue:   -1,
		maxVersion: maxVersion,
		Type:       NumericParameter,
		Doc: Doc{
			DefaultValue: 3,
			DocURLSuffix: "runtime-config-wal.html#GUC-CHECKPOINT-SEGMENTS",
			Abstract:     "This parameter defines how much WAL files can be stored before a automatic CHECKPOINT. All files are stored in the pg_xlog directory.",
			Articles: []map[string]string{
				map[string]string{"WRITE AHEAD LOG + UNDERSTANDING POSTGRESQL.CONF: CHECKPOINT_SEGMENTS, CHECKPOINT_TIMEOUT and CHECKPOINT_WARNING": "https://www.depesz.com/2011/07/14/write-ahead-log-understanding-postgresql-conf-checkpoint_segments-checkpoint_timeout-checkpoint_warning/"},
			},
		},
		computeFunc: computeCheckPointSegments,
	}
)

func computeCheckPointSegments(args Input) (out interface{}, err error) {

	if err = validateArgs(args); err != nil {
		return
	}

	switch args.Env {
	case WebEnvironment, MixedEnvironment:
		out = 32
	case OLTPEnvironment:
		out = 96
	case DWEnvironment:
		out = 256
	default:
		out = 16
	}

	return
}
