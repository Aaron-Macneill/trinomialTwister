package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
	"flag"
)

func main() {
	rand.Seed( time.Now().UTC().UnixNano())

	uptoInt := flag.Int("b", 12, "The Highest B value * 2")
	flag.Parse()

	var i int

	for {
		base := rand.Intn(*uptoInt)
		b := base * 2
		c := int(math.Pow(float64(base), 2))
		fmt.Printf("B=%d c=?", b)
		fmt.Scan(&i)

		if i != c {
			fmt.Printf("Wrong. B=%d C=%d",b, c)
		}
	}
}
