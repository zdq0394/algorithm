package hashtable

import "testing"

func TestSubdomainVisits(t *testing.T) {
	ret := subdomainVisitsAsMap([]string{"9001 discuss.leetcode.com"})
	if !equalMap(ret, map[string]int{"com": 9001, "leetcode.com": 9001, "discuss.leetcode.com": 9001}) {
		t.Fail()
	}
}
