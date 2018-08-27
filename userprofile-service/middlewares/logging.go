package middlewares

import (
	"github.com/go-kit/kit/endpoint"
	"context"
)

func LoggingMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		//custom code
		return next(ctx, request)
	}
}

func test(test string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			//custom code
			return next(ctx, request)
		}
	}
}

//func main(){
//	test("")()
//}
