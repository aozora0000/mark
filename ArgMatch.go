package main

func ArgMatch(args []string, text string) bool {
	for _, arg := range args {
		if text == arg {
			return true
		}
	}
	return false
}
