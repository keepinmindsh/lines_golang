package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	ctx = context.WithValue(ctx, "current_user", "shjeong")

	myFunc(ctx)
}

func myFunc(ctx context.Context) {
	var currentUser any

	// todo 코드가 제대로 작성되지 않아 작성 필요 ( 22-10-24 )
	// 컨텍스트에서 값을 가져옴
	//if v := ctx.Value("current_user"); v != nil {
	//
	//	// 타입 확인(type assertion)
	//	u := v
	//	if !ok {
	//		return errors.New("Not authorized")
	//	}
	//	currentUser = u
	//} else {
	//	return errors.New("Not authorized")
	//}

	// currentUser를 사용하여 로직 처리
	fmt.Println(currentUser)
}
