// Copyright 2017 Gruppe 12 IS-105. All rights reserved.
// Kode fra Gruppe 2 IS-105 kan være delvis gjennbrukt her pga bytte av medlemmer.
package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	tall := os.Args[1]
	tallFloat, _ := strconv.ParseFloat(tall, 64)
	fmt.Println("Logaritmen av ", tall, " blir ", math.Log2(tallFloat), " med base 2")
}
