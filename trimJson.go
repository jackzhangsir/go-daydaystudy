package main

import (
	"encoding/json"
	"fmt"
)

type T1 struct {
	Field1 string `json:"field_1"`
	Field2 []int  `json:"field_2"`
	Field3 int    `json:"field_3"`
}
type T2 struct {
	Field1 string `json:"field_1"`
	Field2 []int  `json:"field_2"`
	Field3 int    `json:"field_3"`
	T1     T1     `json:"t_1"`
}
type ComplexStruct struct {
	T1
	Array []T1 `json:"array"`
	T2    T2   `json:"t_2"`
}

type obj map[string]interface{}

func trim(source obj, target obj) obj {
	out := make(obj)
	for k, v := range source {
		if vv, find := target[k]; find {
			switch vv.(type) {
			case obj:
				out[k] = make(obj)
				if len(vv.(obj)) == 0 {
					out[k] = v
				} else {
					out[k] = trim(v.(map[string]interface{}), vv.(obj))
				}
			case []obj:
				if len(vv.([]obj)) == 0 {
					out[k] = v
				} else {
					out[k] = make([]obj, 0)
					for i := range v.([]interface{}) {
						o := trim(v.([]interface{})[i].(map[string]interface{}), vv.([]obj)[0])
						out[k] = append(out[k].([]obj), o)
					}
				}
			default:
				out[k] = v
			}
		}
	}
	return out
}
func main() {
	t := ComplexStruct{
		T1: T1{
			Field1: "aaaa",
			Field2: []int{1, 2, 3, 4},
			Field3: 12,
		},
		Array: []T1{
			{
				Field1: "cccc",
				Field2: []int{4, 5, 6},
				Field3: 34,
			},
			{
				Field1: "dddd",
				Field2: []int{14, 15, 16},
				Field3: 35,
			},
		},
		T2: T2{
			Field1: "eeee",
			Field2: []int{123, 231},
			Field3: 1234,
			T1: T1{
				Field1: "aaaa",
				Field2: []int{1, 2, 3, 4},
				Field3: 12,
			},
		},
	}
	source := make(obj)
	bytes, _ := json.Marshal(t)
	json.Unmarshal(bytes, &source)


	//
	target := obj{
		"array": []obj{
			{
				"field_1": "",
				"field_3": "",
			},
		},
		"t_2": obj{
			"t_1": obj{},
		},
	}
	fmt.Println(source)
	out := trim(source, target)

	fmt.Println(out)
	bytes, _ = json.MarshalIndent(out, "", "  ")
	fmt.Println(string(bytes))
}