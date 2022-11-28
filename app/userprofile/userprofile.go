package userprofile

import (
	"bytes"
	helper "chat-app-grpc/helper"
	"context"
	"encoding/json"
	"fmt"

	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
)

type UserProfileStruct struct {
	UnimplementedUserProfileServiceServer
}

func (s *UserProfileStruct) UserProfile(ctx context.Context, req *UserProfileRequest) (*httpbody.HttpBody, error) {

	var ret *httpbody.HttpBody

	userData, err := FetchUserData(req.UserId)
	if err != nil {
		helper.SugarObj.Error(err)
		return ret, err
	}
	helper.SugarObj.Info(userData)

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(userData)
	byteData := reqBodyBytes.Bytes()

	ret = &httpbody.HttpBody{
		ContentType: "text/json",
		Data:        byteData,
	}
	return ret, nil
}

func (s *UserProfileStruct) UserLogin(ctx context.Context, req *UserLoginRequest) (*httpbody.HttpBody, error) {

	userName := req.UserName
	emailId := req.EmailId
	password := req.Password

	fmt.Println(userName, emailId, password)
	resp := userName + emailId + password

	return &httpbody.HttpBody{
		ContentType: "application/json",
		Data:        []byte(resp),
	}, nil

}

func (s *UserProfileStruct) UserRegistration(ctx context.Context, req *UserRegistrationRequest) (*httpbody.HttpBody, error) {

	name := req.FirstName + req.MiddleName + req.LastName
	userName := req.UserName
	emailId := req.EmailId
	password := req.Password

	fmt.Println(name, userName, emailId, password)
	resp := name + userName + emailId + password

	return &httpbody.HttpBody{
		ContentType: "application/json",
		Data:        []byte(resp),
	}, nil

}
