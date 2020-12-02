package main

import (
	"context"
	"fmt"

	"github.com/tessellator/fnrun"
)

func Sink(ctx context.Context, result *fnrun.Result) error {
	_, err := fmt.Printf("The result is: %s\n", string(result.Data))
	return err
}
