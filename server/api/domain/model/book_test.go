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
