package main

import (
	"fmt"
)

type Student struct {
	name   string
	grades []int
	age    int
}

func (s Student) getAge() int {
	return s.age
}

func (s *Student) setAge(age int) {
	s.age = age
}

func (s *Student) getAverageGrades() float32 {
	sum := 0
	for _, v := range s.grades {
		sum += v
	}
	return float32(sum) / float32(len(s.grades))
}

func (s Student) getMaxGrade() int {
	curMax := 0
	for _, v := range s.grades {
		if curMax < v {
			curMax = v
		}
	}
	return curMax
}

func main() {
	s1 := Student{"Ze", []int{14, 18, 19}, 19}
	fmt.Println("age:", s1.getAge())

	s1.setAge(20)
	fmt.Println("age is now:", s1.getAge())

	average := s1.getAverageGrades()
	fmt.Println("average:", average)

	maxGrade := s1.getMaxGrade()
	fmt.Println("max grade:", maxGrade)
}
