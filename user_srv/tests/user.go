package main

import (
	"context"
	"fmt"
	"github.com/luhaopei/mxshop_srvs/user_srv/proto"
	"google.golang.org/grpc"
)

var userClient proto.UserClient
var conn *grpc.ClientConn

func Init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	userClient = proto.NewUserClient(conn)
}
func TestGetUserList() {
	resp, err := userClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    1,
		PSize: 2,
	})
	if err != nil {
		panic(err)
	}
	for _, user := range resp.Data {
		fmt.Println(user.Mobile, user.NickName, user.Password)
		checkResp, err := userClient.CheckPassWord(context.Background(), &proto.PasswordCheckInfo{
			Password:         "admin123",
			EncryptedPasword: user.Password,
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(checkResp.Success)
	}
}

func TestCreateUser() {
	for i := 0; i < 10; i++ {
		rsp, err := userClient.CreateUser(context.Background(), &proto.CreateUserInfo{
			NickName: fmt.Sprintf("bobby%d", i),
			Mobile:   fmt.Sprintf("1878222222%d", i),
			PassWord: "admin123",
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(rsp.Id)
	}
}

func main() {
	Init()
	defer conn.Close()
	TestGetUserList()
	//TestCreateUser()
}
