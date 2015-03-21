package bench

import "testing"

// path is a text file path.
const path = "./sample.txt"

// want is an expected number.
const want = 54

func TestCount(t *testing.T) {
	got, err := Count(path)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("Count(%q) => %d, want %d", path, got, want)
	}
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(path)
	}
}
