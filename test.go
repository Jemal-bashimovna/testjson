package main

type Person struct {
	Name      *string  `json:"name"`
	Age       int      `json:"age"`
	IsStudent bool     `json:"is_student"`
	Courses   []string `json:"courses"`
}

func test() {
	// p := Person{
	// 	Name:      nil,
	// 	Age:       25,
	// 	IsStudent: true,
	// 	Courses:   []string{"math", "history", "chemistry"},
	// }

	// jsonData, err := json.Marshal(p) // return []byte and error
	// if err != nil {
	// 	panic(err)
	// }

	// file, err := os.Create("newFile.json")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer file.Close()

	// file.WriteString(string(jsonData))
	// println()

	// jsonFile, err := os.ReadFile("person.json")
	// if err != nil {
	// 	fmt.Println("Can't read file ", err)
	// }

	// var person Person
	// er := json.Unmarshal([]byte(jsonFile), &person)
	// if er != nil {
	// 	panic(er)
	// }
	// fmt.Println(person)
}
