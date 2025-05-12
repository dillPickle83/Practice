package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
)

type contactInfo struct{
	Name string
	email string
}

type purchaseInfo struct{
	Name string
	Price string
	Amount string
}

func main(){
	var contacts []contactInfo = loadJSON[contactInfo]("./contactInfo.json")
	fmt.Printf("\n%+v", contacts)

	var purchases []purchaseInfo = loadJSON[purchaseInfo]("./purchaseInfo.json")
	fmt.Printf("\n%+v", purchases)
}

func loadJSON[T contactInfo | purchaseInfo](filepath string) []T{
	var data, _ = ioutil.ReadFile(filepath)

	var loaded = []T{}
	json.Unmarshal(data, &loaded)

	return loaded
}
