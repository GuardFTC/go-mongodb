// @Author:冯铁城 [17615007230@163.com] 2025-10-11 10:48:20
package main

import (
	"mongodb-demo/client"
)

func main() {

	//1.创建客户端
	mongoClient := client.CreateMongoClient()
	defer client.CloseMongoClient(mongoClient)
}
