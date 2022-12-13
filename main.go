package main

import (
	"fmt"
	"linkedlist/internal"
)

type Product struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func main() {
	ll := internal.NewLinkedList()
	ll.InsertHead(Product{ID: 1, Name: "p1"})
	ll.InsertHead(Product{ID: 2, Name: "p2"})
	ll.InsertTail(Product{ID: 3, Name: "p3"})
	ll.Insert(0, Product{ID: 4, Name: "p4"})
	ll.Println()
	fmt.Println("Length=", ll.Length())
	fmt.Println("****************")
	ll.DeleteHead()
	ll.Println()
	fmt.Println("Length=", ll.Length())
	fmt.Println("****************")
	ll.Delete(2)
	ll.Println()
	fmt.Println("Length=", ll.Length())

}
