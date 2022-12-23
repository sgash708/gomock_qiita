package model

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	mes := "model test"
	fmt.Println("start ", mes)

	code := m.Run()

	fmt.Println("end ", mes)

	os.Exit(code)
}

func TestGetUUID(t *testing.T) {
	cases := []struct {
		UUID string
	}{
		{UUID: ""},
		{UUID: "test2"},
	}

	for _, c := range cases {
		fakeUUID = c.UUID
		uid, err := GetUUID()
		if err != nil {
			t.Log(err)
		}
		if c.UUID != "" && uid != c.UUID {
			t.Log("fakeuuidが異なります")
		}
	}
}

func TestSetFakeUUID(t *testing.T) {
	cases := []struct {
		UUID string
	}{
		{UUID: "test"},
		{UUID: "test2"},
	}

	for _, c := range cases {
		SetFakeUUID(c.UUID)
		if fakeUUID != c.UUID {
			t.Log("setしたfakeuuidが異なります")
		}
	}
}
