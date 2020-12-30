package account

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gofrs/uuid"
)

type service struct {
	repository Repository
	logger log.Logger
}

func (s service) CreateUser(ctx context.Context, email string, password string) (string, error) {
	logger:=log.With(s.logger,"method","CreateUser")

	uid,_ :=uuid.NewV4()
	id:=uid.String()
	user :=User{
		ID:id,
		Email: email,
		Password: password,
	}

    if err:=s.repository.CreateUser(ctx,user);err!=nil{
		_ = level.Error(logger).Log("err", err)
    	return "",err
	}
	_ = logger.Log("create user", id)
    return "Success",nil
}

func (s service) GetUser(ctx context.Context, id string) (string, error) {
	logger:=log.With(s.logger,"method","GetUser")
	email,err:=s.repository.GetUser(ctx,id)
	if err !=nil{
		_ = level.Error(logger).Log("err", err)
		return "",err
	}
	_ = logger.Log("Get User", id)
	return email,nil

}

func NewService(rep Repository,logger log.Logger) Service {
	return service{rep,logger}
}
