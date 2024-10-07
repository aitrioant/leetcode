package minLength

func minLength(s string) int {
	for i := 0; i < len(s)-1; i++ {
		current := s[i]
		next := s[i+1]
		if current == 'A' && next == 'B' || current == 'C' && next == 'D' {
			s = removeCharAt(s, i)
			i = -1
		}
	}

	return len(s)
}

func removeCharAt(s string, i int) string {
	return s[:i] + s[i+2:]
}
