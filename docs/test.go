package main

import (
	"fmt"
	"log"
	"reflect"
)

type resume struct {
	Name string `json:"name" doc:"my name"`
}

func findDoc(stru interface{}) map[string]string {
	t := reflect.TypeOf(stru).Elem()
	doc := make(map[string]string)

	for i := 0; i < t.NumField(); i++ {
		doc[t.Field(i).Tag.Get("json")] = t.Field(i).Tag.Get("doc")
		log.Println(t.Field(i).Tag.Get("json"))
	}

	return doc

}

func main() {
	var stru resume
	doc := findDoc(&stru)
	fmt.Printf("name：%s\n", doc["name"])
}
