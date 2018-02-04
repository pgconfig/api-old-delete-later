package params

// EffectiveCacheSize contains a 'effective_cache_size' GUC of postgresql.conf
var EffectiveCacheSize = Parameter{
	Name:     "effective_cache_size",
	maxValue: -1,
	// maxVersion: maxVersion,
	Type: BytesParameter,
	Doc: &Doc{
		DefaultValue: 3,
		DocURLSuffix: "runtime-config-query.html#GUC-EFFECTIVE-CACHE-SIZE",
		Abstract:     "This parameter does not allocate any resource, just tells to the query planner how much of the operating system cache are available to use. Remember that shared_buffers needs to smaller than 8GB, then the query planner will prefer read the disk because it will be on memory.",
	},
	computeFunc: computeEffectiveCacheSize,
}

func computeEffectiveCacheSize(p *Parameter, args *Input) (out interface{}, err error) {

	if err = validateArgs(p, args); err != nil {
		return
	}

	if args.PGVersion <= 9.2 {
		p.Doc.DefaultValue = 128 * MegaByte
	} else {
		p.Doc.DefaultValue = 4 * GigaByte
	}

	switch args.Env {
	case DesktopEnvironment:
		out = args.TotalRAM / 4
	default:
		out = args.TotalRAM / 4 * 3
	}

	return
}
