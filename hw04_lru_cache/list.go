package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

// ListItem is an element of a doubly linked list.
type ListItem struct {
	Key   Key         // key of the item
	Value interface{} // value of the item
	Next  *ListItem   // next item in the list
	Prev  *ListItem   // previous item in the list
}

type list struct {
	size int       // number of elements in the list
	head *ListItem // first element of the list
	tail *ListItem // last element of the list
}

// NewList creates a new list.
func NewList() List {
	return new(list)
}

// Len returns the number of elements in the list.
func (l *list) Len() int {
	return l.size
}

// Front returns the first element of the list or nil if the list is empty.
func (l *list) Front() *ListItem {
	return l.head
}

// Back returns the last element of the list or nil if the list is empty.
func (l *list) Back() *ListItem {
	return l.tail
}

// PushFront adds a new element with the specified value at the beginning of the list.
func (l *list) PushFront(v interface{}) *ListItem {
	if l.size == 0 {
		l.head = &ListItem{Value: v}
		l.tail = l.head
	} else {
		l.head = &ListItem{Value: v, Next: l.head}
		l.head.Next.Prev = l.head
	}

	l.size++
	return l.head
}

// PushBack adds a new element with the specified value at the end of the list.
func (l *list) PushBack(v interface{}) *ListItem {
	if l.size == 0 {
		l.tail = &ListItem{Value: v}
		l.head = l.tail
	} else {
		l.tail = &ListItem{Value: v, Prev: l.tail}
		l.tail.Prev.Next = l.tail
	}

	l.size++
	return l.tail
}

// Remove removes the specified element from the list.
func (l *list) Remove(i *ListItem) {
	if i == nil {
		return
	}

	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.head = i.Next
	}

	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.tail = i.Prev
	}

	l.size--
}

// MoveToFront moves the specified element to the beginning of the list.
func (l *list) MoveToFront(i *ListItem) {
	if i == nil {
		return
	}

	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		return
	}

	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.tail = i.Prev
	}

	i.Prev = nil
	i.Next = l.head
	l.head.Prev = i
	l.head = i
}
