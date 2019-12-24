package jjdoes

import (
	"fmt"
	"strconv"
)

type Student struct {
	ID           string
	Name         string
	AnonymizedID string
	FinalScore   string
	Columns      []string
}

func isNumeric(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}

func anonymizeFinalScore(score string) string {
	if isNumeric(score) {
		scoreNumber, err := strconv.Atoi(score)
		if err != nil {
			return score
		}
		if scoreNumber < 60 {
			return "F"
		} else if scoreNumber < 70 {
			return "C"
		} else if scoreNumber < 80 {
			return "B"
		} else if scoreNumber < 90 {
			return "A"
		} else if scoreNumber < 100 {
			return "S"
		}
	}
	return score
}

func BuildMapping(students []Student) map[string]string {
	mapping := map[string]string{}
	for _, student := range students {

	}
	return mapping
}

func Anonymize(students []Student) []Student {
	// TODO implement this function
	return students
}

func BuildScore(scoreFile string) ([]Student, error) {
	return nil, fmt.Errorf("implement please.")
}
