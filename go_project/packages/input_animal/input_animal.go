package input_animal

import (
	"fmt"
	"strings"
)

func Animal_input() string {
	animals_list := []string{"лев", "собака", "кошка", "карась", "птица"}
	var user_animal string
	var error = 1
	for {
		fmt.Println("Выберите животное для вывода данных:\n Лев\n Собака\n Кошка\n Карась\n Птица")
		fmt.Scanln(&user_animal)
		user_animal = strings.ToLower(user_animal)
		fmt.Println(user_animal)
		for _, animal := range animals_list {
			if animal == user_animal {
				error = 0
				break
			}

		}
		if error == 0 {
			return user_animal
		} else {
			fmt.Println("no")
			continue
		}

	}

}
