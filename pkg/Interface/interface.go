package Interface

type Dispatcher interface {
	Send(mobile, code string) error
}
