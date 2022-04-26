package main

// siapkan struct student
type Student struct {
	Id    string
	Name  string
	Grade int32
}

// variable global penampung student
var students = []*Student{}

// func GetStudent
func GetStudents() []*Student {
	return students
}

// func SelectStudent
func SelectStudent(id string) *Student {
	for _, val := range students {
		if val.Id == id {
			return val
		}
	}

	return nil
}

// func init untuk data dummy, karena func ini pertama kali akan dipanggil ketika package diimport atau dirun
func init() {
	students = append(students, &Student{Id: "1", Name: "Budi", Grade: 10})
	students = append(students, &Student{Id: "2", Name: "Joko", Grade: 9})
	students = append(students, &Student{Id: "3", Name: "Sri", Grade: 8})
}
