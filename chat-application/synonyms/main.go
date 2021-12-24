package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"go-practice.com/chat-application/thesaurus"
)

func main() {
	apiKey := os.Getenv("BHT_APIKEY")
	thesaurus := &thesaurus.BigHuge{APIKey: apiKey}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := s.Text()
		syns, err := thesaurus.Synonyms(word)
		if err != nil {
			log.Fatalf("Failed when looking for synonyms for word '%s' %s", word, err)
		}
		if len(syns) == 0 {
			log.Fatalf("Couldn't find any synonyms for word '%s'", word)
		}
		for _, syn := range syns {
			fmt.Println(syn)
		}
	}
}
