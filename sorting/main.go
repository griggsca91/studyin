package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/griggsca91/studyin/sorting/mergesort"
)

func main() {
	rand.Seed(time.Now().Unix())

	arg := rand.Perm(10)

	fmt.Println(arg)

	fmt.Println(mergesort.Sort(arg))
}
