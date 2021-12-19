package container

import "testing"

func TestEmptyListPopInt(t *testing.T) {
	list := []int{}
	_, e, err := Pop(list)
	if err == nil {
		t.Fatalf("Popped from empty list")
	} else if e != 0 {
		t.Fatalf("Wrong return value from empty list")
	}
}

func TestEmptyListPopString(t *testing.T) {
	list := []string{}
	_, e, err := Pop(list)
	if err == nil {
		t.Fatalf("Popped from empty list")
	} else if e != "" {
		t.Fatalf("Wrong return value from empty list")
	}
}

func TestListPushPopInt(t *testing.T) {
	list := []int{}
	list = Push(list, 5)
	list, e, err := Pop(list)
	if err != nil {
		t.Fatalf("Error popping from stack")
	} else if e != 5 {
		t.Fatalf("Wrong value popped from stack")
	} else if len(list) != 0 {
		t.Fatalf("Stack size not decreased after pop")
	}
}

func TestListPushPopString(t *testing.T) {
	list := []string{}
	list = Push(list, "hello")
	list, e, err := Pop(list)
	if err != nil {
		t.Fatalf("Error popping from stack")
	} else if e != "hello" {
		t.Fatalf("Wrong value popped from stack")
	} else if len(list) != 0 {
		t.Fatalf("Stack size not decreased after pop")
	}
}
