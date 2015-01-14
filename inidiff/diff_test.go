package inidiff

import (
	"reflect"
	"testing"
)

var base_sample_string = `
[section1]
option1=value1
option2=value2

[section2]
option1=value12
option2=value2
option3=value3
`

var modif_order_string = `
[section2]
option3=value3
option1=value12
option2=value2

[section1]
option2=value2
option1=value1
`

var modif_one_value = `
[section1]
option1=value12
option2=value2

[section2]
option1=value12
option2=value2
option3=value3
`

var modif_two_values = `
[section1]
option1=value12
option2=value2

[section2]
option1=value12
option2=value22
option3=value3
`

var fixtures = []struct {
	inputString string
	expected    []ComparisonRecord
	err         error
}{
	{
		modif_order_string,
		[]ComparisonRecord{},
		nil,
	},
	{
		modif_one_value,
		[]ComparisonRecord{{"section1", "option1", "value1", "value12"}},
		nil,
	},
	{
		modif_two_values,
		[]ComparisonRecord{
			{"section1", "option1", "value1", "value12"},
			{"section2", "option2", "value2", "value22"},
		},
		nil,
	},
}

func TestCommand(t *testing.T) {
	for _, test := range fixtures {
		result, err := compareStrings(base_sample_string, test.inputString)

		// check for errors
		if err != test.err {
			t.Error("error err")
		}

		// compare results
		if !reflect.DeepEqual(result, test.expected) {
			t.Error("Expected ", test.expected, "got", result)
		}
	}
}
