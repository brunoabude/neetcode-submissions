func isValid(s string) bool {
    stack := make([]rune, len(s))
	stack_top := -1
	
	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack_top++
			stack[stack_top] = char
		} else {
			matching := rune('0')

			switch char {
				case ')':
				matching = '('
				case '}':
				matching = '{'
				case ']':
				matching = '['
			}

			if stack_top < 0 || stack[stack_top] != matching {
				return false
			} else {
				stack_top--
			}
		}
	} 

	return stack_top == -1
}
