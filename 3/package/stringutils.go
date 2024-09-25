package stringutils

func Reverse(str string) string {
	ans := make([]byte, len(str))
	i := 0
	j := len(str) - 1

	for i < len(str) {
		ans[i] = str[j]
		i++
		j--
	}

	return string(ans)
}
