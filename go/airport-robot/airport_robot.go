package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.


type Greeter interface {
	LanguageName() string;
	Greet(name string) string; 
}

type German struct{}
func (i German) LanguageName() string {
	return "German"
}

func (i German) Greet(name string) string {
	return fmt.Sprintf("Hallo %s!", name)
}

type Italian struct{}
func (i Italian) LanguageName() string {
	return "Italian"
}

func (i Italian) Greet(name string) string {
	return fmt.Sprintf("Ciao %s!", name)
}

type Portuguese struct{}
func (i Portuguese) LanguageName() string {
	return "Portuguese"
}

func (i Portuguese) Greet(name string) string {
	return fmt.Sprintf("Olá %s!", name)
}

func SayHello(name string, greeter Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", greeter.LanguageName(), greeter.Greet(name))
}