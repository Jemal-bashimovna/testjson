package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Students struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Grades Grades `json:"grades"`
}

type Grades struct {
	Math      int `json:"math"`
	Physics   int `json:"physics"`
	Chemistry int `json:"chemistry"`
}

func UnmarshalJson(filename string) ([]Students, error) {
	var student []Students

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error!", r)
		}
	}()

	//file, err := os.Open(filename)
	//if err != nil {
	//	fmt.Println("Error opening file")
	//	return nil, err
	//}
	//defer file.Close()

	fileData, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file")
		return nil, err
	}
	if len(fileData) == 0 {
		return nil, fmt.Errorf("File is empty")
	}

	er := json.Unmarshal(fileData, &student)
	if er != nil {
		return nil, fmt.Errorf("Error processing JSON. %v\n", er)
	}
	fmt.Println("Successfully Unmarshalling")
	fmt.Println()
	fmt.Println("Students who scored more than 90 points in math:")

	for _, val := range student {
		if val.Grades.Math > 90 {
			fmt.Printf("Name: %v, age: %d, score - %d\n", val.Name, val.Age, val.Grades.Math)
		}
	}
	fmt.Println()

	for _, val := range student {
		if val.Age == 20 {
			grade := 0
			grade += val.Grades.Math
			grade += val.Grades.Physics
			grade += val.Grades.Chemistry
			fmt.Println(val.Name, " : ", grade)
		}
	}

	highScore := student[0].Grades.Physics
	person := student[0].Name
	for _, val := range student {
		if highScore < val.Grades.Physics {
			highScore = val.Grades.Physics
			person = val.Name
		}
	}
	fmt.Printf("%s has the highest score in physics : %d\n", person, highScore)

	shortName := student[0].Name
	for _, val := range student {
		if len(shortName) >= len(val.Name) {
			shortName = val.Name
		}
	}

	fmt.Println(strings.ToUpper(shortName))
	return student, nil
}
