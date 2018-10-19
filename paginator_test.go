package gopager

import (
	"strings"
	"testing"
)

func TestSliderCloseToBeginning(t *testing.T) {
	paginator := NewLengthAwarePaginator(make([]int, 100), 100, 5, 1, nil)

	elements := paginator.Elements()

	if len(elements[0]) != 8 {
		t.Errorf("The first part should contains 8 links, but got %v", len(elements[0]))
	}

	if len(elements[1]) != 0 {
		t.Errorf("The slider part should be empty, but got %v", len(elements[0]))
	}

	if len(elements[2]) != 2 {
		t.Errorf("The last part should contains 2 links, but got %v", len(elements[0]))
	}
}

func TestSliderCloseToEnding(t *testing.T) {
	paginator := NewLengthAwarePaginator(make([]int, 100), 100, 5, 20, nil)

	elements := paginator.Elements()

	if len(elements[0]) != 2 {
		t.Errorf("The first part should contains 2 links, but got %v", len(elements[0]))
	}

	if len(elements[1]) != 0 {
		t.Errorf("The slider part should be empty, but got %v", len(elements[0]))
	}

	if len(elements[2]) != 9 {
		t.Errorf("The last part should contains 9 links, but got %v", len(elements[0]))
	}
}

func TestFullSlider(t *testing.T) {
	paginator := NewLengthAwarePaginator(make([]int, 100), 100, 5, 10, nil)

	elements := paginator.Elements()

	if len(elements[0]) != 2 {
		t.Errorf("The first part should contains 2 links, but got %v", len(elements[0]))
	}

	if len(elements[1]) != 7 {
		t.Errorf("The slider part should contains 7 links, but got %v", len(elements[0]))
	}

	if len(elements[2]) != 2 {
		t.Errorf("The last part should contains 2 links, but got %v", len(elements[0]))
	}
}

func TestPath(t *testing.T) {
	paginator1 := NewPaginator(make([]int, 20), 20, 10, 1, nil)

	elements1 := paginator1.Elements()

	if elements1[1] != "/?page=1" {
		t.Errorf("The path should be /?page=1, but got %v", elements1[0])
	}

	paginator2 := NewPaginator(make([]int, 20), 20, 10, 1, map[string]string{
		"path": "/foo/bar",
	})

	elements2 := paginator2.Elements()

	if elements2[1] != "/foo/bar?page=1" {
		t.Errorf("The path should be /?page=1, but got %v", elements2[0])
	}
}

func TestPageName(t *testing.T) {
	paginator1 := NewPaginator(make([]int, 20), 20, 10, 1, nil)

	elements1 := paginator1.Elements()

	if ! strings.Contains(elements1[1], "?page=1") {
		t.Errorf("The page name isn't `page`")
	}

	paginator2 := NewPaginator(make([]int, 20), 20, 10, 1, map[string]string{
		"pageName": "p",
	})

	elements2 := paginator2.Elements()

	if ! strings.Contains(elements2[1], "?p=1") {
		t.Errorf("The page name isn't `p`")
	}
}

func TestAppends(t *testing.T) {
	paginator := NewPaginator(make([]int, 20), 20, 10, 1, nil)

	paginator.Appends(map[string][]string{
		"keyword": {"andy"},
		"names":   {"tom", "jack"},
	})

	elements := paginator.Elements()

	if ! strings.Contains(elements[1], "keyword=andy") {
		t.Errorf("The path doesn't contains `keyword=andy`")
	}
}
