package services

import (
	"context"
	"errors"
	"grpc_v1/common"
	"grpc_v1/pbfiles"
	"log"
	"net/mail"
)

type UserService struct {

}

func (*UserService) LoginUser(ctx context.Context, in *pbfiles.User) (*pbfiles.UserResponse, error)  {
	validationErr := validate(in)
	if validationErr != nil{
		return &pbfiles.UserResponse{
			Code: 500,
			Message: validationErr.Error(),
		}, validationErr
	}
	// GORM操作
	db := common.GetDB()
	// 获取参数
	email := in.Email
	password := in.Password
	var user pbfiles.User
	db.Where("email = ?", email).First(&user)
	if user.Id == "" {
		return &pbfiles.UserResponse{
			Code: 422,
			Message: "用户不存在",
		}, nil
	}
	if user.Password != password {
		return &pbfiles.UserResponse{
			Code: 400,
			Message: "密码错误",
		}, nil
	}
	// 登陆成功
	//fmt.Println("登陆成功，用户为：",user.Email)
	return &pbfiles.UserResponse{
		Code: 200,
		Message: "OK",
		User: &user,
	}, nil
}

func (*UserService) GetUserStream(in *pbfiles.UsersRequest, stream pbfiles.UserService_GetUserStreamServer) error {
	// GORM操作
	db := common.GetDB()
	users := make([] *pbfiles.UserResponse, 0)
	for index, user := range in.Users{
		validationErr := validate(user)
		// 输入格式验证失败
		if validationErr != nil{
			users = append(users, &pbfiles.UserResponse{
				Code: 500,
				Message: validationErr.Error(),
			})
		}
		// 查询数据库
		id := user.Id
		var userRes pbfiles.User
		db.First(&userRes, id) // 查询id为1的user
		users = append(users, &pbfiles.UserResponse{
			Code: 200,
			Message: "OK",
			User: &userRes,
		})

		// 使用stream流返回结果
		if (index+1) % 2 == 0 && index > 0 {
			err := stream.Send(&pbfiles.UsersStreamResponse{UserRep: users})
			if err != nil {
				return err
			}
		}
		// 每次发送完成，清除users的缓存
		users = users[0:0]
	}
	// 发送最后一批数据
	if len(users) > 0{
		err := stream.Send(&pbfiles.UsersStreamResponse{UserRep: users})
		if err != nil {
			return err
		}
	}
	return nil
}

func (*UserService) LoginUserSql(ctx context.Context, in *pbfiles.User) (*pbfiles.UserResponse, error)  {
	validationErr := validate(in)
	if validationErr != nil{
		return &pbfiles.UserResponse{
			Code: 500,
			Message: validationErr.Error(),
		}, validationErr
	}
	// MYSQL操作
	db := common.GetMysql()
	// 获取参数
	email := in.Email
	password := in.Password
	var user pbfiles.User
	//db.First(&user, id) // 查询id为1的user
	err := db.QueryRow("SELECT * FROM users WHERE email = ?", email).Scan(&user.Id, &user.Email,&user.Name ,&user.Password)
	if err != nil {
		log.Fatal("查询发生异常！")
	}
	if user.Id == "" {
		return &pbfiles.UserResponse{
			Code: 422,
			Message: "用户不存在",
		}, nil
	}
	if user.Password != password {
		return &pbfiles.UserResponse{
			Code: 400,
			Message: "密码错误",
		}, nil
	}
	// 登陆成功
	//log.Println("登陆成功，用户为：",user.Email)
	return &pbfiles.UserResponse{
		Code: 200,
		Message: "OK",
		User: &user,
	}, nil
}

func (*UserService) GetUserStreamSql(in *pbfiles.UsersRequest, stream pbfiles.UserService_GetUserStreamSqlServer) error {
	// GORM操作
	db := common.GetMysql()
	users := make([] *pbfiles.UserResponse, 0)
	for index, user := range in.Users{
		validationErr := validate(user)
		// 输入格式验证失败
		if validationErr != nil{
			users = append(users, &pbfiles.UserResponse{
				Code: 500,
				Message: validationErr.Error(),
			})
		}
		// 查询数据库
		//id := user.Id
		email := user.Email
		password := user.Password
		var userRes pbfiles.User
		//db.First(&userRes, id) // 查询id为1的user
		err := db.QueryRow("SELECT * FROM users WHERE email = ?", email).Scan(&userRes.Id, &userRes.Email,&userRes.Name ,&userRes.Password)
		if err != nil {
			log.Fatal("查询发生异常！")
		}
		// 数据校验
		if userRes.Id == "" {
			users = append(users, &pbfiles.UserResponse{
				Code: 422,
				Message: "用户不存在",
			})
		}else if userRes.Password != password{
			users = append(users, &pbfiles.UserResponse{
				Code: 400,
				Message: "密码错误",
			})
		}else {
			users = append(users, &pbfiles.UserResponse{
				Code: 200,
				Message: "OK",
				User: &userRes,
			})
		}
		// 使用stream流返回结果
		if (index+1) % 2 == 0 && index > 0 {
			err := stream.Send(&pbfiles.UsersStreamResponse{UserRep: users})
			if err != nil {
				return err
			}
			// 每次发送完成，清除users的缓存
			users = users[0:0]
		}

	}
	// 发送最后一批数据
	if len(users) > 0{
		err := stream.Send(&pbfiles.UsersStreamResponse{UserRep: users})
		if err != nil {
			return err
		}
	}
	return nil
}

func validate(in *pbfiles.User) error {
	_, err := mail.ParseAddress(in.Email)
	if err != nil {
		return err
	}

	if len(in.Name) < 4 {
		return errors.New("Name is too short")
	}

	if len(in.Password) < 4 {
		return errors.New("Password is too weak")
	}

	return nil
}

