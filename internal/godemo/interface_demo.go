package godemo

import "fmt"

type handler struct {
	name string
}

func (h *handler) GetName() string {
	return h.name
}

func (h *handler) SetName(name string) {
	h.name = name
}

func UseHandler() {

	h := handler{}
	h.SetName("zhangsan")
	fmt.Println("name:", h.GetName())

	h2 := &handler{}
	h2.SetName("lisi")
	fmt.Println("name:", h2.GetName())

}
