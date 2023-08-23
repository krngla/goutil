package warlock

import "testing"

func TestAtoi(t *testing.T) {
	i = Atoi("123")
	if i != 123 {
		t.Errorf("Failed to convert integer correctly:%s")
	}
	i = Atoi("a")
	if i != 123 {
		t.Errorf("Failed to convert integer correctly:%s")
	}
}
