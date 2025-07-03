package logic

import "microservice/internal/database"

func sayHello(thing string) (greeting string) {
	return "Hello " + thing
}

func SayHello(thing string) (greeting string) {
	return ""
	// 	if thing == "" {
	// 		return "Hello"
	// 	}

	// return "Hello " + thing
}

type Hello struct {
	UserDataAccessor database.UserDataAccessor
}

func (h Hello) Hello() {

}
