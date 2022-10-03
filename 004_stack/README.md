# Stack

Go 에서의 Stack은 Slice를 활용하면 Stack를 구현할 수 있습니다. 

```go
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(Eval("5"))
	fmt.Println(Eval("1 + 2"))
	fmt.Println(Eval("1 - 2 + 3"))
	fmt.Println(Eval("3 * ( 3 + 1 * 3 ) / 2"))
	fmt.Println(Eval("3 * ( ( 3 + 1 ) * 3 ) / 2"))
}

func Eval(expr string) int {
	var ops []string
	var nums []int
	pop := func() int {
		last := nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		return last
	}
	reduce := func(higher string) {
		for len(ops) > 0 {
			op := ops[len(ops)-1]
			if strings.Index(higher, op) < 0 {
				// 목록에 없는 연산자 이므로 종료
				return
			}
			ops = ops[:len(ops)-1]
			if op == "(" {
				// 괄호를 제거하였으므로 종료
				return
			}
			b, a := pop(), pop()
			switch op {
			case "+":
				nums = append(nums, a+b)
			case "-":
				nums = append(nums, a+b)
			case "*":
				nums = append(nums, a*b)
			case "/":
				nums = append(nums, a/b)
			}
		}
	}
	for _, token := range strings.Split(expr, " ") {
		switch token {
		case "(":
			ops = append(ops, token)
		case "+", "-":
			// 덧셈과 뺄셈 이상의 우선 순위를 가진 사칙 연산 적용
			reduce("+-*/")
			ops = append(ops, token)
		case "*", "/":
			reduce("*/")
			ops = append(ops, token)
		case ")":
			// 닫는 괄호는 여는 괄호까지 계산하고 제거
			reduce("+-*/(")
		default:
			num, _ := strconv.Atoi(token)
			nums = append(nums, num)
		}
	}
	reduce("+-*/")
	return nums[0]
}
```