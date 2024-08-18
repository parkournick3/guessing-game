package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Guessing Game")
	fmt.Println("A random number will be drawn. Try to get it right. The number is an integer between 0 and 100.")

	scanner := bufio.NewScanner(os.Stdin)

	attempts := [10]int64{}

	n := rand.Int64N(101)

	for i := range attempts {
		fmt.Println("what is the number?")
		scanner.Scan()

		attempt := scanner.Text()
		attempt = strings.TrimSpace(attempt)

		attemptInt, err := strconv.ParseInt(attempt, 10, 64)
		if err != nil {
			fmt.Println("your attempt must be an integer")
			return
		}

		switch {
		case attemptInt < n:
			fmt.Println("wrong attempt, number is greater than ", attemptInt)
		case attemptInt > n:
			fmt.Println("wrong attempt, number is less than ", attemptInt)
		case attemptInt == n:
			fmt.Printf(
				"Congratulations, you got the number right, which was: %d\n"+
					"you got it right in %d tries\n"+
					"These were your attempts: %v\n",
				n,
				i,
				attempts[:i],
			)
			return
		}

		attempts[i] = attemptInt
	}

	fmt.Printf(
		"Unfortunately you didn't get the number right, which was: %d\n"+
			"You had 10 attempts\n"+
			"These were your attempts: %v\n",
		n,
		attempts,
	)
}
