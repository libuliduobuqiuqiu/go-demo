package structdemo

import (
	"fmt"
)

// 不同情况下嵌入的结构体类型
// 值嵌入T:
// 1. 零值可用 2. 更像是组成部分，而不是引入外部对象 3. 希望拷贝外层对象是独立副本
// 指针嵌入*T:
// 1. 被嵌入对象很大，避免值拷贝成本 2. 需要共享底层状态（多个对象共用一个嵌入对象） 3. 被嵌入类型的方法主要是指针接受者
//
// 为什么“值嵌入 + 指针接收者方法”时，外层结构体的指针类型 *S 能实现接口，而值类型 S 不能？
// 1. 方法调用可以使用语法糖 2. 接口实现必须基于严格的方法集

type PrintInt interface {
	Print()
}

type Info struct {
	Name string `json:"name"`
	Addr string `json:"addr"`
}

func (i *Info) Print() {
	fmt.Printf("My name is %s, addr is %s\n", i.Name, i.Addr)
}

type Student struct {
	*Info
}

func (s *Student) Log() {
	fmt.Println("I'm student.")
}

func NewStudent(Name, Addr string) Student {
	s := Student{Info: &Info{}}
	s.Name = Name
	s.Addr = Addr
	return s
}

type Teacher struct {
	Info
}

func (t *Teacher) Log() {
	fmt.Println("I'm teacher.")
}

func NewTeacher(Name, Addr string) Teacher {
	t := Teacher{}
	t.Name = Name
	t.Addr = Addr
	return t
}

func RunDemo() {
	var s PrintInt
	s = NewStudent("zhangsan", "Beijing")

	// var t PrintInt
	// error: 作为值被拷贝，拷贝后的值和原来的值是独立的，如果调用指针接受者的方法，这时会产生歧义，到底需要修改哪一份？ （即不可寻址）
	// t = NewTeacher("wangwu", "Guangzhou")

	s.Print()
	// t.Print()
}
