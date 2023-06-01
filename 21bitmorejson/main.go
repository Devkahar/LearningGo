package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"name"`
	Price    int
	Platform string   `json:"platform"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Learning how to create json data.")
	encodeJSON()
	decodeJSON()

}

func encodeJSON() {
	courses := []course{
		{"React JS", 299, "Udemy", "123@123", []string{"react js", "javascript"}},
		{"Go", 0, "Youtube", "123@1234", []string{"goplers", "backend"}},
		{"Vue JS", 299, "vuejs.org", "123@1234", []string{"Vue JS", "js", "vue"}},
		{"Javascript", 399, "mdn.org", "123@1234", nil},
	}
	jsonData, err := json.MarshalIndent(courses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonData))
}
func decodeJSON() {
	jsonData := []byte(`
	{
		"name": "Vue JS",
		"Price": 299,
		"platform": "vuejs.org",
		"tags": ["Vue JS","js","vue"]
	}
	`)

	validJSON := json.Valid(jsonData)

	if validJSON {
		fmt.Println("JSON IS VALID")
		var decodeCourse course
		json.Unmarshal(jsonData, &decodeCourse)
		fmt.Printf("%#v\n", decodeCourse)
	} else {
		fmt.Println("JSON IS NOT VALID")
	}

	// some cases where you just want to add data to key value pair

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonData, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("The key %v and value %v, and its type is %T\n", k, v, v)
	}
}
