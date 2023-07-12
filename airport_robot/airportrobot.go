package airportrobot

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

type Italian struct{}
type Portuguese struct{}

func SayHello(name string, greeter Greeter) string {
	return fmt.Sprintf("I can speak %s : %s", greeter.LanguageName(), greeter.Greet(name))
}

func (ital Italian) LanguageName() string {
	return "Italian"
}

func (ital Italian) Greet(name string) string {
	return fmt.Sprintf("Ciao %s!", name)
}

func (porto Portuguese) LanguageName() string {
	return "Portuguese"
}

func (porto Portuguese) Greet(name string) string {
	return fmt.Sprintf("Ola_Ola %s!", name)
}
