package account

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
	"net/http"
)

type EndPoints struct {
	CreateUser endpoint.Endpoint
	GetUser endpoint.Endpoint
}

func MakeEndPoint(s Service) EndPoints {
	return EndPoints{
		CreateUser: makeCreateUserEndPoint(s),
		GetUser: makeGetUserEndPoint(s),

	}
}
func makeCreateUserEndPoint(s Service) endpoint.Endpoint {

	return func(ctx context.Context,request interface{})(interface{},error){
       req:=request.(CreateUserRequest)
       ok,err:=s.CreateUser(ctx,req.Email,req.Password)
       return CreateUserResponse{Ok: ok},err
	}
}
func makeGetUserEndPoint(s Service) endpoint.Endpoint{
	return func(ctx context.Context,request interface{})(interface{},error){
      req:=request.(GetUserRequest)
      email,err:=s.GetUser(ctx,req.Id)
      return GetUserResponse{
      	Email: email,

	  },err
	}
}

func decodeEmailReq(ctx context.Context,r *http.Request)(interface{},error){
	var req GetUserRequest
	vars:=mux.Vars(r)
	req=GetUserRequest{
		Id: vars["id"],
	}
	return req,nil
}
