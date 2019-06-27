package main

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"love"
)

func main() {
	addr := "localhost:9091"
	var transport thrift.TTransport
	var defaultCtx = context.Background()
	transport, err := thrift.NewTSocket(addr)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	if err != nil {
		fmt.Println("巧弄殷勤不知惫，强作 error")
	}
	transport, err = transportFactory.GetTransport(transport)
	if err != nil {
		fmt.Println("get transport error")
	}
	defer transport.Close()
	if err := transport.Open(); err != nil {
		fmt.Println("open error")
	}
	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)
	client := love.NewWoquClient(thrift.NewTStandardClient(iprot, oprot))
	client.GetId(defaultCtx)
}
