package main

import (
	"fmt"
	"strings"
)

type Animal interface {
	sound() string
	move() string
	feed()
	sleep()
}

type swim interface {
	swim() string
}

type Lion struct{}

func (l Lion) MakeSound() {
	fmt.Println("Лев рычит")
}
func (l Lion) move() {
	fmt.Println("Лев бежит")
}
func (l Lion) feed() {
	fmt.Println("Лев ест")
}
func (l Lion) sleep() {
	fmt.Println("Лев спит")
}
func (l Lion) hunt() {
	fmt.Println("Лев охотится")
}
func (l Lion) swim() {
	fmt.Println("Лев плывет")
}

type Cat struct{}

func (c Cat) MakeSound() {
	fmt.Println("кошка мяукает")
}
func (c Cat) move() {
	fmt.Println("кошка бежит")
}
func (c Cat) feed() {
	fmt.Println("кошка ест")
}
func (c Cat) sleep() {
	fmt.Println("кошка спит")
}
func (c Cat) purr() {
	fmt.Println("кошка мурчит")
}
func (c Cat) swim() {
	fmt.Println("кошка плавает")
}

type Dog struct{}

func (d Dog) MakeSound() {
	fmt.Println("собака гавкает")
}
func (d Dog) move() {
	fmt.Println("собака бежит")
}
func (d Dog) feed() {
	fmt.Println("собака ест")
}
func (d Dog) sleep() {
	fmt.Println("собака спит")
}
func (d Dog) tail() {
	fmt.Println("собака виляет хвостом")
}
func (d Dog) swim() {
	fmt.Println("собака плывёт")
}

type Karas struct{}

func (k Karas) MakeSound() {
	fmt.Println("карась булькает")
}
func (k Karas) move() {
	fmt.Println("карась плывёт")
}
func (k Karas) feed() {
	fmt.Println("карась ест")
}
func (k Karas) sleep() {
	fmt.Println("карась спит")
}
func (k Karas) cook() {
	fmt.Println("карась жарится")
}
func (k Karas) swim() {
	fmt.Println("карась плывёт")
}

type Bird struct{}

func (b Bird) MakeSound() {
	fmt.Println("птица поёт")
}
func (b Bird) move() {
	fmt.Println("птица летит")
}
func (b Bird) feed() {
	fmt.Println("птица клюёт семечки")
}
func (b Bird) sleep() {
	fmt.Println("птица спит")
}
func (b Bird) seet() {
	fmt.Println("птица сидит на электропроводе")
}
func (b Bird) swim() {
	fmt.Println("птица не плавает")
}

func main() {
	var user_animal string
	var error int
	for {
		animals_list := []string{"лев", "собака", "кошка", "карась", "птица"}
		fmt.Println("Выберите животное для вывода данных:\n Лев\n Собака\n Кошка\n Карась\n птица")
		fmt.Scanln(&user_animal)
		user_animal = strings.ToLower(user_animal)
		fmt.Println(user_animal)
		for _, animal := range animals_list {
			if animal == user_animal {
				error = 0
				break
			} else {
				fmt.Println("проверьте, правильно ли вы написали название животного")
				error = 1
				break
			}
		}
		if error == 0 {
			break
		} else {
			continue
		}
	}
	if user_animal == "лев" {
		parrot := Lion{}
		parrot.MakeSound()
		parrot.move()
		parrot.feed()
		parrot.sleep()
		parrot.hunt()
		parrot.swim()
	}
	if user_animal == "собака" {
		parrot := Dog{}
		parrot.MakeSound()
		parrot.move()
		parrot.feed()
		parrot.sleep()
		parrot.tail()
		parrot.swim()
	}
	if user_animal == "кот" {
		parrot := Cat{}
		parrot.MakeSound()
		parrot.move()
		parrot.feed()
		parrot.sleep()
		parrot.purr()
		parrot.swim()
	}
	if user_animal == "карась" {
		parrot := Karas{}
		parrot.MakeSound()
		parrot.move()
		parrot.feed()
		parrot.sleep()
		parrot.cook()
		parrot.swim()
	}
	if user_animal == "птица" {
		parrot := Bird{}
		parrot.MakeSound()
		parrot.move()
		parrot.feed()
		parrot.sleep()
		parrot.seet()
		parrot.swim()
	}
}
