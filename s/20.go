package s

func IsValid(s string) bool {
	var stack []byte
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '[', '(', '{':
			stack = append(stack, s[i])
		case ']', ')', '}':
			if len(stack) == 0 ||
				(s[i] == ']' && stack[len(stack)-1] != '[') ||
				(s[i] == ')' && stack[len(stack)-1] != '(') ||
				(s[i] == '}' && stack[len(stack)-1] != '{') {
				return false
			}
			stack = stack[:len(stack)-1]
		default:
			return false
		}
	}
	return len(stack) == 0
}
