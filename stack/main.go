package stack

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

var ErrStackPopEmpty = errors.New("Stack empty, can't pop")

type Stack struct {
	Head  int
	Stack []string
}

func NewStack(stack []string) Stack {
	ret := Stack{
		Head:  len(stack),
		Stack: stack,
	}
	return ret
}

func (s *Stack) Push(str string) {
	if s.Head < len(s.Stack) {
		s.Stack[s.Head] = str
	} else {
		s.Stack = append(s.Stack, str)
	}
	s.Head += 1
}
func (s *Stack) Pop() (string, error) {
	s.Head -= 1
	if s.Head < 0 {
		return "", ErrStackPopEmpty
	}
	ret := s.Stack[s.Head]
	return ret, nil
}

func (s Stack) Ser() []byte {
	if s.Head != 0 {
		s.Stack = s.Stack[:s.Head]
	}
	data, err := json.Marshal(s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to serialize struct\n")
	}
	return data
}

func (s *Stack) Deser(data []byte) {
	err := json.Unmarshal(data, s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to deserialize data: %s\n", data)
	}
}

func (s Stack) Print() string {
	str := ""
	for i := 0; i < s.Head; i++ {
		str += s.Stack[i]
	}
	return str
}
