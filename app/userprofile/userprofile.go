package userprofile

import (
	"bytes"
	helper "chat-app-grpc/helper"
	"context"
	"encoding/json"

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
