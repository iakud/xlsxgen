package stringcase

import (
	"log"
	"testing"
)

func TestStringCase(t *testing.T) {
	const str = "test1 Test2 teSt3aTt.TesFt4"
	log.Println(ToCamel(str))
	log.Println(ToPascal(str))
	log.Println(ToSnake(str))
	log.Println(ToKebab(str))
}
