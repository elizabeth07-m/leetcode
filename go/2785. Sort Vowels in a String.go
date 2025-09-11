

func sortVowels(s string) string {
	// Define vowels
	vowelSet := "aeiouAEIOU"


	vowels := []rune{}
	for _, ch := range s {
		if strings.ContainsRune(vowelSet, ch) {
			vowels = append(vowels, ch)
		}
	}

	sort.Slice(vowels, func(i, j int) bool {
		return vowels[i] < vowels[j]
	})


	result := []rune(s)
	idx := 0
	for i, ch := range result {
		if strings.ContainsRune(vowelSet, ch) {
			result[i] = vowels[idx]
			idx++
		}
	}

	return string(result)
}


