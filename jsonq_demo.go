package main

import gojsonq "github.com/thedevsaddam/gojsonq/v2"

func main() {
	const json = `{"name":{"first":"Tom","last":"Hanks"},"age":61}`
	name := gojsonq.New().FromString(json).Find("name.first")
	println(name.(string)) // Tom
}