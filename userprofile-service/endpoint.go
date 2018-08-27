package userprofile_service

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

//func (ps *ProfileService) PostProfile(ctx *context.Context,request interface{}) (response interface{},err error){
//	p:=request.(Profile)
//
//}

func MakePostProfileEndPoint(ps ProfileService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		p := request.(Profile)
		err = ps.PostProfile(ctx, p)
		return nil, err
	}
}

func makeGetProfileEndPoint(ps ProfileService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		id := request.(string)
		profile, err := ps.GetProfile(ctx, id)
		return profile, err
	}
}
