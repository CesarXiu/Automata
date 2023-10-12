// You can edit this code!
// Click here and start typing.
// buffio, os
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewReader(os.Stdin)
	fmt.Println("Valor en X:")
	xValue, _ := scanner.ReadString('\n')
	xValue = strings.TrimSpace(xValue)
	fmt.Println("Valor en Y:")
	yValue, _ := scanner.ReadString('\n')
	yValue = strings.TrimSpace(yValue)
	x, _ := strconv.Atoi(xValue)
	y, err := strconv.Atoi(yValue)

	if err != nil {
		fmt.Println("Error al convertir la cadena a entero:", err)
	}

	fmt.Println("\n El resultado de la suma es: ", suma(x, y))
}

func suma(x, y int) int {
	return x + y
}
