package main

import (
	"container/list"
	"fmt"
)

type linkData struct {
	Next  *linkData
	Value int
}

func main() {
	list.New()
	array := []int{1, 2, 3, 4, 5}
	var head linkData
	var last linkData
	head.Value = array[0]
	head.Next = nil
	last.Next = nil
	last.Value = array[len(array)-1]
	var data *linkData
	for i := 0; i < len(array); i++ {
		data = formLinkData(data, array[i])
	}
	data.rotateRight()
	printList(data)
}

func formLinkData(data *linkData, value int) *linkData {
	if data == nil {
		return &linkData{Next: nil, Value: value}
	}
	data.Next = formLinkData(data.Next, value)
	return data
}

func printList(data *linkData) {
	for data != nil {
		fmt.Println("      ", data.Value)
		data = data.Next
	}
}

func (data *linkData) rotateRight() {
	head := data
	//currentData := data
	var nxData *linkData
	if data != nil {
		for data.Next != nil {
			if data.Next.Next == nil {
				fmt.Println("Next Node ", data.Next.Value)
				fmt.Println("Head ", head.Value)
				data.Next.Next = head
				fmt.Println("Head ", head.Value)
				//data.Next = data.Next.Next

				fmt.Println("Next Next Node ", data.Next.Next.Value)
				nxData.Next.Next = nil
				break
			}
			nxData = data
			data = data.Next
		}
		fmt.Println("Head Out ", data.Value)
	}
}
