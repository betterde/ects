package notify

type (
	Notify interface {
		Send() error
	}
)