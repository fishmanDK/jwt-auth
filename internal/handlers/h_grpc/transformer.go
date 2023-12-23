package h_grpc

import (
	jwtauth "github.com/fishmanDK"
	convert_proto "github.com/fishmanDK/internal/convert-proto/proto"
)

func ProtoNewUserToNewUser(protoNewUser *convert_proto.NewUser) jwtauth.CreateUser {
	return jwtauth.CreateUser{
		AppId:     protoNewUser.AppId,
		Role:      protoNewUser.Role,
		Email:     protoNewUser.Email,
		UserName:  protoNewUser.UserName,
		FirstName: protoNewUser.FirstName,
		LastName:  protoNewUser.LastName,
		Password:  protoNewUser.Password,
	}
}

func ProtoUserToUser(protoUser *convert_proto.User) jwtauth.User {
	return jwtauth.User{
		AppName:  protoUser.AppName,
		Email:    protoUser.Email,
		Password: protoUser.Password,
	}
}
