package model

import "github.com/google/uuid"

var fakeUUID string

func GetUUID() (string, error) {
	if fakeUUID == "" {
		uid, err := uuid.NewRandom()
		if err != nil {
			return "", err
		}
		return uid.String(), nil
	}

	return fakeUUID, nil
}

func SetFakeUUID(uid string) {
	fakeUUID = uid
}
