package text_test

import (
	"testing"
	"reflect"

	"github.com/iuryfukuda/grapnel/text"
)

type testToWords struct {
	raw string
	want map[string]int
}

var allTestsToWords = []testToWords{
	testToWords{
		raw: "abacate tomate alface abacate",
		want: map[string]int{
			"abacate": 2,
			"tomate": 1,
			"alface": 1,
		},
	},
}

func TestToWords(t *testing.T) {
	for _, tt := range allTestsToWords {
		got := text.ToWords(tt.raw)
		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("want [%q] got [%q]", tt.want, got)
		}
	}
}
