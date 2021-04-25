package main

import (
	"fmt"
)

// 本质上是一个指针
type AnimalIF interface {
	Sleep()
	GetColor() string
	GetType() string
}

// 具体的类
type Cat struct {
	color string
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is Sleep..")
}
func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}

// 具体的类
type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is sleep..")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

func showAnimal(animalIF AnimalIF) {
	animalIF.Sleep() //多态的现象
	fmt.Println("color = ", animalIF.GetColor())
	fmt.Println("kind = ", animalIF.GetType())
}

func main() {

	var animal AnimalIF // 接口的数据类型。父类指针
	animal = &Cat{"red"}
	animal.Sleep() // 调用的是Cat的Sleep（）方法，多态的现象
	animal = &Dog{"yello"}
	animal.Sleep() // 调用Dog的Sleep方法，多态的现象

	cat := Cat{"wilte"}
	dog := Dog{"green"}
	showAnimal(&cat)
	showAnimal(&dog)
}
