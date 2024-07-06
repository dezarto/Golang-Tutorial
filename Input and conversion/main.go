package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Enter a number: ")
	reader := bufio.NewReader(os.Stdin)
	value, errr := reader.ReadString('\n')

	if errr != nil {
		fmt.Println(errr)
		return
	}

	value = strings.TrimSpace(value)
	num, err := strconv.Atoi(value)

	if err != nil {
		fmt.Println("Dönüştürme hatası:", err)
	}

	fmt.Println("After convert: ", num)
}
