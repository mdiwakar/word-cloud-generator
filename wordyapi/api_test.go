package wordyapi

import (
	"encoding/json"
	"testing"
)

func TestParseText(t *testing.T) {

	// Send a string with just two words
	s := TextToParse{Title: "Sample", Text: "test test test"}

	out := ParseText(s)

	var v map[string]interface{}

	err := json.Unmarshal(out, &v)
	if err != nil {
		t.Error("Got this error:", err)
	}

	// test to make sure it counted number of occurances
	if v["test"].(float64) != 3 {
		t.Error("Expected 3 occurances, but got:", v)
	}
}

func TestParseTextMixedCase(t *testing.T) {

        // Send a string with just two words
        s := TextToParse{Title: "Sample", Text: "test test test Test"}

        out := ParseText(s)

        var v map[string]interface{}

        err := json.Unmarshal(out, &v)
        if err != nil {
                t.Error("Got this error:", err)
        }

        // test to make sure it counted number of occurances
        if v["test"].(float64) != 4 {
                t.Error("Expected 4 occurances, but got:", v)
        }
}

func TestParseNumericalText(t *testing.T) {

        // Send a string with just two words
        s := TextToParse{Title: "Sample", Text: "123"}

        out := ParseText(s)

        var v map[string]interface{}

        err := json.Unmarshal(out, &v)
        if err != nil {
                t.Error("Got this error:", err)
        }

        // test to make sure it counted number of occurances
        if v["123"].(float64) != 1 {
                t.Error("Expected 4 occurances, but got:", v)
        }
}


func TestParseNumericalTextString(t *testing.T) {

        // Send a string with just two words
        s := TextToParse{Title: "Sample", Text: "123 123 123"}

        out := ParseText(s)

        var v map[string]interface{}

        err := json.Unmarshal(out, &v)
        if err != nil {
                t.Error("Got this error:", err)
        }

        // test to make sure it counted number of occurances
        if v["123"].(float64) != 3 {
                t.Error("Expected 4 occurances, but got:", v)
        }
}


