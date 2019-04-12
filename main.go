package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Histogram struct {
	Buckets []Bucket
}

type Bucket struct {
	Value   int
	Counter uint
}

func NewHistogram() *Histogram {
	return &Histogram{}
}

func (h *Histogram) Add(value int) error {
	for i, b := range h.Buckets {
		if b.Value == value {
			h.Buckets[i].Counter++
			return nil
		}
	}
	histogram := Bucket{
		Value:   value,
		Counter: 1,
	}
	h.Buckets = append(h.Buckets, histogram)
	return nil
}

func (h *Histogram) TotalBytes() (uint, error) {
	counter := uint(0)
	for _, b := range h.Buckets {
		counter = counter + b.Counter
	}
	return counter, nil
}

func main() {
	stat := NewHistogram()

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalln("Read data from stdin failed")
	}
	for _, x := range data {
		stat.Add(int(x))
	}
	for _, result := range stat.Buckets {
		fmt.Println("Bucket: " + fmt.Sprint(result.Value) + " | " + fmt.Sprint(result.Counter))
	}
	total, _ := stat.TotalBytes()
	fmt.Println("Total bytes hashed: " + fmt.Sprint(total))
}
