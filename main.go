package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
	"flag"
)

func main() {
	//seedin random number to current time ensures it will be reasonably random
	rand.Seed( time.Now().UTC().UnixNano())

	//create and parse command line flags
	uptoInt := flag.Int("b", 12, "The Highest B value * 2")
	flag.Parse()

	//variable i will hold the user input
	var i int

	for {
		base := rand.Intn(*uptoInt) //base number to build trinomial
		b := base * 2 // b value
		c := int(math.Pow(float64(base), 2)) //c value
		fmt.Printf("B=%d c=?", b) //ask question
		fmt.Scan(&i) //get user input

		if i != c { //check answer, print if wrong
			fmt.Printf("Wrong. B=%d C=%d",b, c)
		}
	}
}
