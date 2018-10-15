package main

import "fmt"

func findRestaurant(list1 []string, list2 []string) []string {
	if list1 == nil || list2 == nil {
		return []string{}
	}
	if len(list1) == 0 || len(list2) == 0 {
		return []string{}
	}
	h := map[string][]int{}
	for i, v := range list1 {
		h[v] = []int{i}
	}
	for i, v := range list2 {
		if iList, ok := h[v]; ok {
			iList = append(iList, i)
			h[v] = iList
		}
	}
	minDepth := 2001
	minDepthRest := []string{}
	for k, v := range h {
		if len(v) == 2 {
			if v[0]+v[1] < minDepth {
				minDepth = v[0] + v[1]
				minDepthRest = []string{k}
			} else if v[0]+v[1] == minDepth {
				minDepthRest = append(minDepthRest, k)
			}
		}
	}
	return minDepthRest
}

func main() {
	l1 := []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
	l2 := []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}
	fmt.Println(findRestaurant(l1, l2))
}
