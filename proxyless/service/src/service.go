package src

const (
	PodNameKey = "POD_NAME"
)

type TestService interface {
	Run() error
}
