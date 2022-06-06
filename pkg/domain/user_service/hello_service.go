package user_service

import (
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/resolver/manual"

	UserService "github.com/solost23/go_interface/gen-go/user_service"
)

var client *grpc.ClientConn

func GetUserServiceClient() (UserService.UserServiceClient, error) {
	grpcClient, err := getClient()
	if err != nil {
		return nil, err
	}
	return UserService.NewUserServiceClient(grpcClient), nil
}

func getClient() (*grpc.ClientConn, error) {
	var err error
	if client == nil {
		client, err = newClient()
	}
	return client, err
}

func newClient() (conn *grpc.ClientConn, err error) {
	addrSlice := viper.GetStringSlice("gRpc.user_service.addrSlice")
	var addrs []resolver.Address
	r := manual.NewBuilderWithScheme("whatever")
	for _, addr := range addrSlice {
		addrs = append(addrs, resolver.Address{Addr: addr})
	}
	conn, err = grpc.Dial(
		r.Scheme()+":///test.server",
		grpc.WithInsecure(),
		grpc.WithResolvers(r),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`))
	r.UpdateState(resolver.State{Addresses: addrs})
	return
}
