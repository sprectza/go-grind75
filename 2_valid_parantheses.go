// Problem link: https://leetcode.com/problems/valid-parentheses/description/

func isValid(s string) bool {
    stack := make([]rune, 0) 

    for _, ch := range s {
        switch ch {
        case '(', '{', '[':
            stack = append(stack, ch)
        case ')', '}', ']':
            if len(stack) == 0 {
                return false 
            }

            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]

            if  (top == '(' && ch != ')') || 
                (top == '[' && ch != ']') ||
                (top == '{' && ch != '}') { 
                return false 
            }
        }
    }

    return len(stack) == 0 
}
