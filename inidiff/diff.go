package inidiff

import (
	"github.com/mikelikespie/configparser-go/configparser"
)

func compareStrings(source string, destination string) ([]ComparisonRecord, error) {

	records := make([]ComparisonRecord, 0)

	confSrc, errSrc := configparser.ParseString(source)
	confDst, errDst := configparser.ParseString(destination)

	if errSrc != nil {
		return nil, errSrc
	}
	if errDst != nil {
		return nil, errDst
	}

	// assuming same sections, same options exist on both sides... for now
	for section, record := range confSrc {
		for opt, _ := range record {
			old := confSrc[section][opt]
			new := confDst[section][opt]
			if old != new {
				records = append(records, ComparisonRecord{section, opt, old, new})
			}
		}
	}

	return records, nil

}

type ComparisonRecord struct {
	section  string
	option   string
	oldValue string
	newValue string
}
