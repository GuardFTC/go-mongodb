// @Author:冯铁城 [17615007230@163.com] 2025-10-11 10:48:20
package main

import (
	"mongodb-demo/client"
	"mongodb-demo/db_and_collection"
)

func main() {

	//1.创建客户端
	mongoClient := client.CreateMongoClient()
	defer client.CloseMongoClient(mongoClient)

	//2.测试数据库以及集合操作

	//1.创建/删除数据库
	db_and_collection.CreateAndDropDb(mongoClient.GetClient(), mongoClient.GetCtx())

	//2.创建/删除集合
	db_and_collection.CreateAndDropCollection(mongoClient.GetClient(), mongoClient.GetCtx())
}
