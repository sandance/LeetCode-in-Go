package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/*
✅ Use json.NewEncoder and json.NewDecoder when:
Scenario									Description
Working with streams						The data source is a stream (e.g., a file, network connection, or os.Stdin).
Reading/writing large JSON					You want to avoid loading large JSON blobs entirely into memory.
Multiple JSON objects in a stream			Useful for processing sequences of JSON objects line-by-line.
Better performance with buffered writers	Encoder works with io.Writer, which can be a file, buffer, or network connection.



🆚 Summary 					Comparison:
Feature					Marshal/Unmarshal					Encoder/Decoder
Works on				[]byte, string						io.Writer, io.Reader
Suitable for			In-memory data						Streaming/large/multiple objects
Output/Input type		Byte slices ([]byte)				Streams (file, socket, etc.)
Multiple JSON objects?	No (requires manual splitting)		Yes
*/
/*
func main() {
	dec := json.NewDecoder(os.Stdin)
	enc := json.NewEncoder(os.Stdout)

	for {
		var m map[string]interface{}
		if err := dec.Decode(&m); err != nil {
			log.Println(err)
			return
		}

		for k := range m {
			if k != "Name" {
				delete(m, k)
			}
		}

		if err := enc.Encode(m); err != nil {
			log.Println(err)
		}
	}

}*/

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func main() {
	user := User{
		Name:  "John",
		Email: "john@example.com",
		Age:   20,
	}

	file, err := os.Create("user.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(user); err != nil {
		panic(err)
	}
	fmt.Println("Successfully written user.json")

	readFile, err := os.Open("user.json")
	if err != nil {
		panic(err)
	}
	defer readFile.Close()

	var decodeUser User
	decoder := json.NewDecoder(readFile)
	if err := decoder.Decode(&decodeUser); err != nil {
		panic(err)
	}

	fmt.Printf("Decoded User: %+v\n", decodeUser)
}
