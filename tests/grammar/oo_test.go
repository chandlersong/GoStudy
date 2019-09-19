package grammar

import (
	"testing"
)

type Person struct {
	hello   Hello
	address string
}

type Hello struct {
	name     string
	password string
}

type Duration int

func (h Hello) notification(t *testing.T) {
	h.name = "new"
	t.Logf("print persion,%v", h)
}

func (h *Hello) notificationWithPointer(t *testing.T) {
	h.name = "pointer"
	t.Logf("print pointer persion,%v", &h)
}

func TestCreateObject(t *testing.T) {
	hello1 := Hello{
		name:     "hello one",
		password: "password",
	}
	hello1.notification(t)
	t.Logf("hello1 %v", hello1) // it won't change the outside problem

	hello1.notificationWithPointer(t)
	t.Logf("hello1 %v", hello1)

	hello2 := Hello{"hello two", "password"}
	t.Logf("hello1 %v", hello2)

	person := Person{
		hello:   Hello{"person", "password"},
		address: "address",
	}
	t.Logf("person %v", person)

	duration := Duration(1)
	t.Logf("duration %v", duration)
}

type notifier interface {
	notify(t *testing.T)
}

type user struct {
	name  string
	email string
}

func (u *user) notify(t *testing.T) {
	t.Logf("sending user email %s<%s> \n",
		u.name, u.email)
}

func TestInterface(t *testing.T) {
	u := &user{"chandler", "chandler605@gmail.com"}
	sendNotification(u, t)
}

func sendNotification(n notifier, t *testing.T) {
	n.notify(t)
}
