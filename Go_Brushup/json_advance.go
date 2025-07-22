package main

import (
	"encoding/json"
	"fmt"
)

/*
To decode JSON data we use the Unmarshal function.

func Unmarshal(data []byte, v interface{}) error

# We must first create a place where the decoded data will be stored

var m Message

and call json.Unmarshal, passing it a []byte of JSON data and a pointer to m

err := json.Unmarshal(b, &m)
*/

type Message struct {
	Name string
	Body string
	Time int64
}

func main() {

	b := []byte(`{"Name":"Bob","Food":"Pickle"}`)
	//m := Message{
	//	Name: "Alice",
	//	Body: "Hello",
	//	Time: 1294706395881547000,
	//}
	var m Message
	err := json.Unmarshal(b, &m)
	if err != nil {
		panic(err)
	}
	fmt.Println(m)
	/*
		Unmarshal will decode only the fields that it can find in the destination type.
		In this case, only the Name field of m will be populated, and the Food field will be ignored.
		This behavior is particularly useful when you wish to pick only a few specific fields out of a large JSON blob.
		It also means that any unexported fields in the destination struct will be unaffected by Unmarshal.
	*/

	// But what if you don’t know the structure of your JSON data beforehand?

	jsonData := `{
		"name": "Alice",
		"age": 30,
		"height": 5.5,
		"active": true,
		"hobbies": ["reading", "cycling"],
		"address": {
			"city": "Austin",
			"zip": "73301"
		}
	}`

	var data interface{}
	err = json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		panic(err)
	}

	if obj, ok := data.(map[string]interface{}); ok {
		for key, value := range obj {
			fmt.Printf("Key: %s", key)
			switch v := value.(type) {
			case string:
				fmt.Printf("String value: %s\n", v)
			case float64:
				fmt.Printf("Number value: %f\n", v)
			case bool:
				fmt.Printf("Boolean value: %t\n", v)
			case []interface{}:
				fmt.Printf("Array value: %v\n", v)
			case map[string]interface{}:
				fmt.Println("Object value:")
				for subKey, subVal := range v {
					fmt.Printf("    %s: %v\n", subKey, subVal)
				}
			default:
				fmt.Println("Unknown type")
			}
		}
	} else {
		fmt.Println("Expected a json object")
	}
}
