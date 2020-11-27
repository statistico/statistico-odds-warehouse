package market

func isSupportedOverUnderMarket(s string) bool {
	markets := []string{
		"OVER_UNDER_05",
		"OVER_UNDER_15",
		"OVER_UNDER_25",
		"OVER_UNDER_35",
		"OVER_UNDER_45",
		"OVER_UNDER_25_CARDS",
		"OVER_UNDER_35_CARDS",
		"OVER_UNDER_45_CARDS",
		"OVER_UNDER_65_CARDS",
		"OVER_UNDER_55_CORNR",
		"OVER_UNDER_85_CORNR",
		"OVER_UNDER_105_CORNR",
		"OVER_UNDER_135_CORNR",
	}

	for _, m := range markets {
		if m == s {
			return true
		}
	}

	return false
}
