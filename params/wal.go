package params

// MinWalSize contains a 'min_wal_size' GUC of postgresql.conf
var MinWalSize = Parameter{
	Name:       "min_wal_size",
	minVersion: 9.5,
	maxValue:   -1,
	Type:       BytesParameter,
	Doc: &Doc{
		DefaultValue: 80 * MegaByte,
		DocURLSuffix: "runtime-config-wal.html#GUC-MIN-WAL-SIZE",
		Abstract:     "This parameter defines the minimum size of the pg_xlog directory. pgx_log directory contains the WAL files.",
		Articles: map[string]string{
			"Configuration changes in 9.5: transaction log size": "http://www.databasesoup.com/2016/01/configuration-changes-in-95-transaction.html",
		},
	},
	computeFunc: computeMinWalSize,
}

func computeMinWalSize(p *Parameter, args *Input) (out interface{}, err error) {
	if err = validateArgs(p, args); err != nil {
		return
	}

	switch args.Env {
	case WebEnvironment, MixedEnvironment:
		out = 512 * MegaByte
	case DesktopEnvironment:
		out = 256 * MegaByte
	case OLTPEnvironment:
		out = 1 * GigaByte
	default:
		out = 2 * GigaByte
	}

	return
}
