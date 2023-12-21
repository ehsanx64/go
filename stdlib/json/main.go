package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Person struct {
	Name         string   `json:"name"`
	Age          int      `json:"age"`
	IsProgrammer bool     `json:"is_programmer"`
	Favorites    []string `json:"favorites"`
	// fields with - will be omitted from marshalization
	PrivateField string `json:"-"`
}

const jsonFile string = "data.json"

func readTheFile() []byte {
	fmt.Println("### Reading the file '" + jsonFile + "'")
	content, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		panic("Could not read the file")
	}

	return content
}

func printTest(title string, data []byte) {
	fmt.Println("### " + title)
	fmt.Println(string(data))
	fmt.Println()
}

func marshalInt() {
	a, _ := json.Marshal(1337)
	printTest("marshalInt", a)
}

func marshalBool() {
	a, _ := json.Marshal(true)
	printTest("marshalBool", a)
}

func marshalSlice() {
	a, _ := json.Marshal([]string{"adam", "jack"})
	printTest("marshalSlice", a)
}

func marshalMap() {
	a, _ := json.Marshal(map[int]string{
		1: "one",
		2: "two",
		3: "three",
	})
	printTest("marshalMap", a)
}

func marshalStruct() []byte {
	p := Person{
		Name:         "john",
		Age:          30,
		IsProgrammer: true,
		Favorites:    []string{"go", "js"},
		PrivateField: "empty",
	}

	fmt.Println("### struct")
	fmt.Println(p)

	jsonData, err := json.Marshal(p)
	if err != nil {
		panic("Could not marshall the data")
	}
	fmt.Println("### JSON data")
	fmt.Println(string(jsonData))

	return jsonData
}

func unmarshalStruct(jsonData []byte) {
	var p Person

	err := json.Unmarshal(jsonData, &p)
	if err != nil {
		panic("Could not unmarshall the data")
	}

	fmt.Println("### Unmarshalled data")
	fmt.Println(p)
	fmt.Println("Name: " + p.Name)
	fmt.Println("Age:", p.Age)
	fmt.Println("IsProgrammer:", p.IsProgrammer)
	fmt.Println("PrivateField:", p.PrivateField)

}

func main() {
	content := readTheFile()
	printTest("content:", content)

	marshalInt()
	marshalBool()
	marshalSlice()
	marshalMap()
	jsonData := marshalStruct()
	unmarshalStruct(jsonData)
}
