package golister_test

import (
	"testing"

	"github.com/mekramy/golister"
)

func TestEmptyLister(t *testing.T) {
	lister := golister.New().SetTotal(0)

	if v := lister.Page(); v != 0 {
		t.Errorf("expected Page() to be 0, got %d", v)
	}

	if v := lister.Pages(); v != 0 {
		t.Errorf("expected Pages() to be 0, got %d", v)
	}

	if v := lister.Total(); v != 0 {
		t.Errorf("expected Total() to be 0, got %d", v)
	}

	if v := lister.From(); v != 0 {
		t.Errorf("expected From() to be 0, got %d", v)
	}

	if v := lister.To(); v != 0 {
		t.Errorf("expected To() to be 0, got %d", v)
	}

	expected := ` ORDER BY "id" ASC LIMIT 25 OFFSET 0`
	if sorts := lister.SQLSortOrder(); sorts != expected {
		t.Errorf(`expected SQLSortOrder() to be "%s", got "%s"`, expected, sorts)
	}
}

func TestLister(t *testing.T) {
	lister := golister.New(
		golister.WithLimits(20, 10, 20, 30),
		golister.WithSorts("id", "id", "name", "mobile"),
	).
		AddSort("name", golister.ParseOrder(-1)).
		AddSort("mobile", "desc").
		SetLimit(100).
		SetPage(100).
		SetTotal(101)

	if v := lister.Page(); v != 6 {
		t.Errorf("expected Page() to be 6, got %d", v)
	}

	if v := lister.Pages(); v != 6 {
		t.Errorf("expected Pages() to be 6, got %d", v)
	}

	if v := lister.Limit(); v != 20 {
		t.Errorf("expected Limit() to be 20, got %d", v)
	}

	if v := lister.Total(); v != 101 {
		t.Errorf("expected Total() to be 101, got %d", v)
	}

	if v := lister.From(); v != 100 {
		t.Errorf("expected From() to be 100, got %d", v)
	}

	if v := lister.To(); v != 101 {
		t.Errorf("expected To() to be 101, got %d", v)
	}

	expected := ` ORDER BY "name" DESC, "mobile" DESC LIMIT 20 OFFSET 100`
	if sorts := lister.SQLSortOrder(); sorts != expected {
		t.Errorf(`expected SQLSortOrder() to be "%s", got "%s"`, expected, sorts)
	}
}
