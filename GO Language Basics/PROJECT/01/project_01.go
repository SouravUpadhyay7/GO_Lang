package main

import (
	"fmt"
	"sort"
)

// Struct for each student
type Student struct {
	Name    string
	RollNo  int
	Scores  map[string]float64
	Average float64
}

// Slice to store students
var students []Student

// Add a student
func addStudent(name string, rollNo int, scores map[string]float64) {
	avg := calculateAverage(scores)
	newStudent := Student{
		Name:    name,
		RollNo:  rollNo,
		Scores:  scores,
		Average: avg,
	}
	students = append(students, newStudent)
	fmt.Println("✅ Student added:", name)
}

// Calculate average of a student's scores
func calculateAverage(scores map[string]float64) float64 {
	sum := 0.0
	for _, val := range scores {
		sum += val
	}
	return sum / float64(len(scores))
}

// Sort students by average marks
func sortByGrade() {
	sort.Slice(students, func(i, j int) bool {
		return students[i].Average > students[j].Average // descending
	})
	fmt.Println("🔃 Students sorted by grade.")
}

// Display all students
func displayStudents() {
	fmt.Println("📋 List of Students:")
	for _, s := range students {
		fmt.Printf("Name: %s | Roll: %d | Avg: %.2f | Scores: %v\n", s.Name, s.RollNo, s.Average, s.Scores)
	}
}

func main() {
	// Example students
	addStudent("Alice", 101, map[string]float64{
		"Math":    85,
		"Science": 90,
		"English": 78,
	})

	addStudent("Bob", 102, map[string]float64{
		"Math":    75,
		"Science": 80,
		"English": 70,
	})

	addStudent("Charlie", 103, map[string]float64{
		"Math":    95,
		"Science": 92,
		"English": 88,
	})

	displayStudents()

	sortByGrade()
	fmt.Println("\nAfter Sorting:")
	displayStudents()
}
// This program manages a list of students, allowing you to add students with their scores,
// calculate their average scores, sort them by grades, and display the list of students.