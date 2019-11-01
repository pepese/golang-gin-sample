package usecase

/*
PingUsecase interface contains Usecase methods.
*/
type PingUsecase interface {
	Say()
}

/*
PingUsecase model.
*/
type pingUsecase struct {
	Message string `json:"message"`
}

func (uc *pingUsecase) Say(message string) {
	uc.Message = message
}

/*
NewPingUsecase returns *pingUsecase.
*/
func NewPingUsecase() *pingUsecase {
	return &pingUsecase{}
}
