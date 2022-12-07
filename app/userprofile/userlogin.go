package userprofile

import (
	config "chat-app-grpc/config"
	helper "chat-app-grpc/helper"

	_ "github.com/lib/pq"
)

func GetUserLogin(userName, password string) bool {

	query := "SELECT user_name,password  FROM public.user WHERE user_name = '" + userName + "' OR email_id = '" + userName + "'"

	helper.SugarObj.Info(query)

	rows, err := config.GetDb().Query(query)
	if err != nil {
		helper.SugarObj.Error(err)
		return false
	}
	defer rows.Close()

	for rows.Next() {
		var userName string
		var passString string

		err = rows.Scan(&userName, &passString)

		if err != nil {
			helper.SugarObj.Error(err)
			return false
		}

		if password == passString {
			return true
		}

	}

	return false
}
