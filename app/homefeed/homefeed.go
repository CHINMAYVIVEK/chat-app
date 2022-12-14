package homefeed

import (
	"bytes"
	helper "chat-app-grpc/helper"
	"context"
	"encoding/json"

	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	userprofile "chat-app-grpc/app/userprofile"
)

type HomePageStruct struct {
	UnimplementedHomeFeedServiceServer
}

func (s *HomePageStruct) HomeFeed(ctx context.Context, req *HomeFeedRequest) (*httpbody.HttpBody, error) {

	var ret *httpbody.HttpBody

	userData, err := userprofile.FetchUserData(req.UserId)
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
