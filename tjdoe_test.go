package tjdoe

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestOutput(t *testing.T) {
	tjdoe := New(1)
	students, _ := tjdoe.BuildScores([]string{"testdata/scores.csv"})
	buffer := bytes.Buffer{}
	tjdoe.OutputAnonymizedScores(students, &buffer)
	wontResults := `id,final score,a01,a02,a03,a04,a05,a06,a07,a08,a09,a10
2245bd5f,B,5,6,7,4,3,1,8,9,8,6
22eb9250,F,4,3,3,2,3,,,,,4
7382d1e7,A,10,10,10,10,10,10,,10,10,
`
	results := string(buffer.Bytes())
	if results != wontResults {
		t.Errorf("OutputAnonymizedScores generates unexpected results, wont %s, got %s", wontResults, results)
	}
}

func TestMapping(t *testing.T) {
	tjdoe := New(1)
	students, _ := tjdoe.BuildScores([]string{"testdata/scores.csv"})
	mappings := tjdoe.buildMappings(students)
	wontLength := 20
	if len(mappings) != wontLength {
		t.Errorf("mapping length did not match, wont %d, got %d", wontLength, len(mappings))
	}
	if mappings[0].String() != `{ from: "123456 Tamada Haruaki", to: "2245bd5f" }` {
		t.Errorf("mappings[0] did not match, got %s", mappings[0].String())
	}
}

func TestCopyDirectories(t *testing.T) {
	// os.RemoveAll("testdata/anonymity")
	tjdoe := New(1)
	students, _ := tjdoe.BuildScores([]string{"testdata/scores.csv"})
	tjdoe.AnonymizeDirectory("testdata/assignments", "testdata/anonymity", students)
	defer os.RemoveAll("testdata/anonymity")

	paths := []string{"0000/2245bd5f/c/hello.c", "0000/2245bd5f/go/hello.go", "0000/2245bd5f/java/HelloWorld_2245bd5f.java", "0000/2245bd5f/node/hello.js", "0000/22eb9250/java/HelloWorld_22eb9250.java"}
	for _, path := range paths {
		_, err := os.Stat(filepath.Join("testdata/anonymity", path))
		if err != nil { // file did not exists
			t.Errorf("copied file did not exist: %s", filepath.Join("testdata/anonymity", path))
		}
	}
}
