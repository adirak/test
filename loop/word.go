package loop

// GenWord to gen word
func GenWord() string {
	count := 1000
	result := ""
	for i := 0; i < count; i++ {
		result += "  Hello world!....\n"
	}

	return result
}
