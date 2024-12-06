package interfaces

import (
	"fmt"
)

type Animal interface {
	MakeSound()
	Move()
	Feed()
	Sleep()
}

type swim interface {
	Swim()
}

type Lion struct{}

func (l Lion) MakeSound() {
	fmt.Println("Лев рычит")
}
func (l Lion) Move() {
	fmt.Println("Лев бежит")
}
func (l Lion) Feed() {
	fmt.Println("Лев ест")
}
func (l Lion) Sleep() {
	fmt.Println("Лев спит")
}
func (l Lion) Hunt() {
	fmt.Println("Лев охотится")
}
func (l Lion) Swim() {
	fmt.Println("Лев плывет")
}

type Cat struct{}

func (c Cat) MakeSound() {
	fmt.Println("кошка мяукает")
}
func (c Cat) Move() {
	fmt.Println("кошка бежит")
}
func (c Cat) Feed() {
	fmt.Println("кошка ест")
}
func (c Cat) Sleep() {
	fmt.Println("кошка спит")
}
func (c Cat) Purr() {
	fmt.Println("кошка мурчит")
}
func (c Cat) Swim() {
	fmt.Println("кошка плавает")
}

type Dog struct{}

func (d Dog) MakeSound() {
	fmt.Println("собака гавкает")
}
func (d Dog) Move() {
	fmt.Println("собака бежит")
}
func (d Dog) Feed() {
	fmt.Println("собака ест")
}
func (d Dog) Sleep() {
	fmt.Println("собака спит")
}
func (d Dog) Tail() {
	fmt.Println("собака виляет хвостом")
}
func (d Dog) Swim() {
	fmt.Println("собака плывёт")
}

type Karas struct{}

func (k Karas) MakeSound() {
	fmt.Println("карась булькает")
}
func (k Karas) Move() {
	fmt.Println("карась плывёт")
}
func (k Karas) Feed() {
	fmt.Println("карась ест")
}
func (k Karas) Sleep() {
	fmt.Println("карась спит")
}
func (k Karas) Cook() {
	fmt.Println("карась жарится")
}
func (k Karas) Swim() {
	fmt.Println("карась плывёт")
}

type Bird struct{}

func (b Bird) MakeSound() {
	fmt.Println("птица поёт")
}
func (b Bird) Move() {
	fmt.Println("птица летит")
}
func (b Bird) Feed() {
	fmt.Println("птица клюёт семечки")
}
func (b Bird) Sleep() {
	fmt.Println("птица спит")
}
func (b Bird) Seet() {
	fmt.Println("птица сидит на электропроводе")
}
func (b Bird) Swim() {
	fmt.Println("птица не плавает")
}
