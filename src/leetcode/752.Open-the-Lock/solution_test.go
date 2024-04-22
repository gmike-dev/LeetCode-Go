package _52_Open_the_Lock

import "testing"

func TestOpenLock1(t *testing.T) {
	actual := openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202")
	expect := 6
	if actual != expect {
		t.Errorf("expect %d, but got %d", expect, actual)
	}
}
func TestOpenLock2(t *testing.T) {
	actual := openLock([]string{"8888"}, "0009")
	expect := 1
	if actual != expect {
		t.Errorf("expect %d, but got %d", expect, actual)
	}
}
func TestOpenLock3(t *testing.T) {
	actual := openLock([]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888")
	expect := -1
	if actual != expect {
		t.Errorf("expect %d, but got %d", expect, actual)
	}
}

func TestOpenLock4(t *testing.T) {
	actual := openLock([]string{"0234"}, "1234")
	expect := 10
	if actual != expect {
		t.Errorf("expect %d, but got %d", expect, actual)
	}
}
func TestOpenLock5(t *testing.T) {
	actual := openLock([]string{"0000"}, "8888")
	expect := -1
	if actual != expect {
		t.Errorf("expect %d, but got %d", expect, actual)
	}
}
