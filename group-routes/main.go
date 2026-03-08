package main

import "github.com/gin-examples/group-routes/routes"

func main() {
	routes.Run()
}


// # 访问 v1 分组的 ping 路由
// curl http://localhost:8080/v1/ping
// # 响应："pong"

// # 访问 v1 分组的 users 路由
// curl http://localhost:8080/v1/users
// # 响应："users"

// # 访问 v2 分组的 ping 路由
// curl http://localhost:8080/v2/ping
// # 响应："pong"