package strings

func SplitChars(line string) []string {
	list := make([]string, 0)
	for i := 0; i < len(line); i++ {
		list = append(list, string(line[i]))
	}
	return list
}

func SplitBytes(line string) []byte {
	list := make([]byte, 0)
	for i := 0; i < len(line); i++ {
		list = append(list, line[i])
	}
	return list
}

func SplitByChar(line string, char byte) []string {
	list := make([]string, 0)
	currentString := ""
	for i := 0; i < len(line); i++ {
		if line[i] == char {
			list = append(list, currentString)
		} else {
			currentString += string(line[i])
		}
	}
	return list
}

func Count(line string, char byte) int {
	count := 0
	for i := 0; i < len(line); i++ {
		if line[i] == char {
			count += 1
		}
	}
	return count
}