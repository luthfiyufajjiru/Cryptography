package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Zn(a, n int) int {
	if n == 0 {
		return 0
	} else if a < n {
		return a
	}
	return a % n
}

func CaesarShift(s string, k int) string {
	var sb strings.Builder

	for _, cr := range s {

		x := 0

		if cr >= 'a' && cr <= 'z' {
			x = 'a'
		} else if cr >= 'A' && cr <= 'Z' {
			x = 'A'
		}

		y := int(cr)

		if x > 0 {
			y = Zn(int(cr)+k-x, 26) + x
		}

		sb.WriteString(string(y))
	}

	return sb.String()
}

func readInput() (plainText string, k int64) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Input your text:")
	plainText, _ = reader.ReadString('\n')
	plainText = strings.Replace(plainText, "\n", "", -1)

	fmt.Println("Input your amount of shift:")

	var kStr string
	fmt.Scan(&kStr)

	isAlpha := regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(kStr)

	if isAlpha {
		panic("please input a number!\n")
	}

	k, _ = strconv.ParseInt(kStr, 10, 64)

	return
}

func writeResult(plainText, cipher string) {
	fmt.Println(fmt.Sprintf("\nResults:\nPlain text: %s\nCipher text: %s", plainText, cipher))
}

func main() {
	plainText, k := readInput()
	cipherText := CaesarShift(plainText, int(k))
	writeResult(plainText, cipherText)
}
