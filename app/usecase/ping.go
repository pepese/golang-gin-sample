package usecase

import "context"

/*
PingUsecase interface contains Usecase methods.
*/
type PingUsecaser interface {
	Say(c context.Context, m string) (string, error)
}

/*
PingUsecase model.
*/
type PingUsecase struct{}

func (uc PingUsecase) Say(c context.Context, m string) (string, error) {
	return m, nil
}

/*
NewPingUsecase returns *pingUsecase.
*/
/*func NewPingUsecase() *pingUsecase {
	return &pingUsecase{}
}*/
