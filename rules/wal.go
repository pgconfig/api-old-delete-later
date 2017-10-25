package rules

import (
	"fmt"
)

// MinWalSize : Computes a 'min_wal_size' GUC of postgresql.conf
func MinWalSize(args ParameterArgs) (interface{}, DatabaseParameter, error) {
	return computeParameter(args, setMinWalSize)
}

func setMinWalSize(args ParameterArgs) DatabaseParameter {

	var newParam = DatabaseParameter{
		Name:         "min_wal_size",
		MaxValue:     -1,
		DefaultValue: 80 * MEGABYTE,
		MinVersion:   9.5,
		Type:         BytesParameter,
		Category:     ChekPointRelatedCategory,
		DocURLSuffix: "runtime-config-wal.html#GUC-MIN-WAL-SIZE",
		Abstract:     "This parameter defines the minimum size of the pg_xlog directory. pgx_log directory contains the WAL files.",
		Articles: []ArticleRecommendation{
			{
				Title: "Configuration changes in 9.5: transaction log size",
				URL:   "http://www.databasesoup.com/2016/01/configuration-changes-in-95-transaction.html",
			},
			{
				Title: "Configuration changes in 9.5: transaction log size",
				URL:   "http://www.databasesoup.com/2016/01/configuration-changes-in-95-transaction.html",
			},
		},
	}

	var ruleValue int

	if args.Env == WebEnvironment || args.Env == MixedEnvironment {
		ruleValue = 512 * MEGABYTE
	} else if args.Env == DesktopEnvironment {
		ruleValue = 256 * MEGABYTE
	} else if args.Env == OLTPEnvironment {
		ruleValue = 1 * GIGABYTE
	} else {
		ruleValue = 2 * GIGABYTE
	}

	newParam.Rule = fmt.Sprint(ruleValue)

	return newParam
}

// MaxWalSize : Computes a 'min_wal_size' GUC of postgresql.conf
func MaxWalSize(args ParameterArgs) (interface{}, DatabaseParameter, error) {
	return computeParameter(args, setMaxWalSize)
}

func setMaxWalSize(args ParameterArgs) DatabaseParameter {

	var newParam = DatabaseParameter{
		Name:         "max_wal_size",
		MaxValue:     -1,
		DefaultValue: 1 * GIGABYTE,
		MinVersion:   9.5,
		Type:         BytesParameter,
		Category:     ChekPointRelatedCategory,
		DocURLSuffix: "runtime-config-wal.html#GUC-MIN-WAL-SIZE",
		Abstract:     "This parameter defines the minimum size of the pg_xlog directory. pgx_log directory contains the WAL files.",
		Articles: []ArticleRecommendation{
			{
				Title: "Configuration changes in 9.5: transaction log size",
				URL:   "http://www.databasesoup.com/2016/01/configuration-changes-in-95-transaction.html",
			},
			{
				Title: "Configuration changes in 9.5: transaction log size",
				URL:   "http://www.databasesoup.com/2016/01/configuration-changes-in-95-transaction.html",
			},
		},
	}

	if args.PGVersion <= 9.5 {
		newParam.Abstract = "This parameter defines the maximum size of the pg_xlog directory. pg_xlog directory contains the WAL files."
	} else {
		newParam.Abstract = "This parameter defines the maximum size of the pg_wal directory. pg_wal directory contains the WAL files."
	}

	var ruleValue int

	if args.Env == WebEnvironment || args.Env == MixedEnvironment {
		ruleValue = 2 * GIGABYTE
	} else if args.Env == DWEnvironment {
		ruleValue = 6 * GIGABYTE
	} else if args.Env == OLTPEnvironment {
		ruleValue = 3 * GIGABYTE
	} else {
		ruleValue = 1 * GIGABYTE
	}

	newParam.Rule = fmt.Sprint(ruleValue)

	return newParam
}
