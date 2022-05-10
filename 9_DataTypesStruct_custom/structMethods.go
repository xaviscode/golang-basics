package main

import "fmt"

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

func (s Student) getAverageGrade() float32 {
	sum := 0
	for _, v := range s.grades {
		sum += v
	}
	return float32(sum) / float32(len(s.grades))
}

func (s *Student) getMaxGrade() int {
	curMax := 0

	for _, v := range s.grades {
		if curMax < v {
			curMax = v
		}
	}
	return curMax
}

func main() {

	s1 := Student{"Name", []int{7, 4, 2, 8, 5, 5}, 19}
	s2 := Student{"Name1", []int{5, 7, 8, 7, 4, 9}, 19}
	fmt.Println(s1)

	//s1.age
	s1.getAge()
	s1.setAge(20)

	fmt.Println(s1)

	average := s1.getAverageGrade()
	average2 := s2.getAverageGrade()
	fmt.Println(average, average2)

	max := s1.getMaxGrade()
	max2 := s2.getMaxGrade()
	fmt.Println(max, max2)
}
