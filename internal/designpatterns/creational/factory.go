package creational

type Animal interface {
	Speak() string
}

type Wolf struct{}

func (w Wolf) Speak() string {
	return "wowowo"
}

type Cow struct{}

func (c Cow) Speak() string {
	return "moumoumou"
}

// Factory patterns: create objects using a factory function.
func NewAnimal(kind string) Animal {
	switch kind {
	case "wolf":
		return Wolf{}
	case "cow":
		return Cow{}
	}
	return nil
}
