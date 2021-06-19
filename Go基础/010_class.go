package main

import "fmt"

type Human struct {
	Name string
	sex  string
}

func (this Human) Eat() {
	fmt.Println("Human.Eat()..")
}
func (this Human) Walk() {
	fmt.Println("Human.Walk..")
}

type SuperMan struct {
	Human // 继承Human
	lever int
}

// 重定义父类的方法Eat()
func (this *SuperMan) Eat() {
	fmt.Println("SurperMan.Eat()..")
}

//子类的新方法
func (this *SuperMan) Fly() {
	fmt.Println("SurperMan.Fly()..")
}

// 定义子类的新方法
func (this *SuperMan) Print() {
	fmt.Println("Name = ", this.Name)
	fmt.Println("Sex = ", this.sex)
	fmt.Println("Lever = ", this.lever)
}

func main() {
	h := Human{
		"zhangsan",
		"female",
	}
	h.Eat()
	h.Walk()

	// 定义一个子类对象
	s := SuperMan{Human{"lisi", "female"}, 99}
	fmt.Println("name = ", s.Name)
	fmt.Println("sex = ", s.sex)
	fmt.Println("lever = ", s.lever)

	s.Eat()
	s.Walk()
	s.Fly()
	s.Print()
}
