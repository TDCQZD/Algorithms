package src

// Stack 是用于存放 int 的 队列
type Queue struct {
	nums []int
}

// NewQueue 返回
func NewQueue() *Queue {
	return &Queue{nums: []int{}}
}

// Push 把 n 放入 队列
func (s *Queue) Push(n int) {
	s.nums = append(s.nums, n)
}

// Pop 从 s 中取出最后放入 队列 的值
func (s *Queue) Pop() int {
	res := s.nums[0]
	s.nums = s.nums[1:]
	return res
}

// Len 返回 s 的长度
func (s *Queue) Len() int {
	return len(s.nums)
}

// IsEmpty 反馈 s 是否为空
func (s *Queue) IsEmpty() bool {
	return s.Len() == 0
}
