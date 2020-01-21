package _1324_print_words_vertically

import (
	"log"
	"testing"
)

func TestTask2(t *testing.T) {
	res := printVertically("HOW ARE YOU")
	for i := range res {
		log.Println(len(res[i]))
	}
	log.Println(res)
	log.Println(printVertically("TO BE OR NOT TO BE"))
	log.Println(printVertically("CONTEST IS COMING"))
}
