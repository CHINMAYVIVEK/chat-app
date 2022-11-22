package homefeed

import (
	"context"

	_ "github.com/go-sql-driver/mysql"
)

type HomePageStruct struct {
	UnimplementedHomeFeedServiceServer
}

func (s *HomePageStruct) HomeFeed(ctx context.Context, req *HomeFeedRequest) (*HomeFeedResponse, error) {

	userDetailResp := &UserDetail{
		UserName:   "Chinmay",
		UserId:     "9876543210",
		IsVerified: 0,
		ProfilePic: "pic",
	}
	ret := &HomeFeedResponse{
		UserDetail: userDetailResp,
		Tweet:      "Test tweet",
	}

	return ret, nil
}
