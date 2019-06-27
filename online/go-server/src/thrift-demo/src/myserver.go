package main

import (
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"love"
)

func main() {
	addr := "localhost:9091"
	transport, err := thrift.NewTServerSocket(addr)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	if err != nil {
		fmt.Println("transport error")
	}
	handler := NewWoquHandler()
	processor := love.NewWoquProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
	server.Serve()
}
