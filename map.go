package main

import "fmt"

func main() {
	mymap := make(map[string]int)
	var numpair int
	fmt.Print("enter the no of key pairs")
	fmt.Scan(&numpair)
	for i := 0; i < numpair; i++ {
		var key string
		var value int
		fmt.Printf("enter the key %d:", i+1)
		fmt.Scan(&key)
		fmt.Printf("enter the value %d:", i+1)
		fmt.Scan(&value)
		mymap[key] = value
	}
	var key string
	fmt.Println("enter the key to retrive value:")
	fmt.Scan(&key)

	value, exists := mymap[key]
	if exists {
		fmt.Printf("%s value of %d", key, value)

	} else {
		fmt.Print("%s not found", key)
	}
	var updatekey string
	var updatevalue int
	fmt.Println("enter the key to update:")
	fmt.Scan(&updatekey)
	fmt.Println("enter the value:")
	fmt.Scan(&updatevalue)
	mymap[updatekey] = updatevalue
	var deletekey string
	fmt.Println("enter the key to delete:")
	fmt.Scan(&deletekey)
	delete(mymap, deletekey)
	fmt.Print("Adfter deleting", deletekey, ":", mymap)
	for key, value := range mymap {
		fmt.Println("key:", key, "value:", value)

	}
	fmt.Println("length is:", len(mymap))

}
