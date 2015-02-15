package models

import (
	"testing"
	"time"
)

func TestSort(t *testing.T) {
	p1 := Post{}
	p1.PubTime = time.Now()
	p1.Title = "1"
	time.Sleep(10 * time.Millisecond)
	p2 := Post{}
	p2.Title = "2"
	p2.PubTime = time.Now()
	time.Sleep(10 * time.Millisecond)
	p3 := Post{}
	p3.Title = "3"
	p3.PubTime = time.Now()
	time.Sleep(10 * time.Millisecond)
	p4 := Post{}
	p4.Title = "4"
	p4.PubTime = time.Now()
	time.Sleep(10 * time.Millisecond)
	p5 := Post{}
	p5.Title = "5"
	p5.PubTime = time.Now()

	ps1 := []Post{p3, p1}
	sort(ps1)
	if !checkSort(ps1) {
		t.Log(ps1)
		t.Fail()
	}

	ps2 := []Post{p3, p1, p2}
	sort(ps2)
	if !checkSort(ps2) {
		t.Log(ps2)
		t.Fail()
	}
	// test1
	ps := []Post{p2, p3, p1, p5, p4}
	sort(ps)
	if !checkSort(ps) {
		t.Log(ps)
		t.Fail()
	}

}

func checkSort(ps []Post) bool {

	for i := 0; i < len(ps); i++ {
		for j := i; j < len(ps); j++ {
			if ps[i].PubTime.After(ps[j].PubTime) {
				return false
			}
		}
	}
	return true
}
