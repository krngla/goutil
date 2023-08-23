package warlock

import (
	"log"
	"strconv"
)

func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func rec() {
	if err := recover(); err != nil {
		log.Println("Fear ensued:", err)
		//TODO: implement all the panic cases
	}
}
