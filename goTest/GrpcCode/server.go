package main

import (
	"google.golang.org/grpc"
	"GOlang/goTest/GrpcCode/message"
	"context"
	"errors"
	"net"
	"time"
)

//订单服务
type OrderServiceImpl struct {
}

func (os *OrderServiceImpl) GetOrderInfo(ctx context.Context, request *message.OrderRequest) (*message.OrderInfo, error) {
	orderMap := map[string]message.OrderInfo{
		"201907300001": {OrderId: "201907300001", OrderName: "衣服", OrderStatus: "已付款"},
		"201907300002": {OrderId: "201907300002", OrderName: "零食", OrderStatus: "已付款"},
		"201907300003": {OrderId: "201907300003", OrderName: "食品", OrderStatus: "未付款"},
	}
	var response *message.OrderInfo
	current := time.Now().Unix()
	if request.TimeStamp > current {
		*response = message.OrderInfo{OrderId: "0", OrderName: "", OrderStatus: "订单信息异常"}
	} else {
		result := orderMap[request.OrderId]
		if result.OrderId != "" {
			*response = orderMap[request.OrderId]
			return &result, nil
		} else {
			return nil, errors.New("server error")
		}
	}
	return response, nil
}

func main() {
	server := grpc.NewServer()
	message.RegisterOrderServiceServer(server, new(OrderServiceImpl))
	lis, err := net.Listen("tcp", ":8090")
	if err != nil {
		panic(err.Error())
	}
	server.Serve(lis)
}
