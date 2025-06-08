Student Marks Management System (in Go)

==========================================
Description:
------------------------------------------
A simple CLI-based Student Marks Management System developed in Golang.

It uses:
- Structs: to represent each student's data.
- Slices: to store a list of students.
- Maps: to store subject-wise scores.

==========================================
Features:
------------------------------------------
1. Add a new student with subject scores.
2. Calculate average marks for each student.
3. Sort all students based on their average marks.
4. Display the list of students with full details.

==========================================
Code Overview:
------------------------------------------

1. Struct: 
   - Student (Name, RollNo, map of Scores, Average)

2. Functions:
   - addStudent(name, rollNo, scores): Adds a new student.
   - calculateAverage(scores): Computes average of given marks.
   - sortByGrade(): Sorts students in descending order of average.
   - displayStudents(): Displays all students and their scores.

3. Data:
   - students []Student: Slice of student records.

4. Example Subjects:
   - Math, Science, English

==========================================
How to Run:
------------------------------------------
1. Save the main Go code in a file named: main.go
2. Use terminal to run:
   go run main.go

==========================================
Sample Output:
------------------------------------------
✅ Student added: Alice
✅ Student added: Bob
✅ Student added: Charlie

📋 List of Students:
Name: Alice | Roll: 101 | Avg: 84.33 | Scores: map[English:78 Math:85 Science:90]
Name: Bob | Roll: 102 | Avg: 75.00 | Scores: map[English:70 Math:75 Science:80]
Name: Charlie | Roll: 103 | Avg: 91.67 | Scores: map[English:88 Math:95 Science:92]

🔃 Students sorted by grade.

📋 List of Students (Sorted):
Name: Charlie | Roll: 103 | Avg: 91.67 | Scores: map[English:88 Math:95 Science:92]
Name: Alice   | Roll: 101 | Avg: 84.33 | Scores: map[English:78 Math:85 Science:90]
Name: Bob     | Roll: 102 | Avg: 75.00 | Scores: map[English:70 Math:75 Science:80]

==========================================
Future Improvements (optional):
------------------------------------------
- Add grade (A/B/C) based on average
- Use file/database storage
- Take user input via console
- Add web interface using Go’s net/http
