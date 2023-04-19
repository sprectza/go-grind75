// Problem desc: https://leetcode.com/problems/implement-queue-using-stacks/

type Stack []int

func (s *Stack) Push(val int) {
    *s = append(*s, val)
} 

func (s *Stack) Pop() int {
    val := (*s)[len(*s)-1]
    *s = (*s)[:len(*s)-1]
    return val
}

func (s *Stack) IsEmpty() bool {
    return len(*s) == 0
}

type MyQueue struct {
    stack1, stack2 Stack
}

func Constructor() MyQueue {
    return MyQueue{}    
}


func (q *MyQueue) Push(x int)  {
    q.stack1.Push(x)
}


func (q *MyQueue) Pop() int {
    if q.stack2.IsEmpty() {
        for !q.stack1.IsEmpty() {
            q.stack2.Push(q.stack1.Pop())
        }
    }
    return q.stack2.Pop()
}


func (q *MyQueue) Peek() int {
    if q.stack2.IsEmpty() {
        for !q.stack1.IsEmpty() {
            q.stack2.Push(q.stack1.Pop())
        }
    }
    return q.stack2[len(q.stack2)-1]
}


func (q *MyQueue) Empty() bool {
    return q.stack1.IsEmpty() && q.stack2.IsEmpty()
}
