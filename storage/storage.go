package storage

import pb "github.com/Mubinabd/auth_service/genproto"

type StorageI interface {
	User() UserI
}

type UserI interface {
	RegisterUser(user *pb.UserCreate) (*pb.User, error)
	GetUserInfo(id *pb.ByUsername) (*pb.User,error)
	Loginuser(logreq *pb.LoginReq) (*pb.Token,error)
}