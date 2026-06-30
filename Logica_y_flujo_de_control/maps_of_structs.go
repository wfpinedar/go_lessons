package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Read input
	var numStudentsStr string
	var studentData string
	fmt.Scanln(&numStudentsStr)
	fmt.Scanln(&studentData)

	// TODO: Define Student struct here
	type Student struct {
		ID    int
		Grade string
	}

	// TODO: Create a map to store students (name as key, Student struct as value)
	//students :=
	numStudents, _ := strconv.Atoi(numStudentsStr)
	students := make(map[string]Student, numStudents)

	// TODO: Parse the student data and populate the map
	studentIn := strings.Split(studentData, ",")

	for _, st := range studentIn {
		values := strings.Split(st, ":")
		id, _ := strconv.Atoi(values[1])
		students[values[0]] = Student{id, values[2]}
	}

	// TODO: Display all students in alphabetical order by name
	// Crear un slice para las claves
	keys := make([]string, 0, len(students))

	// Popular el slice de claves
	for k := range students {
		keys = append(keys, k)
	}

	// Ordenar las claves alfabéticamente
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%s: ID %d, Grade %s\n", k, students[k].ID, students[k].Grade)
	}

	// TODO: Calculate and display grade statistics
	gradeStats := make(map[string]int)
	for _, valor := range students {
		gradeStats[valor.Grade] += 1
	}
	new_keys := make([]string, 0, len(gradeStats))
	for k, _ := range gradeStats {
		new_keys = append(new_keys, k)
	}

	// Ordenar las claves alfabéticamente
	sort.Strings(new_keys)
	for _, k := range new_keys {
		fmt.Printf("Grade %s: %d students\n", k, gradeStats[k])
	}

	// TODO: Find and display the student with the highest ID
	max := 0
	max_st := ""
	for name, st := range students {
		if st.ID > max {
			max = st.ID
			max_st = name
		}
	}
	fmt.Printf("Highest ID: %s (%d)\n", max_st, max)

	// TODO: Display the total number of students
	fmt.Printf("Total students: %d\n", len(students))
}
