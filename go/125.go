func filter(c rune) rune {
	if 97 <= c && c <= 122 {
		return c
	}
	if 65 <= c && c <= 90 {
		return c + 32
	}
	if 48 <= c && c <= 57 {
		return c
	}

	return 0
}

func isPalindrome(s string) bool {
	real := make([]byte, 0)
	for _, c := range s {
		if a := filter(c); a > 0 {
			real = append(real, byte(a))
		}
	}

	reverse := make([]byte, len(real))
	for idx, length := 0, len(real); idx < length; idx++ {
		reverse[idx] = real[length-idx-1]
	}

	return string(real) == string(reverse)
}