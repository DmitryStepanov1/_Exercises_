package main

import (
	"fmt"
	"log"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	err := os.WriteFile(p.Title, p.Body, 0666)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func load(title string) (*Page, error) {
	data, err := os.ReadFile(title)
	if err != nil {
		log.Fatal(err)
	}

	var b = new(Page)
	b.Title = title
	b.Body = append(b.Body, data...)

	fmt.Println(b)

	return b, err
}

func main() {

	var b = new(Page)
	//var b *Page
	b.Title = "java"
	b.Body = append(b.Body, 56)
	fmt.Println(b)

	b.save()
	load("golang")

}
