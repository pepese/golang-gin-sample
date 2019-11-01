package usecase

/*
PingUsecase interface contains Usecase methods.
*/
type PingUsecase interface {
	GET() string
}

type pingUsecase struct {
	response string
}

func (uc *pingUsecase) GET() string {
	return uc.response
}

/*
NewPingUsecase returns *pingUsecase.
*/
func NewPingUsecase() *pingUsecase {
	return &pingUsecase{response: "pong"}
}
