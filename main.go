package main

import (
	"a21hc3NpZ25tZW50/helper"
	"a21hc3NpZ25tZW50/model"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type StudentManager interface {
	Login(id string, name string) error
	Register(id string, name string, studyProgram string) error
	GetStudyProgram(code string) (string, error)
	ModifyStudent(name string, fn model.StudentModifier) error
}

type InMemoryStudentManager struct {
	students             []model.Student
	studentStudyPrograms map[string]string
}

func NewInMemoryStudentManager() *InMemoryStudentManager {
	return &InMemoryStudentManager{
		students: []model.Student{
			{
				ID:           "A12345",
				Name:         "Aditira",
				StudyProgram: "TI",
			},
			{
				ID:           "B21313",
				Name:         "Dito",
				StudyProgram: "TK",
			},
			{
				ID:           "A34555",
				Name:         "Afis",
				StudyProgram: "MI",
			},
		},
		studentStudyPrograms: map[string]string{
			"TI": "Teknik Informatika",
			"TK": "Teknik Komputer",
			"SI": "Sistem Informasi",
			"MI": "Manajemen Informasi",
		},
	}
}

func (sm *InMemoryStudentManager) GetStudents() []model.Student { //mengembalikan semua data mahasiswa
	return sm.students // TODO: replace this
}

func (sm *InMemoryStudentManager) Login(id string, name string) (string, error) { // login sesuai id dan nama mahasiswa terdaftar
	for _, student := range sm.students {
		if student.ID == id && student.Name == name {
			return fmt.Sprintf("Login berhasil: %s", name), nil
		}
	}
	return "", errors.New("Login gagal: data mahasiswa tidak ditemukan") // TODO: replace this
}

func (sm *InMemoryStudentManager) Register(id string, name string, studyProgram string) (string, error) { // registrasi mahasiswa baru
	if id == "" || name == "" || studyProgram == "" {
		return "", errors.New("ID, Name or StudyProgram is undefined!")
	}
	for _, student := range sm.students {
		if student.ID == id {
			return "", errors.New("Registrasi gagal: id sudah digunakan")
		}
	}
	if _, ok := sm.studentStudyPrograms[studyProgram]; !ok {
		return "", errors.New(fmt.Sprintf("Study program %s is not found", studyProgram))
	}
	sm.students = append(sm.students, model.Student{
		ID:           id,
		Name:         name,
		StudyProgram: studyProgram,
	})
	return fmt.Sprintf("Registrasi berhasil: %s (%s)", name, studyProgram), nil // TODO: replace this
}

func (sm *InMemoryStudentManager) GetStudyProgram(code string) (string, error) { // return nama program studi berdasarkan kode program studi
	if studyProgram, ok := sm.studentStudyPrograms[code]; ok {
		return studyProgram, nil
	}
	return "", errors.New("Kode program studi tidak ditemukan") // TODO: replace this
}

func (sm *InMemoryStudentManager) ModifyStudent(name string, fn model.StudentModifier) (string, error) { // memodifikasi data mahasiswa
	for i, student := range sm.students {
		if student.Name == name {
			if err := fn(&sm.students[i]); err != nil {
				return "", err
			}
			return "Program studi mahasiswa berhasil diubah.", nil
		}
	}
	return "", errors.New("Mahasiswa tidak ditemukan.") // TODO: replace this
}

func (sm *InMemoryStudentManager) ChangeStudyProgram(studyProgram string) model.StudentModifier { // mengubah program studi mahasiswa
	return func(s *model.Student) error {
		if _, ok := sm.studentStudyPrograms[studyProgram]; !ok {
			return errors.New("Kode program studi tidak ditemukan")
		}
		s.StudyProgram = studyProgram
		return nil
	}
}

func main() {
	manager := NewInMemoryStudentManager()

	for {
		helper.ClearScreen()
		students := manager.GetStudents()
		for _, student := range students {
			fmt.Printf("ID: %s\n", student.ID)
			fmt.Printf("Name: %s\n", student.Name)
			fmt.Printf("Study Program: %s\n", student.StudyProgram)
			fmt.Println()
		}

		fmt.Println("Selamat datang di Student Portal!")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Get Study Program")
		fmt.Println("4. Modify Student")
		fmt.Println("5. Exit")

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Pilih menu: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			helper.ClearScreen()
			fmt.Println("=== Login ===")
			fmt.Print("ID: ")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)

			fmt.Print("Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			msg, err := manager.Login(id, name)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			}
			fmt.Println(msg)
			helper.Delay(5)
		case "2":
			helper.ClearScreen()
			fmt.Println("=== Register ===")
			fmt.Print("ID: ")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)

			fmt.Print("Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Study Program Code (TI/TK/SI/MI): ")
			code, _ := reader.ReadString('\n')
			code = strings.TrimSpace(code)

			msg, err := manager.Register(id, name, code)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			}
			fmt.Println(msg)
			helper.Delay(5)
		case "3":
			helper.ClearScreen()
			fmt.Println("=== Get Study Program ===")
			fmt.Print("Program Code (TI/TK/SI/MI): ")
			code, _ := reader.ReadString('\n')
			code = strings.TrimSpace(code)

			if studyProgram, err := manager.GetStudyProgram(code); err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			} else {
				fmt.Printf("Program Studi: %s\n", studyProgram)
			}
			helper.Delay(5)
		case "4":
			helper.ClearScreen()
			fmt.Println("=== Modify Student ===")
			fmt.Print("Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Program Studi Baru (TI/TK/SI/MI): ")
			code, _ := reader.ReadString('\n')
			code = strings.TrimSpace(code)

			msg, err := manager.ModifyStudent(name, manager.ChangeStudyProgram(code))
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			}
			fmt.Println(msg)
			helper.Delay(5)
		case "5":
			helper.ClearScreen()
			fmt.Println("Goodbye!")
			return
		default:
			helper.ClearScreen()
			fmt.Println("Pilihan tidak valid!")
			helper.Delay(5)
		}

		fmt.Println()
	}
}
