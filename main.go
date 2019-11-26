package main

import (
	"encoding/csv"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"os"
)

func main() {
	f, err := os.Open("name_file.htm")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	doc, err := goquery.NewDocumentFromReader(f)

	if err != nil {
		log.Fatal(err)
	}

	sel := doc.Find("table")

	tr := sel.Find("tr")

	file, err := os.Create("result_file.csv")
	checkError("Cannot create file", err)

	defer file.Close()

	writer := csv.NewWriter(file)

	defer writer.Flush()

	for i := range tr.Nodes {
		cols := []string{}
		td := tr.Eq(i).Find("td")
		for n := range td.Nodes {
			cols = append(cols, td.Eq(n).Text())
		}
		fmt.Printf("%q\n", cols)
		if err := writer.Write(cols); err != nil {
			checkError("Cannot write to file", err)
		}
	}
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
