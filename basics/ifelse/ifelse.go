package main

import (
	"errors"
)

func getStudentGrade(marks int) (string, error) {
	if marks < 0 || marks > 100 {
		return "", errors.New("Invalid marks")
	}

	if marks >= 90 {
		return "A", nil
	} else if marks >= 80 {
		return "B", nil
	} else if marks >= 70 {
		return "C", nil
	} else if marks >= 60 {
		return "D", nil
	} else {
		return "F", nil
	}

}

func main() {
	grade, err := getStudentGrade(185)
	if err != nil {
		panic(err)
	}
	println("Grade:", grade)
}
