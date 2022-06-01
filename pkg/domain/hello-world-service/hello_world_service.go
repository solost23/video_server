package hello_world_service

import (
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/resolver/manual"

	HelloWorldService "github.com/solost23/my_interface/hello_world_service"
)

var client *grpc.ClientConn

func GetHelloWorldServiceClient() (HelloWorldService.HelloWorldServiceClient, error) {
	grpcClient, err := getClient()
	if err != nil {
		return nil, err
	}
	return HelloWorldService.NewHelloWorldServiceClient(grpcClient), nil
}

func getClient() (*grpc.ClientConn, error) {
	var err error
	if client == nil {
		client, err = newClient()
	}
	return client, err
}

func newClient() (conn *grpc.ClientConn, err error) {
	addrSlice := viper.GetStringSlice("gRpc.hello_world_service.addrSlice")
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
