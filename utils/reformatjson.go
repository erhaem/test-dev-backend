package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

// below is input structure as in case.json
type Item struct {
	Category string `json:"category"`
	Code     string `json:"code"`
	Name     string `json:"name"`
	Total    int    `json:"total"`
}

type CaseData struct {
	Data []Item `json:"data"`
}

// below is output structure as in expectation.json
type Expectation struct {
	Total int           `json:"total"`
	Data  []CategoryOut `json:"data"`
}

type CategoryOut struct {
	Category string             `json:"category"`
	Total    int                `json:"total"`
	Data     map[string]CodeOut `json:"data"`
}

type CodeOut struct {
	Total int        `json:"total"`
	Data  []NameItem `json:"data"`
}

type NameItem struct {
	Name  string `json:"name"`
	Total int    `json:"total"`
}

// to reformat json
// refs: 
// - https://gobyexample.com/reading-files
// - https://gobyexample.com/writing-files
// - https://gobyexample.com/json
// - chatgpt, see prompt.txt
func ReformatJSON(inputFile, outputFile string) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	var caseData CaseData
	if err := json.Unmarshal(bytes, &caseData); err != nil {
		log.Fatal(err)
	}

	exp := Expectation{
		Total: 0,
		Data:  []CategoryOut{},
	}

	// map for grouping categories
	catMap := make(map[string]*CategoryOut)

	for _, item := range caseData.Data {
		if _, ok := catMap[item.Category]; !ok {
			catMap[item.Category] = &CategoryOut{
				Category: item.Category,
				Total:    0,
				Data:     make(map[string]CodeOut),
			}
		}
		cat := catMap[item.Category]

		// update code group inside category
		code := cat.Data[item.Code]
		code.Total += item.Total
		code.Data = append(code.Data, NameItem{
			Name:  item.Name,
			Total: item.Total,
		})
		cat.Data[item.Code] = code

		// update totals for category and global
		cat.Total += item.Total
		exp.Total += item.Total
	}

	// convert map to slice
	for _, cat := range catMap {
		exp.Data = append(exp.Data, *cat)
	}

	// encode json
	formatted, err := json.MarshalIndent(exp, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(outputFile, formatted, 0644); err != nil {
		log.Fatal(err)
	}

	fmt.Println("JSON saved to:", outputFile)
}
