package golister_test

import (
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/mekramy/golister"
)

func TestParamConstructor(t *testing.T) {
	params := golister.ListerParams{
		Page:   3,
		Limit:  50,
		Search: "John Doe",
		Sorts: []golister.Sort{
			{Field: "name", Order: golister.Ascending},
			{Field: "id", Order: golister.Descending},
		},
		Filters: map[string]any{"author": "asc7DsX"},
	}

	lister := golister.NewFromParams(params)

	if v := lister.Page(); v != 3 {
		t.Errorf("expected Page() to be 3, got %d", v)
	}

	if v := lister.Limit(); v != 50 {
		t.Errorf("expected Limit() to be 50, got %d", v)
	}

	if v := lister.Search(); v != "John Doe" {
		t.Errorf(`expected Search() to be "John Doe", got "%s"`, v)
	}

	if v := lister.CastFilter("author").StringSafe(""); v != "asc7DsX" {
		t.Errorf(`expected Filter() to be "asc7DsX", got "%s"`, v)
	}

	expected := ` ORDER BY "name" ASC, "id" DESC LIMIT 50 OFFSET 0`
	if sorts := lister.SQLSortOrder(); sorts != expected {
		t.Errorf(`expected SQLSortOrder() to be "%s", got "%s"`, expected, sorts)
	}
}

func TestJsonConstructor(t *testing.T) {
	raw := `{
		"page": 2,
		"limit": 100,
		"search": "Jack ma",
		"filters": {
			"author": "asc7DsX"
		},
		"sorts": [
			{ "field": "name", "order": "desc" },
			{ "field": "id", "order": "asc" }
		]
	}`

	lister, err := golister.NewFromJson(raw)
	if err != nil {
		t.Error(err.Error())
	}
	lister.SetTotal(1000)

	if v := lister.Page(); v != 2 {
		t.Errorf("expected Page() to be 2, got %d", v)
	}

	if v := lister.Limit(); v != 100 {
		t.Errorf("expected Limit() to be 100, got %d", v)
	}

	if v := lister.Search(); v != "Jack ma" {
		t.Errorf(`expected Search() to be "Jack ma", got "%s"`, v)
	}

	if v := lister.CastFilter("author").StringSafe(""); v != "asc7DsX" {
		t.Errorf(`expected Filter() to be "asc7DsX", got "%s"`, v)
	}

	expected := ` ORDER BY "name" DESC, "id" ASC LIMIT 100 OFFSET 100`
	if sorts := lister.SQLSortOrder(); sorts != expected {
		t.Errorf(`expected SQLSortOrder() to be "%s", got "%s"`, expected, sorts)
	}
}

func TestBase64Constructor(t *testing.T) {
	encoded := base64.URLEncoding.EncodeToString([]byte(
		`{
			"page": 2,
			"limit": 100,
			"search": "Jack ma",
			"filters": {
				"author": "asc7DsX"
			},
			"sorts": [
				{ "field": "name", "order": "desc" },
				{ "field": "id", "order": "asc" }
			]
		}`,
	))
	fmt.Println(encoded)

	lister, err := golister.NewFromBase64Json(encoded)
	if err != nil {
		t.Error(err.Error())
	}
	lister.SetTotal(1000)

	if v := lister.Page(); v != 2 {
		t.Errorf("expected Page() to be 2, got %d", v)
	}

	if v := lister.Limit(); v != 100 {
		t.Errorf("expected Limit() to be 100, got %d", v)
	}

	if v := lister.Search(); v != "Jack ma" {
		t.Errorf(`expected Search() to be "Jack ma", got "%s"`, v)
	}

	if v := lister.CastFilter("author").StringSafe(""); v != "asc7DsX" {
		t.Errorf(`expected Filter() to be "asc7DsX", got "%s"`, v)
	}

	expected := ` ORDER BY "name" DESC, "id" ASC LIMIT 100 OFFSET 100`
	if sorts := lister.SQLSortOrder(); sorts != expected {
		t.Errorf(`expected SQLSortOrder() to be "%s", got "%s"`, expected, sorts)
	}
}
