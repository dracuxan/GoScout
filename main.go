package main

import (
	"flag"
	"log"
	"strings"
	"time"

	"github.com/dracuxan/GoScout/utils"
)

func main() {
	var dumpPath, query string
	flag.StringVar(&dumpPath, "p", "./", "file to search")
	flag.StringVar(&query, "q", "", "query to search for")
	flag.Parse()

	log.Println("Full text seach in progress...")

	start := time.Now()
	docs, err := utils.LoadDocuments(dumpPath)
	if err != nil {
		log.Fatalf("Files not found error: %v", err)
	}
	log.Printf("Loaded %d documents in %v", len(docs), time.Since(start))

	start = time.Now()
	idx := make(utils.Index)
	idx.Add(docs)
	log.Printf("Loaded %d documents in %v", len(docs), time.Since(start))

	start = time.Now()
	matchIds := idx.Search(query)
	log.Printf("Search found: %d documents in %v", len(matchIds), time.Since(start))

	for _, id := range matchIds {
		doc := docs[id]
		for line := range strings.SplitSeq(doc.Text, "\n") {
			if strings.Contains(strings.ToLower(line), strings.ToLower(query)) {
				log.Printf("File %d:\t%v\t%s\n", doc.ID+1, doc.Title, strings.TrimSpace(line)) // Print only matching lines
			}
		}
	}
}
