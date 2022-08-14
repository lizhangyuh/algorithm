package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Keyword struct {
	Word      string
	Visit     int
	UpdatedAt *time.Time
}

// 通过序列化和反序列化深拷贝
func (k *Keyword) Clone() *Keyword {
	var newKeyword Keyword
	b, _ := json.Marshal(k)
	json.Unmarshal(b, &newKeyword)
	return &newKeyword
}

type Keywords map[string]*Keyword

func (words Keywords) Clone(update []*Keyword) Keywords {
	newKeywords := Keywords{}

	for k, v := range words {
		newKeywords[k] = v
	}

	for _, word := range update {
		newKeywords[word.Word] = word.Clone()
	}

	return newKeywords
}

func main() {
	updateAt, _ := time.Parse("2006", "2020")
	words := Keywords{
		"testA": &Keyword{
			Word:      "testA",
			Visit:     1,
			UpdatedAt: &updateAt,
		},
		"testB": &Keyword{
			Word:      "testB",
			Visit:     2,
			UpdatedAt: &updateAt,
		},
		"testC": &Keyword{
			Word:      "testC",
			Visit:     3,
			UpdatedAt: &updateAt,
		},
	}

	now := time.Now()
	updatedWords := []*Keyword{
		{
			Word:      "testB",
			Visit:     10,
			UpdatedAt: &now,
		},
	}

	newWords := words.Clone(updatedWords)

	fmt.Printf("testA: %v\n", words["testA"] == newWords["testA"]) // testA: true

	fmt.Printf("testB content: %v\n", words["testB"])                         // testB content: &{testB 2 2020-01-01 00:00:00 +0000 UTC}
	fmt.Printf("testB new content: %v\n", newWords["testB"])                  // testB new content: &{testB 10 2022-08-09 18:53:40.411434 +0800 CST}
	fmt.Printf("testB: %v\n", words["testB"] == newWords["testB"])            // testB: false
	fmt.Printf("updatedWords[0]: %v\n", updatedWords[0] == newWords["testB"]) // updatedWords[0]: false

	newWords["testC"].Word = "new value"                           // 修改 testC 会影响源对象
	fmt.Printf("testC: %v\n", words["testC"] == newWords["testC"]) // testC: true
	fmt.Printf("testC content: %v\n", words["testC"])              // estC content: &{new value 3 2020-01-01 00:00:00 +0000 UTC}

}
