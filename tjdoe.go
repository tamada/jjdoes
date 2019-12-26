package tjdoe

import (
	"encoding/csv"
	"fmt"
	"io"
	"math/rand"
	"sort"
	"strconv"
	"strings"

	"github.com/seehuhn/mt19937"
)

/*
TJDoe shows core type for processing anonymity of programs.
*/
type TJDoe struct {
	random  *rand.Rand
	mapping []Mapping
}

/*
New creates and returns an instance of TJDoe by initializing given seed.
*/
func New(seed int64) *TJDoe {
	tjdoe := new(TJDoe)
	tjdoe.random = rand.New(mt19937.New())
	tjdoe.random.Seed(seed)
	return tjdoe
}

func convertToMappingSlice(mapping map[string]string) []Mapping {
	results := []Mapping{}
	for key, value := range mapping {
		results = append(results, Mapping{fromID: key, toID: value})
	}
	sort.Slice(results, func(i, j int) bool {
		return len(results[i].fromID) > len(results[j].fromID)
	})
	return results
}

/*
BuildMappings creates Mapping array from given students.
*/
func (tjdoe *TJDoe) buildMappings(students []*Student) []Mapping {
	mapping := map[string]string{}
	for _, student := range students {
		updateMapping(mapping, student.AnonymizedID, student.ID)
		updateMapping(mapping, student.AnonymizedID, student.Name)
		updateMapping(mapping, student.AnonymizedID, strings.ReplaceAll(student.Name, " ", ""))
		updateMapping(mapping, student.AnonymizedID, familyName(student.Name))
		updateMapping(mapping, student.AnonymizedID, fmt.Sprintf("%s %s", student.ID, student.Name))
		updateMapping(mapping, student.AnonymizedID, fmt.Sprintf("%s %s", student.ID, strings.ReplaceAll(student.Name, " ", "")))
		updateMapping(mapping, student.AnonymizedID, fmt.Sprintf("%s %s", student.ID, familyName(student.Name)))
	}
	tjdoe.mapping = convertToMappingSlice(mapping)
	return tjdoe.mapping
}

/*
AnonymizeDirectory copies files in from directory to destination directories with given mapping.
*/
func (tjdoe *TJDoe) AnonymizeDirectory(from, to string, students []*Student) error {
	tjdoe.buildMappings(students)
	return tjdoe.copy(from, to)
}

func createCsvItems(student *Student, labels []string) []string {
	array := []string{student.AnonymizedID, student.AnonymizedFinalScore}
	for _, label := range labels {
		value, ok := student.Scores[label]
		valueString := strconv.Itoa(value)
		if !ok {
			valueString = ""
		}
		array = append(array, valueString)
	}
	return array
}

/*
OutputAnonymizedScores generates score file to destination.
*/
func (tjdoe *TJDoe) OutputAnonymizedScores(students []*Student, dest io.Writer) error {
	header := createHeader(students)
	writer := csv.NewWriter(dest)
	writer.Write(header)
	labels := header[2:]
	for _, student := range students {
		writer.Write(createCsvItems(student, labels))
	}
	writer.Flush()
	return nil
}

func contains(array []string, value string) bool {
	for _, item := range array {
		if item == value {
			return true
		}
	}
	return false
}

func assignmentNames(assignments []string, scores map[string]int) []string {
	for key := range scores {
		if !contains(assignments, key) {
			assignments = append(assignments, key)
		}
	}
	return assignments
}

func createHeader(students []*Student) []string {
	assignments := []string{}
	for _, student := range students {
		assignments = assignmentNames(assignments, student.Scores)
	}
	sort.Slice(assignments, func(i, j int) bool {
		if len(assignments[i]) == len(assignments[j]) {
			return strings.Compare(assignments[i], assignments[j]) <= 0
		}
		return len(assignments[i]) < len(assignments[j])
	})
	header := []string{"id", "final score"}
	header = append(header, assignments...)
	return header
}
