package utils

import (
	"log"
	"os"
	"path/filepath"
)

type document struct {
	Title string `xml:"title"`
	URL   string `xml:"url"`
	Text  string `xml:"abstract"`
	ID    int    `xml:"id"`
}

// LoadDocuments reads all files in a directory as plain text
func LoadDocuments(dirPath string) ([]document, error) {
	var docs []document

	// Read all files in the directory
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	docID := 0
	for _, file := range files {
		filePath := filepath.Join(dirPath, file.Name())

		// Skip directories
		if file.IsDir() {
			continue
		}

		// Read file content
		content, err := os.ReadFile(filePath)
		if err != nil {
			log.Println("Error reading file:", filePath, err)
			continue
		}

		// Treat all files as plain text
		docs = append(docs, document{
			Title: file.Name(),
			Text:  string(content),
			ID:    docID,
		})
		docID++
	}

	return docs, nil
}
