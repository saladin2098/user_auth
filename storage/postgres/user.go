package postgres

import (
	"database/sql"
	"errors"
	"github.com/Mubinabd/auth_service/token"

	pb "github.com/Mubinabd/auth_service/genproto"
)

type UserStorage struct {
	db *sql.DB
}
func NewUserStorage(db *sql.DB) *UserStorage {
	return &UserStorage{
        db: db,
    }
}
func (us *UserStorage) RegisterUser(user *pb.UserCreate) (*pb.User, error) {
	query := `
		insert into users(
			id,
			username,
			password,
			email
		) values($1,$2,$3,$4)
		returning username, password, email
	`
	var registeredUser pb.User

	err := us.db.QueryRow(query, user.Id, user.Username, user.Password, user.Email).
		Scan(&registeredUser.Username, &registeredUser.Password, &registeredUser.Email)
	if err != nil {
		return nil, err
	}

	return &registeredUser, nil
}
func (us *UserStorage) GetUserInfo(id *pb.ByUsername) (*pb.User,error) {
	query := `
        select
            username,
            password,
            email
        from users
        where username = $1
    `
    var user pb.User

    err := us.db.QueryRow(query, id.Username).
        Scan(&user.Username, &user.Password, &user.Email)
    if err!= nil {
        return nil, err
    }

    return &user, nil
}
func (us *UserStorage) Loginuser(logreq *pb.LoginReq) (*pb.Token,error) {
	var usernameDB, passwordDB, user_id string
	query := `select id,username,password from users where username = $1`
	err := us.db.QueryRow(query,logreq.Username).Scan(&user_id,&usernameDB,&passwordDB)
	if err!= nil {
        return nil,err
    }
	qualify := true
	if passwordDB != logreq.Password || usernameDB != logreq.Username {
		qualify = false
	}
	if !qualify {
		return nil,errors.New("username or password incorrect")
	}
	token,err := token.GenereteJWTToken(user_id,logreq.GetUsername())
	if err!= nil {
        return nil,err
    }
	return token,nil
}
