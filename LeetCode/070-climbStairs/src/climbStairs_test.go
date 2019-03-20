package src

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	res := climbStairs(3)
	fmt.Println(res)
	res = climbStairs2(3)
	fmt.Println(res)
	res = climbStairs3(3)
	fmt.Println(res)
}
