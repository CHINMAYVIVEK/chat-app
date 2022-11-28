package userprofile

import (
	config "chat-app-grpc/config"
	helper "chat-app-grpc/helper"

	_ "github.com/lib/pq"
)

type UserData struct {
	UserName   string `json:"username"`
	FirstName  string `json:"firstname"`
	MiddleName string `json:"middlename"`
	LastName   string `json:"lastname"`
	IsVerified string `json:"isverified"`
	IsOfficial string `json:"isofficial"`
	ProfilePic string `json:"profilepic"`
}

func FetchUserData(userId string) (*UserData, error) {

	userData := &UserData{}
	query := "SELECT user_name, first_name,middle_name,last_name,is_verified,is_official,profile_pic FROM public.user WHERE id = '" + userId + "'"

	helper.SugarObj.Info(query)

	rows, err := config.GetDb().Query(query)
	if err != nil {
		helper.SugarObj.Error(err)
		return userData, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&userData.UserName, &userData.FirstName, &userData.MiddleName, &userData.LastName, &userData.IsVerified, &userData.IsOfficial, &userData.ProfilePic)

		if err != nil {
			helper.SugarObj.Error(err)
			return userData, err
		}

	}

	helper.SugarObj.Info(userData)
	return userData, nil

}
