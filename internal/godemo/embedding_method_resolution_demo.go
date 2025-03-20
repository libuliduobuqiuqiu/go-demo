package godemo

import "fmt"

type Model interface {
	Use() error
}

type Car struct {
	Name string
}

func (c *Car) GetObjType() string {
	return "car"
}

func (c *Car) Use() error {
	fmt.Println(c.GetObjType())
	return nil

}

type Byd struct {
	*Car
	Address string
}

func (b *Byd) GetObjType() string {
	return "byd"
}

func StartByd() {
	b := Byd{
		Car:     &Car{Name: "byd"},
		Address: "Hangzhou",
	}

	// 需要注意组合不是继承，实际解析规则根据方法绑定的接收者为准，Use方法调用绑定的是*Car对象的方法
	// 所以调用链顺序是：b.Use() -> b.Car.Use() -> b.Car.GetObjType()
	b.Use()
}
