package main

import (
	"context"
	"errors"
)

func main() {
	ctx := context.Background()

	ctx = context.WithValue(ctx, "current_user", "shjeong")

	myFunc(ctx)
}

func myFunc(ctx context.Context) {
	var currentUser string

	// 컨텍스트에서 값을 가져옴
	if v := ctx.Value("current_user"); v != nil {
		// 타입 확인(type assertion)
		u := string(v)
		if !ok {
			return errors.New("Not authorized")
		}
		currentUser = u
	} else {
		return errors.New("Not authorized")
	}

	// currentUser를 사용하여 로직 처리

	return nil
}
