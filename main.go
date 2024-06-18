package main

import (
	"fmt"

	"github.com/recor-glitch/Go-practice/Link-List/doublyLinkList"
	"github.com/recor-glitch/Go-practice/Link-List/singleLinkList"
	"github.com/recor-glitch/Go-practice/SafeMap/multiSafeMap"
	"github.com/recor-glitch/Go-practice/SafeMap/singleSafeMap"
)

// DUMMY DATA

var data []any = []any{
	map[string]string{"A": "value1", "B": "value2"},
	"value1",
	1,
	true,
	[]int{1, 2, 3, 4, 5},
	"value2",
}

func singleSafeMapFunc() {

	// CREATE A NEW INSTANCE OF SAFE MAP
	safeMap := singleSafeMap.NewSafeMap()

	// SETTING VALUES TO THE SAFE MAP
	safeMap.Set("key1", "value1")
	safeMap.Set("key2", []int{1, 2, 3, 4, 5})

	// GETTING THE VALUE AND PRINTING THE VALUE
	result := safeMap.Get("key1")
	result2 := safeMap.Get("key2")
	fmt.Printf("Result of Key 1: %v", result)
	fmt.Printf("Result of Key 2: %+v", result2)

	// DELETE THE VALUE FROM THE MAP
	safeMap.Delete("key1")
}

func multiSafeMapFunc() {

	// Initialize the multi safe map
	multiSafeMap := multiSafeMap.NewMultiSafeMap(5)
	fmt.Printf("My safe map: %v\n", multiSafeMap)

	// SET VALUE TO THE MULTI SAFE MAP IN PARTICULAR INDEX
	_, err := multiSafeMap.Set("Key1", "value1")
	if err != nil {
		fmt.Printf("Error while setting value to %v: %v", multiSafeMap, err)
		return
	}
	multiSafeMap.Display()

	// GET VALUE FROM THE MULTI SAFE MAP
	value, e := multiSafeMap.Get("Key1")
	if e != nil {
		fmt.Printf("Error: %v\n", e.Error())
		return
	}
	fmt.Printf("The value is: %v\n", value)

	// DELETE VALUE FROM THE MULTI SAFE MAP
	_, dEr := multiSafeMap.Delete("Key1")
	if dEr != nil {
		fmt.Printf("Error: %v\n", dEr.Error())
		return
	}
	fmt.Printf("Successfully deleted the key from the map")
}

func singlyLinkListFunc() {

	ll := singleLinkList.LinkList{}

	// INSERT DATA INTO LINK LIST
	for i, d := range data {
		_, err := ll.Insert(fmt.Sprintf("key%v", i), d)
		if err != nil {
			fmt.Println(err)
		}
	}

	// DISPLAY DATA FROM LINK LIST
	ll.Display()
}

func doublyLinkListFunc() {

	dl := doublyLinkList.LinkList{}

	// INSERT DATA INTO DOUBLY LINK LIST
	for _, d := range data {
		dl.Insert(d)
	}

	// INSERT DATA INTO PARTICULAR INDEX
	dl.InsertElementInIndex(5, "I am here at index 5")
	dl.InsertElementInIndex(0, "I am here at index 0")
	dl.InsertElementInIndex(2, "I am here at index 2")

	// DISPLAY DATA FROM DOUBLY LINK LIST
	dl.Display()
}

func main() {

	// LINK LIST

	// SINGLY LINK LIST
	// singlyLinkListFunc()

	// DOUBLY LINK LIST
	doublyLinkListFunc()

	// SAFE MAP

	// SINGLE SAFE MAP
	// singleSafeMapFunc()

	// MULTI SAFE MAP
	// multiSafeMapFunc()
}
