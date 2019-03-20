package src

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {

	nums := []int{2, 7, 11, 15}
	target := 9
	arr := TwoSum(nums, target)
	println(arr[0], arr[1])

}

type argument struct {
	numberArray []int
	target      int
}

type answer struct {
	numberArray []int
}

type example struct {
	arg argument
	ans answer
}

func TestOk(t *testing.T) {
	examples := []example{
		{
			arg: argument{
				numberArray: []int{3, 2, 4},
				target:      6,
			},
			ans: answer{
				numberArray: []int{1, 2},
			},
		}, {
			arg: argument{
				numberArray: []int{3, 2, 4},
				target:      8,
			},
			ans: answer{
				numberArray: nil,
			},
		},
	}

	ast := assert.New(t)
	for _, exam := range examples {
		arg, ans := exam.arg, exam.ans
		ast.Equal(ans.numberArray, TwoSum(arg.numberArray, arg.target),
			"%v %v", arg.numberArray, arg.target)
	}

}
