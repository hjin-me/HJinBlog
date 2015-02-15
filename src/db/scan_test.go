package db

import "testing"

func TestScan(t *testing.T) {
	err := Connect("10.211.55.8:6379")
	if err != nil {
		t.Fatal(err)
	}
	count := 0
	for _ = range Scan() {
		count++
	}

	if err := Err(); err != nil {
		t.Error(err)
	}

	if count != 610 {
		t.Error("count is not 610")
	}

}
