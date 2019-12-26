package tjdoe

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
Mapping represents id mapping.
*/
type Mapping struct {
	fromID string
	toID   string
}

/*
Student shows a student score of grade.
*/
type Student struct {
	ID                   string
	Name                 string
	AnonymizedID         string
	FinalScore           string
	AnonymizedFinalScore string
	Scores               map[string]int
}

func (m *Mapping) String() string {
	return fmt.Sprintf(`{ from: "%s", to: "%s" }`, m.fromID, m.toID)
}

func (s *Student) String() string {
	return fmt.Sprintf("%s,%s,%s", s.ID, s.Name, s.FinalScore)
}

/*
AnonymizedString returns anonymized information of student.
*/
func (s *Student) AnonymizedString() string {
	return fmt.Sprintf("%s,%s", s.AnonymizedID, s.AnonymizedFinalScore)
}

func isNumeric(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}

func anonymizeFinalScore(score string) string {
	if isNumeric(score) {
		scoreNumber, _ := strconv.Atoi(score)
		switch {
		case scoreNumber < 60:
			return "F"
		case scoreNumber < 70:
			return "D"
		case scoreNumber < 80:
			return "C"
		case scoreNumber < 90:
			return "B"
		default:
			return "A"
		}
	}
	return score
}

func anonymizeID(tjdoe *TJDoe, id string) string {
	return fmt.Sprintf("%0x", tjdoe.random.Int31())
}

func updateMapping(mapping map[string]string, toID, fromID string) {
	val, ok := mapping[fromID]
	if ok {
		toID = fmt.Sprintf("%s, %s", val, toID)
	}
	mapping[fromID] = toID
}

func familyName(name string) string {
	return strings.Split(name, " ")[0]
}

func isValidID(newID string, ids []string) bool {
	for _, id := range ids {
		if id == newID {
			return false
		}
	}
	return true
}

func anonymizeIDs(tjdoe *TJDoe, student *Student, ids []string) string {
	newID := anonymizeID(tjdoe, student.ID)
	for !isValidID(newID, ids) {
		newID = anonymizeID(tjdoe, student.ID)
	}
	student.AnonymizedID = newID
	student.AnonymizedFinalScore = anonymizeFinalScore(student.FinalScore)
	return newID
}

func (tjdoe *TJDoe) anonymize(students []*Student) []*Student {
	ids := []string{}
	for _, student := range students {
		id := anonymizeIDs(tjdoe, student, ids)
		ids = append(ids, id)
	}
	sort.Slice(students, func(i, j int) bool {
		return strings.Compare(students[i].AnonymizedID, students[j].AnonymizedID) < 0
	})
	return students
}

func buildScore(fileName string) ([]*Student, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	header, err := reader.Read()
	if err != nil {
		return nil, err
	}
	return buildStudents(reader, header), nil
}

func buildStudent(header, records []string) *Student {
	student := &Student{ID: records[0], Name: records[1], FinalScore: records[2], Scores: map[string]int{}}
	for i := range header[3:] {
		number, err := strconv.Atoi(records[3+i])
		if err == nil {
			student.Scores[header[i+3]] = number
		}
	}
	return student
}

func buildStudents(reader *csv.Reader, header []string) []*Student {
	results := []*Student{}
	for {
		records, err := reader.Read()
		if err == io.EOF {
			break
		}
		results = append(results, buildStudent(header, records))
	}
	return results
}

/*
BuildScores creates array of Student from given score files.
*/
func (tjdoe *TJDoe) BuildScores(scoreFiles []string) ([]*Student, error) {
	var err error
	results := []*Student{}
	for _, file := range scoreFiles {
		students, err1 := buildScore(file)
		if err1 != nil {
			err = err1
		} else {
			results = append(results, students...)
		}
	}
	return tjdoe.anonymize(results), err
}
