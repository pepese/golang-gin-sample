package usecase

type Ping struct{}

func (*Ping) GET() string {
	return "pong"
}
