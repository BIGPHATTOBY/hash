package main

import (
	//"bytes"
	"fmt"
	//stats "google.golang.org/grpc/benchmark/stats"
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
	fmt.Println(stat.Buckets)
	fmt.Println(stat.TotalBytes())
}
