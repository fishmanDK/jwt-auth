package h_grpc

import (
	"context"
	jwtauth "github.com/fishmanDK"
	convert_proto "github.com/fishmanDK/internal/convert-proto/proto"
)

func (s *ServerApi) CreateUer(ctx context.Context, protoNewUser *convert_proto.NewUser) (*convert_proto.UserId, error) {
	var inp jwtauth.CreateUser
	inp = ProtoNewUserToNewUser(protoNewUser)

	userId, err := s.Service.Authentication.CreateUser(inp)
	if err != nil {
		return &convert_proto.UserId{
			Response: &convert_proto.Response{
				Status: "Error",
				Error:  "Error create new user, please try again later",
			},
		}, err
	}

	return &convert_proto.UserId{
		Response: &convert_proto.Response{
			Status: "OK",
		},
		UserId: userId,
	}, nil
}

func (s *ServerApi) Authentication(ctx context.Context, protoUser *convert_proto.User) (*convert_proto.AuthResponse, error) {
	var inp jwtauth.User
	inp = ProtoUserToUser(protoUser)

	tokens, err := s.Service.Authentication.Authentication(inp)
	if err != nil {
		return &convert_proto.AuthResponse{
			Response: &convert_proto.Response{
				Status: "Error",
				Error:  "Error auth, , please try again later",
			},
		}, err
	}

	return &convert_proto.AuthResponse{
		Tokens: &convert_proto.Tokens{
			AccessToken:  tokens.Access_token,
			RefreshToken: tokens.Refresh_token,
		},
	}, nil
}
