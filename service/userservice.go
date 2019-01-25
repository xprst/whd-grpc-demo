package service

import (
	"context"
	"github.com/go-xorm/xorm"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/xprst/whd-grpc-base/app"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
	pb "whd-grpc-demo/proto"
	"whd-grpc-demo/storge/entity"
)

// UserService 用户服务
type UserService struct {
	orm *xorm.Engine
}

// NewUserService 创建一个用户服务的实例
func NewUserService() *UserService {
	return &UserService{
		// 从app.MyBean中获取数据库对应的xorm引擎
		orm: app.MyBean.GetEngine("default"),
	}
}

// ListUsers 获取用户列表
func (t UserService) ListUsers(ctx context.Context, r *pb.ListUserRequest) (*pb.ListUserResponse, error) {
	users := make([]entity.User, 0)
	err := t.orm.Find(&users)

	if err != nil {
		grpclog.Errorf("ListUsers orm find err:%v, ListUserRequest: %s", err, r.String())

		return nil, err
	}

	pbUsers := make([]*pb.User, 0)

	for i := 0; i < len(users); i++ {
		pbUsers = append(pbUsers, &pb.User{
			Id:   users[i].ID,
			Name: users[i].Name,
			Age:  users[i].Age,
			Sex:  users[i].Sex,
		})
	}

	return &pb.ListUserResponse{
		Users: pbUsers,
	}, err
}

// GetUser 获取用户
func (t UserService) GetUser(ctx context.Context, r *pb.GetUserRequest) (*pb.User, error) {
	user := new(entity.User)
	has, err := t.orm.Id(r.Id).Get(user)

	if err != nil {
		grpclog.Errorf("GetUser orm get err:%v, Id: %d", err, r.Id)

		return nil, err
	}

	if !has {
		err = status.Errorf(codes.NotFound, "cannot find user with id %d", r.Id)
		return nil, err
	}

	var pbUser = &pb.User{
		Id:   user.ID,
		Name: user.Name,
		Age:  user.Age,
		Sex:  user.Sex,
	}

	return pbUser, err
}

// CreateUser 创建用户
func (t UserService) CreateUser(ctx context.Context, r *pb.CreateUserRequest) (*pb.User, error) {
	user := new(entity.User)
	user.Name = r.User.Name
	user.Age = r.User.Age
	user.Sex = r.User.Sex
	_, err := t.orm.Insert(user)

	if err != nil {
		grpclog.Errorf("CreateUser orm insert err:%v, CreateUserRequest: %s", err, r.String())

		return nil, err
	}

	return &pb.User{
		Id:     user.ID,
		Name:   user.Name,
		Age:   user.Age,
		Sex:   user.Sex,
	}, err
}

// UpdateUser 更新用户
func (t UserService) UpdateUser(ctx context.Context, r *pb.UpdateUserRequest) (*pb.User, error) {
	user := new(entity.User)
	user.ID = r.User.Id
	user.Name = r.User.Name
	user.Age = r.User.Age
	user.Sex = r.User.Sex
	_, err := t.orm.Id(user.ID).Update(user)

	if err != nil {
		grpclog.Errorf("UpdateUser orm update err:%v, UpdateUserRequest: %s", err, r.String())

		return nil, err
	}

	return &pb.User{
		Id:     user.ID,
		Name:   user.Name,
		Age:   user.Age,
		Sex:   user.Sex,
	}, err
}

// DeleteUser 删除用户
func (t UserService) DeleteUser(ctx context.Context, r *pb.DeleteUserRequest) (*empty.Empty, error) {
	var user entity.User
	_, err := t.orm.Id(r.Id).Delete(&user)

	if err != nil {
		grpclog.Errorf("DeleteUser orm get err:%v, Id: %d", err, r.Id)

		return nil, err
	}

	return new(empty.Empty), err
}
