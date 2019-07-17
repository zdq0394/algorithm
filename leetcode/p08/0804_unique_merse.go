package p08

func uniqueMorseRepresentations(words []string) int {
	table := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

	if words == nil || len(words) == 0 {
		return 0
	}
	h := map[string]int{}
	for _, word := range words {
		morse := ""
		for _, l := range word {
			morse += table[l-'a']
		}
		h[morse]++
	}
	return len(h)
}
