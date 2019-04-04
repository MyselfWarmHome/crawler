package rpcsupport

import (
	"crawler/distributedConcurrentQueueCrawler/config"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

/**
JsonRpc的服务端
*/
func ServeRpc(host string, service interface{}) error {
	rpc.Register(service)
	listener, err := net.Listen(config.NetProtocol, host)
	if err != nil {
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("接受数据出错： %v", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
	return nil
}

/**
JsonRpc的客户端
*/
func NewClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial(config.NetProtocol, host)
	if err != nil {
		return nil, err
	}
	return jsonrpc.NewClient(conn), nil
}
