package bench

import "testing"

// path is a text file path.
const path = "./data.txt"
const path2 = "./data2.txt"

// word is a word to find.
const word = "aa"

// want is an expected number.
const want = "1:0,1:10,6:0,6:1"
const want2 = "1:9,2:5,2:6,3:0,3:6,3:7"

func TestFind(t *testing.T) {
	got, err := Find(path, word)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("Find(%q, %q) => %q, want %q", path, word, got, want)
	}
}

func TestFind_noTrailingNewLine(t *testing.T) {
	got, err := Find(path2, word)
	if err != nil {
		t.Fatal(err)
	}
	if got != want2 {
		t.Errorf("Find(%q, %q) => %q, want %q", path2, word, got, want2)
	}
}

func TestFind_emptyWord(t *testing.T) {
	if _, err := Find(path, ""); err == nil {
		t.Error("some kind of error should be returned")
	}
}

func TestFind_emptyResult(t *testing.T) {
	got, err := Find(path, "not_exist_word")
	if err != nil {
		t.Fatal(err)
	}
	if got != "" {
		t.Error("empty value should be returned, ", got)
	}
}

func BenchmarkFind(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Find(path, word)
	}
}

