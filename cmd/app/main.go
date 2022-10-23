package main

import (
	"fmt"
	"github.com/andrew55516/Shodan/internal"
	"log"
	"os"
)

func main() {
	log.Println(os.Args)
	if len(os.Args) != 2 {
		log.Fatalln("Usage: shodan searchterm")
	}
	apiKey := os.Getenv("SHODAN_API_KEY")
	s := internal.NewClient(apiKey)
	info, err := s.APIInfo()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf(
		"Query Credits: %d\nScan Credits: %d\n\n",
		info.QueryCredits, info.ScanCredits)

	hostSearch, err := s.HostSearch(os.Args[1])

	if err != nil {
		log.Fatalln(err)
	}
	for _, host := range hostSearch.Matches {
		fmt.Printf("%18s%8d\n", host.IPString, host.Port)
	}
}
