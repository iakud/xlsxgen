package main

import (
	"log"
	"testing"
)

func TestString(t *testing.T) {
	const str = "test1 Test2 test3aTt.TesFt4"
	log.Println(ToPascal(str))
	log.Println(ToSnake(str))
	log.Println(ToCapital(str))

}
