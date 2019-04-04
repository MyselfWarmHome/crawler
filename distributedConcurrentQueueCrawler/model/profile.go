package model

import "encoding/json"

type Profile struct {
	Name        string
	Gender      string
	Age         int
	Height      int
	Weight      int
	Income      string
	Marriage    string
	Hokou       string
	XinZuo      string
	WorkAddress string
}

func FromJsonObj(o interface{}) (Profile, error) {

	var profile Profile
	s, err := json.Marshal(o)
	if err != nil {
		return profile, err
	}
	err = json.Unmarshal(s, &profile)
	return profile, nil
}
