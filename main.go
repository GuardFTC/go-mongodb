// @Author:冯铁城 [17615007230@163.com] 2025-10-11 10:48:20
package main

import (
	"mongodb-demo/client"
	"mongodb-demo/db_and_collection"
	"mongodb-demo/document"

	"go.mongodb.org/mongo-driver/mongo"
)

func main() {

	//1.创建客户端
	mongoClient := client.CreateMongoClient()
	defer client.CloseMongoClient(mongoClient)

	//2.创建数据库
	db := db_and_collection.CreateDb(mongoClient.GetClient(), mongoClient.GetCtx())

	//3.创建集合
	db_and_collection.CreateCollection(db, mongoClient.GetCtx())

	//4.获取集合
	collection := db.Collection("testCollection")

	//5.文档操作
	testDocument(collection, mongoClient)

	//6.删除集合
	db_and_collection.DropCollection(db, mongoClient.GetCtx())

	//7.删除数据库
	db_and_collection.DropDb(mongoClient.GetClient(), mongoClient.GetCtx())
}

// testDocument 测试文档操作
func testDocument(collection *mongo.Collection, mongoClient *client.MongoClient) {

	//1.插入文档
	document.Insert(collection, mongoClient.GetCtx())
}
