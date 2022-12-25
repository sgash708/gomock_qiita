package model

import "testing"

func TestSetUUID(t *testing.T) {
	cases := []struct {
		Name string
	}{
		{Name: "test1"},
		{Name: "test2"},
	}

	for _, c := range cases {
		book := &Book{
			Name: c.Name,
		}

		if err := book.SetUUID(); err != nil {
			t.Error(err)
		}
	}
}

func TestUpdateBookName(t *testing.T) {
	name := "test"
	updatedName := "updated-test"

	book := &Book{
		Name: name,
		UUID: "test-uuid",
		ID:   1,
	}

	book.UpdateBookName(updatedName)

	if book.Name == name {
		t.Error("updateに失敗しています")
	}
}
