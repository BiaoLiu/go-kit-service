package lorem_service

import (
	"context"
	"fmt"
)

type Service interface {
	Lorem(ctx context.Context, requestType string, min int, max int) (string, error)
}

type LoremService struct {
}

func (LoremService) Lorem(ctx context.Context, requestType string, min int, max int) (string, error) {
	fmt.Println(requestType, min, max)
	return "ok", nil
}
