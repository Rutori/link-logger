package interfaces

type APIError interface {
	Code() int
	Error() string
}
