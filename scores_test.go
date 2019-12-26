package tjdoe

import "testing"

func TestBuildScores(t *testing.T) {
	tjdoe := New(1)
	students, err := tjdoe.BuildScores([]string{"testdata/mapping.csv"})
	if err != nil {
		t.Error(err)
		return
	}
	if len(students) != 3 {
		t.Errorf("rows count did not match, wont 3, got %d", len(students))
	}
	s := students[0]
	if s.String() != "123456,Tamada Haruaki,87" {
		t.Errorf("mapping.csv read error, wont %s, got %v", "123456,Tamada Haruaki,87", s.String())
	}
	if s.AnonymizedString() != "2245bd5f,B" {
		t.Errorf("anonymized string did not match, wont %s, got %v", "2245bd5f,B", s.AnonymizedString())
	}
	if len(s.Scores) != 10 {
		t.Errorf("score count did not match, wont 10, got %d", len(s.Scores))
	}
}

func TestAnonymizedFinalScore(t *testing.T) {
	testdata := []struct {
		giveScore string
		wontScore string
	}{
		{"K", "K"},
		{"/", "/"},
		{"120", "A"},
		{"90", "A"},
		{"80", "B"},
		{"70", "C"},
		{"60", "D"},
		{"54", "F"},
		{"-3", "F"},
	}
	for _, td := range testdata {
		gotScore := anonymizeFinalScore(td.giveScore)
		if gotScore != td.wontScore {
			t.Errorf("results of anonymizedFinalScore(%s) did not match, wont %s, got %s", td.giveScore, td.wontScore, gotScore)
		}
	}
}

func TestFamilyName(t *testing.T) {
	testdata := []struct {
		giveString string
		wontString string
	}{
		{"Tamada Haruaki", "Tamada"},
		{"Tamada", "Tamada"},
	}
	for _, td := range testdata {
		gotString := familyName(td.giveString)
		if gotString != td.wontString {
			t.Errorf("familyName(%s) wont %s, but got %s", td.giveString, td.wontString, gotString)
		}
	}
}
