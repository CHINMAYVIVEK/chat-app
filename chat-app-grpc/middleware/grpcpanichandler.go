package middleware

import (
	"time"

	helper "chat-app-grpc/helper"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func toPanicError(r interface{}) error {
	return grpc.Errorf(codes.Internal, "panic: %v", r)
}

var additionalHandlers []func(interface{})

func InstallPanicHandler(handler func(interface{})) {
	additionalHandlers = append(additionalHandlers, handler)
}

func handleCrash(handler func(interface{})) {
	if r := recover(); r != nil {
		handler(r)

		if additionalHandlers != nil {
			for _, fn := range additionalHandlers {
				fn(r)
			}
		}
	}
}

func UnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	defer handleCrash(func(r interface{}) {
		err := toPanicError(r)
		helper.SugarObj.Error(err)
		helper.SugarObj.Error(err)
	})

	h, err := handler(ctx, req)

	if err != nil {
		helper.SugarObj.Error(err)
		return nil, err
	}
	//logging
	helper.SugarObj.Info("request - Method:%s\tDuration:%s\tError:%v\n", info.FullMethod, time.Now())
	helper.SugarObj.Info(h)

	return h, err

}
