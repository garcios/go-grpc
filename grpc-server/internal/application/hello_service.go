package application

type HelloService struct{}

func (a *HelloService) Greet(name string) string {
	return "Hello " + name
}
