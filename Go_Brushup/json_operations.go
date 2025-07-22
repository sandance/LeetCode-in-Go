package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type response1 struct {
	Page   int
	Fruits []string
}

// Only exported fields will be encoded/decoded in JSON.
// Fields must start with capital letters to be exported.
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	// converts
	// json.Marshal is a function from the encoding/json package used to convert Go data structures into JSON formatted data.
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	response1 := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "punch"},
	}

	res1B, _ := json.Marshal(response1)
	fmt.Println(string(res1B))

	response2 := &response2{
		Page:   2,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res2B, _ := json.Marshal(response2)
	fmt.Println(string(res2B))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var dat map[string]interface{}

	// convert json into struct
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	fmt.Println(dat["num"])

	if err := json.NewDecoder(bytes.NewReader(byt)).Decode(&dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
}
