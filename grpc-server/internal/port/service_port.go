package port

type HelloServicePort interface {
	Greet(name string) string
}
