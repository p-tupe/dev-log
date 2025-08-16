// This code show how to create, update, and read a context
package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "abc", "xyz")
	f(ctx)
}

func f(ctx context.Context) {
	fmt.Println(ctx)
	fmt.Println(ctx.Err())
	fmt.Println(ctx.Deadline())
	fmt.Println(ctx.Done())
	fmt.Println(ctx.Value("abc"))
}
