package main

import (
	"collectionInterface/controller"
	"fmt"
)

func main() {
	//mySlice := []int{}
	myArrayList := &controller.ArrayListSample{}
	myVector := &controller.VectorSample{}

	fmt.Println(myArrayList.Len())
	fmt.Println(myVector.Len())

	myVector.Push(20)
	myArrayList.Push(10)
	myArrayList.Push(20)
	myArrayList.Push(30)
	myArrayList.Print()
	myVector.Print()
}
