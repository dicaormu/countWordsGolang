package wordCore

import "strings"

const delim = "?!.;,*"
const space = " "

func isDelim(c string) bool {
	/*if strings.Contains(delim, c) {
		return true
	}
	return false*/
	return strings.Contains(delim, c)
}

func cleanDelimiter(input string) string {
	size := len(input)
	temp := ""
	var prevChar string

	for i := 0; i < size; i++ {
		str := string(input[i]) // convert to string for easier operation
		if (str == space && prevChar != space) || !isDelim(str) {
			temp += str
			prevChar = str
		} else if prevChar != space && isDelim(str) {
			temp += ""
		}
	}
	return temp
}

func SingleWordCount(s string, uniqueSensitive bool) map[string]int {
	strSlice := strings.Fields(cleanDelimiter(s))
	result := make(map[string]int)

	if uniqueSensitive {
		for _, str := range strSlice {
			result[str]++
		}

	} else {
		for _, str := range strSlice {
			result[strings.ToLower(str)]++
		}
	}
	return result
}

func TotalWordCount(s string) int {
	strSlice := strings.Fields(cleanDelimiter(s))
	return len(strSlice)
}

/*func main() {

}*/
