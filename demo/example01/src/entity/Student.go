package entity

import (
	"fmt"
	"math/rand"
)

type Student struct {
	Id   int
	Name string
}

func (s Student) String() string {
	return fmt.Sprintf("{ %d: %s }", s.Id, s.Name)
}

type Students []Student

func (arr Students) Len() int {
	return len(arr)
}

func (arr Students) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func (arr Students) Less(i, j int) bool {
	return arr[i].Id < arr[j].Id
}

func GetStudents() Students {
	var arr = Students{}
	for i := 0; i < 10; i++ {
		s := Student{
			Id:   rand.Intn(1000),
			Name: fmt.Sprintf("Jay%d", rand.Intn(10)),
		}
		arr = append(arr, s)
	}
	return arr
}
