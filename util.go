package main

func trimFirstRune(s string) string {
	for i := range s {
		if i > 0 {

			return s[i:]
		}
	}
	return ""
}
