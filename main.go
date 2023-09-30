package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gocolly/colly"
)

func main() {
	fName := "data.csv"
	file, err := os.Create(fName)

	if err != nil {
		log.Fatalf("Could not create file, err : %q", err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	cllct := colly.NewCollector(colly.AllowedDomains("internshala.com"))

	cllct.OnHTML(".internship_meta", func(elm *colly.HTMLElement) {
		writer.Write([]string{
			elm.ChildText("a"),
			elm.ChildText("span"),
		})
	})

	for i := 0; i < 312; i++ {
		fmt.Printf("Scraping Page : %d\n", i)

		cllct.Visit("https://internshala.com/internships/page-" + strconv.Itoa(i))
	}

	log.Printf("Scraping Complete !")
	log.Println(cllct)
}
