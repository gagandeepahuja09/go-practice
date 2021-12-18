package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const otherWord = "*"

var transforms = []string{
	otherWord,
	otherWord + "app",
	otherWord + "site",
	otherWord + "time",
	"get" + otherWord,
	"go" + otherWord,
	"lets" + otherWord,
	otherWord + "hq",
}

func main() {
	// computers can't actually generate random no.s but changing the
	// seed gives the illusion that it can. As the seed would be diff
	// everytime the program is run.
	rand.Seed(time.Now().UnixNano())
	// since it takes io.Reader as input, we can have different
	// sources of input. For unit tests we can specify our own input source
	// without worrying about simulating standard input.
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		t := transforms[rand.Intn(len(transforms))]
		fmt.Println(strings.Replace(t, otherWord, s.Text(), -1))
	}
}
