package stack

import (
	"testing"
)

func TestSer(t *testing.T) {
	colum := Stack{
		Head:  4,
		Stack: []string{"A", "B", "C", "D"},
	}
	d := colum.Ser()
	expected := "{\"Head\":4,\"Stack\":[\"A\",\"B\",\"C\",\"D\"]}"
	if string(d) != expected {
		t.Errorf("failed to serialize!\n\tgot:\t\t%s\n\texpected:\t%s\n", string(d), expected)
	}
	var s Stack
	s.Deser(d)
	d2 := s.Ser()
	if string(d2) != string(d) {
		t.Errorf("failed to deserialize!\n\tgot:\t\t%s\n\texpected:\t%s\n", string(d2), string(d))
	}
}

func TestPush(t *testing.T) {
	type testcase struct {
		inputs   []string
		expected string
	}
	tc := testcase{
		inputs:   []string{"A", "B", "C", "D"},
		expected: "{\"Head\":4,\"Stack\":[\"A\",\"B\",\"C\",\"D\"]}",
	}
	s := Stack{}
	for _, inp := range tc.inputs {
		s.Push(inp)
	}

	ser := string(s.Ser())

	if ser != string(tc.expected) {
		t.Errorf("Failed expected:\n%s \n got:\n%s\n", ser, tc.expected)
	}
}

func TestPop(t *testing.T) {
	type testcase struct {
		stack        Stack
		n            int
		expectedHead int
		expected     []byte
	}
	tc := testcase{
		stack: Stack{
			Head:  4,
			Stack: []string{"a", "b", "c", "d"},
		},
		n:            2,
		expectedHead: 2,
		expected:     []byte("{\"Head\":2,\"Stack\":[\"a\",\"b\"]}"),
	}

	for i := 0; i < tc.n; i++ {
		tc.stack.Pop()
	}
	if tc.stack.Head != tc.expectedHead {
		t.Errorf("Head value not correct\n\texpected: %d\n\tgot: %d\n",
			tc.expectedHead,
			tc.stack.Head)
	}
	ser := string(tc.stack.Ser())

	if ser != string(tc.expected) {
		t.Errorf("Failed expected:\n%s\ngot:\n%s\n", tc.expected, ser)
	}
	tc.stack.Pop()
	tc.stack.Pop()
	_, err := tc.stack.Pop()
	if err != ErrStackPopEmpty {
		t.Errorf("Failed to prevent popping of empty stack,(%s)", tc.stack.Ser())
	}
}
