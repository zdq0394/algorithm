package hashtable

import (
	"testing"
)

func TestUncommonFromSentences(t *testing.T) {
	a := "this apple is sweet"
	b := "this apple is sour"
	if !equalArrayString(uncommonFromSentences(a, b), []string{"sweet", "sour"}) {
		t.Fail()
	}
}
