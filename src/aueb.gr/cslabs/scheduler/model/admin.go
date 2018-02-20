package model

import "encoding/json"

type Admin struct {
	Name        string
	Preferences map[string]Preference
	Distance    int
}

func (admin Admin) String() string {
	return admin.Name
}

func (admin Admin) StringDetailed() string {
	jsonString, err := json.Marshal(admin.Preferences)
	if err != nil {
		panic(err)
	}
	return admin.Name + "\n" + string(jsonString)
}
