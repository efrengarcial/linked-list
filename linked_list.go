package linkedlist

import "fmt"

type Node struct {
	Data interface{}
	Next *Node
}

func NewNode(data interface{}) *Node {
	return &Node{
		Data: data,
		Next: nil,
	}
}

type LinkedList struct {
	Head   *Node
	length int
}

func New() *LinkedList {
	return &LinkedList{
		Head:   nil,
		length: 0,
	}
}

// InsertHead Insertar un nuevo nodo segun la data a la cabeza de la lista
func (l *LinkedList) InsertHead(data interface{}) {
	newNode := NewNode(data)
	newNode.Next = l.Head
	l.Head = newNode
	l.length++
}

// InsertTail Agregar un nuevo nodo al final de la lista
func (l *LinkedList) InsertTail(data interface{}) {
	if l.Head == nil {
		l.InsertHead(data)
		return
	}

	cur := l.Head
	for cur.Next != nil {
		cur = cur.Next
	}
	newNode := NewNode(data)
	cur.Next = newNode
	l.length++
}

// Insert Agregue un nodo  antes del indice  en la lista enlazada.
func (l *LinkedList) Insert(index int, data interface{}) error {
	size := l.length
	if index < 0 || index > size {
		return fmt.Errorf("index out of range")
	}

	if index == 0 {
		l.InsertHead(data)
		return nil
	}

	if index == size {
		l.InsertTail(data)
		return nil
	}

	cur := l.Head
	for i := 0; i < index-1; i++ {
		cur = cur.Next
	}
	node := NewNode(data)
	node.Next = cur.Next
	cur.Next = node
	l.length++

	return nil
}

// DeleteHead Borrar la cabeza de la lista
func (l *LinkedList) DeleteHead() error {
	if l.Head == nil {
		return fmt.Errorf("list is empty")
	}

	l.Head = l.Head.Next
	l.length--
	return nil
}

// Delete Borra el nodo dado el indice  de la lista
func (l *LinkedList) Delete(index int) error {
	size := l.length
	if index < 0 || index > size-1 {
		return fmt.Errorf("index out of range")
	}
	if index == 0 {
		return l.DeleteHead()
	}

	cur := l.Head
	for i := 0; i < index-1; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	l.length--
	return nil

}

// Get Obtiene la data del nodo dado el indice de la lista
func (l *LinkedList) Get(index int) (interface{}, bool) {
	var t interface{}
	if index < 0 || index > l.length-1 {
		return t, false
	}
	cur := l.Head
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	return cur.Data, true
}

// Length Obtiene la longitud de la lista
func (l *LinkedList) Length() int {
	return l.length
}

// Println Imprime la lista
func (l *LinkedList) Println() {
	cur := l.Head
	for cur != nil {
		fmt.Printf("%+v\n", cur.Data)
		cur = cur.Next
	}
}
