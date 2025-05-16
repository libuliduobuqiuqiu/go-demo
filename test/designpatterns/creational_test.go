package designpatterns

import (
	"godemo/internal/designpatterns/creational"
	"testing"
)

func TestAbstractFactory(t *testing.T) {
	w := creational.WinFactory{}
	m := creational.MacFactory{}

	creational.Render(w)
	creational.Render(m)
}

func TestBuilder(t *testing.T) {
	s := creational.NewServer(creational.WithHost("127.0.0.1"), creational.WithPort(8090))
	t.Log(s.Host, s.Port)
}
