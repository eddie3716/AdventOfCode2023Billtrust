package main

import (
	"testing"
)

func TestFitsPattern1(t *testing.T) {
	pattern := Pattern{1, 3, 7}
	template := Template("???.###")
	set := NewSet(pattern, template)
	if !set.PatternFits(0, 0) {
		t.Fatalf(`PatternFits(0, 0) = false, want true`)
	}
	if !set.PatternFits(0, 1) {
		t.Fatalf(`PatternFits(0, 1) = false, want true`)
	}
	if !set.PatternFits(0, 2) {
		t.Fatalf(`PatternFits(0, 2) = false, want true`)
	}
	if set.PatternFits(0, 3) {
		t.Fatalf(`PatternFits(0, 3) = true, want false`)
	}
	if set.PatternFits(0, 4) {
		t.Fatalf(`PatternFits(0, 4) = true, want false`)
	}
	if set.PatternFits(0, 5) {
		t.Fatalf(`PatternFits(0, 5) = true, want false`)
	}
	if set.PatternFits(0, 6) {
		t.Fatalf(`PatternFits(0, 6) = true, want false`)
	}

	if !set.PatternFits(1, 0) {
		t.Fatalf(`PatternFits(0, 0) = false, want true`)
	}
	if set.PatternFits(1, 1) {
		t.Fatalf(`PatternFits(0, 1) = true, want false`)
	}
	if set.PatternFits(1, 2) {
		t.Fatalf(`PatternFits(0, 2) = true, want false`)
	}
	if set.PatternFits(1, 3) {
		t.Fatalf(`PatternFits(0, 3) = true, want false`)
	}
	if !set.PatternFits(1, 4) {
		t.Fatalf(`PatternFits(0, 4) = false, want true`)
	}
	if set.PatternFits(1, 5) {
		t.Fatalf(`PatternFits(0, 5) = true, want false`)
	}
	if set.PatternFits(1, 6) {
		t.Fatalf(`PatternFits(0, 6) = true, want false`)
	}

	if set.PatternFits(2, 0) {
		t.Fatalf(`PatternFits(0, 0) = true, want false`)
	}
	if set.PatternFits(2, 1) {
		t.Fatalf(`PatternFits(0, 1) = true, want false`)
	}
	if set.PatternFits(2, 2) {
		t.Fatalf(`PatternFits(0, 2) = true, want false`)
	}
	if set.PatternFits(2, 3) {
		t.Fatalf(`PatternFits(0, 3) = true, want false`)
	}
	if set.PatternFits(2, 4) {
		t.Fatalf(`PatternFits(0, 4) = true, want false`)
	}
	if set.PatternFits(2, 5) {
		t.Fatalf(`PatternFits(0, 5) = true, want false`)
	}
	if set.PatternFits(2, 6) {
		t.Fatalf(`PatternFits(0, 6) = true, want false`)
	}
}

func TestFitsPattern2(t *testing.T) {
	pattern := Pattern{3}
	template := Template("?###?????#?#")
	set := NewSet(pattern, template)
	if set.PatternFits(0, 0) {
		t.Fatalf(`PatternFits(0, 0) = true, want false`)
	}
	if !set.PatternFits(0, 1) {
		t.Fatalf(`PatternFits(0, 1) = false, want true`)
	}
	if set.PatternFits(0, 2) {
		t.Fatalf(`PatternFits(0, 2) = true, want false`)
	}
	if set.PatternFits(0, 3) {
		t.Fatalf(`PatternFits(0, 3) = true, want false`)
	}
	if set.PatternFits(0, 4) {
		t.Fatalf(`PatternFits(0, 4) = true, want false`)
	}
	if !set.PatternFits(0, 5) {
		t.Fatalf(`PatternFits(0, 5) = false, want true`)
	}
	if set.PatternFits(0, 6) {
		t.Fatalf(`PatternFits(0, 6) = true, want false`)
	}
	if !set.PatternFits(0, 7) {
		t.Fatalf(`PatternFits(0, 7) = false, want true`)
	}
	if set.PatternFits(0, 8) {
		t.Fatalf(`PatternFits(0, 8) = true, want false`)
	}
	if !set.PatternFits(0, 9) {
		t.Fatalf(`PatternFits(0, 9) = false, want true`)
	}
	if set.PatternFits(0, 10) {
		t.Fatalf(`PatternFits(0, 10) = true, want false`)
	}
	if set.PatternFits(0, 11) {
		t.Fatalf(`PatternFits(0, 11) = true, want false`)
	}
}

// func TestFindCombinations(t *testing.T) {
// 	pattern := Pattern{1, 1, 3}
// 	template := Template("???.###")
// 	set := Set{pattern, template}
// 	c := set.FindCombinations(0, 0, 0)
// 	if c != 1 {
// 		t.Fatalf(`FindCombinations(0, 0) = %d, want 1`, c)
// 	}
// }

func TestFindCombinations2(t *testing.T) {
	pattern := Pattern{1, 1, 3}
	template := Template(".??..??...?##.")
	set := NewSet(pattern, template)
	c := set.FindCombos()
	if c != 4 {
		t.Fatalf(`FindCombinations2(0, 0) = %d, want 1`, c)
	}
}

func TestFindCombinations3(t *testing.T) {
	pattern := Pattern{1, 1, 7, 1, 1}
	template := Template("????????#?#????#.??")
	set := NewSet(pattern, template)
	c := set.FindCombos()
	if c != 45 {
		t.Fatalf(`FindCombinations3(0, 0) = %d, want 1`, c)
	}
}
