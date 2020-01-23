package practice_log

// Contains the accepted 3-letter abbreviations for rudiments
var accepted_abbrev = []string{
	"ssr",
	"ss4",
	"5sr",
	"13sr",
	"3sr",
	"6sr",
	"17sr",
	"tpd",
	"ss7",
	"11sr",
	"spd",
	"dpd",
	"7sr",
	"9sr",
	"10sr",
	"15sr",
	"spdd",
}

// Determines if the abbrev is one that is supported or not (or if it even exists...)
func ValidAbbreviation(abbrev string) bool {
	for _, tla := range accepted_abbrev {
		if tla == abbrev {
			return true
		}
	}
	return false
}

