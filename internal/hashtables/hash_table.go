package hashtables

import (
	"fmt"

	"simplecode/internal/linkedlists"
)

type HashMap struct {
	Data   []*linkedlists.LinkedList
	Length int
}

func Init(h *HashMap, length int) {
	h.Length = length
	h.Data = make([]*linkedlists.LinkedList, length)
	for i := range h.Data {
		h.Data[i] = &linkedlists.LinkedList{}
	}
}

func (h *HashMap) hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % h.Length
}

func (h *HashMap) Insert(key string) {
	index := h.hash(key)
	h.Data[index].AddNode(key)
}

func (h *HashMap) Get(key string) *linkedlists.Node {
	index := h.hash(key)
	node := h.Data[index].SearchList(key)
	if node == nil {
		fmt.Printf("Key %q not found\n", key)
	}
	return node
}

func (h *HashMap) Delete(key string) {
	index := h.hash(key)
	h.Data[index].DeleteNode(key)

}

func (h HashMap) PrintMap() {
	for i, v := range h.Data {
		fmt.Printf("Map Index #%d: ", i)
		v.PrintList()
	}
}
