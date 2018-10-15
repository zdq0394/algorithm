package main

import (
	"fmt"
	"strconv"
)

func subDomains(domain string) []string {
	result := make([]string, 0)
	for i := len(domain) - 1; i >= 0; i-- {
		if domain[i] == '.' {
			result = append(result, domain[i+1:])
		}
	}
	result = append(result, domain)
	return result
}
func subdomainVisits(cpdomains []string) []string {
	if cpdomains == nil || len(cpdomains) == 0 {
		return []string{}
	}
	result := make(map[string]int, 0)
	for _, visitsDomain := range cpdomains {
		spaceIndex := 0
		for i := 0; i < len(visitsDomain); i++ {
			if visitsDomain[i] == ' ' {
				spaceIndex = i
				if spaceIndex > 0 && spaceIndex < len(visitsDomain)-1 {
					visitsStr := visitsDomain[:spaceIndex]
					fullDomain := visitsDomain[spaceIndex+1:]
					visitsInt, _ := strconv.Atoi(visitsStr)
					t := subDomains(fullDomain)
					for _, d := range t {
						result[d] += visitsInt
					}

				}
			}
		}
	}
	ret := make([]string, 0)
	for k, v := range result {
		ret = append(ret, fmt.Sprintf("%d %s", v, k))
	}
	return ret
}

func main() {
	fmt.Println(subdomainVisits([]string{"9001 discuss.leetcode.com"}))
}
