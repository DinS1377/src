package output_animal

import (
	"fmt"
	"go_project/packages/interfaces"
)

func Output(user_animal string) {
	if user_animal == "лев" {
		animal := interfaces.Lion{}
		animal.MakeSound()
		animal.Move()
		animal.Feed()
		animal.Sleep()
		animal.Swim()
		animal.Hunt()

	}
	if user_animal == "кошка" {
		animal := interfaces.Cat{}
		animal.MakeSound()
		animal.Move()
		animal.Feed()
		animal.Sleep()
		animal.Swim()
		animal.Purr()

	}
	if user_animal == "собака" {
		animal := interfaces.Dog{}
		animal.MakeSound()
		animal.Move()
		animal.Feed()
		animal.Sleep()
		animal.Swim()
		animal.Tail()

	}
	if user_animal == "карась" {
		animal := interfaces.Karas{}
		animal.MakeSound()
		animal.Move()
		animal.Feed()
		animal.Sleep()
		animal.Swim()
		animal.Cook()

	}
	if user_animal == "птица" {
		animal := interfaces.Bird{}
		animal.MakeSound()
		animal.Move()
		animal.Feed()
		animal.Sleep()
		animal.Swim()
		animal.Seet()
	}
}

func Hi() {
	fmt.Println("xui")
}
