package A

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)

	line, _, _ := reader.ReadLine()

	strArr := strings.Split(string(line), " ")

	intA, _ := strconv.Atoi(strArr[0])
	intB, _ := strconv.Atoi(strArr[1])

	fmt.Println(intA + intB)
}
