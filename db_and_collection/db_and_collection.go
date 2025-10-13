// Package db_and_collection @Author:冯铁城 [17615007230@163.com] 2025-10-13 14:46:41
package db_and_collection

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateAndDropDb 创建数据库并删除数据库
func CreateAndDropDb(client *mongo.Client, ctx context.Context) {

	//1.创建数据库
	db := client.Database("testDb")
	log.Printf("创建数据库成功：%s\n", db.Name())

	//2.删除数据库
	err := client.Database("testDb").Drop(ctx)
	if err != nil {
		log.Printf("删除数据库失败：%v\n", err)
	} else {
		log.Println("删除数据库成功")
	}
}

// CreateAndDropCollection 创建集合并删除集合
func CreateAndDropCollection(client *mongo.Client, ctx context.Context) {

	//1.创建数据库
	db := client.Database("testDb")

	//2.创建集合
	err := db.CreateCollection(ctx, "testCollection")
	if err != nil {
		log.Printf("创建集合失败：%v\n", err)
	} else {
		log.Println("创建集合成功")
	}

	//3.创建一个固定集合
	err = db.CreateCollection(ctx, "testCappedCollection", options.CreateCollection().
		SetCapped(true).
		SetSizeInBytes(1024).
		SetMaxDocuments(100))
	if err != nil {
		log.Printf("创建固定集合失败：%v\n", err)
	} else {
		log.Println("创建固定集合成功")
	}

	//4.列出所有集合
	collections, err := db.ListCollectionNames(ctx, bson.D{})
	if err == nil {
		log.Println("所有集合：")
		for _, collection := range collections {
			log.Println(collection)
		}
	}

	//5.删除集合
	err = db.Collection("testCollection").Drop(ctx)
	if err != nil {
		log.Printf("删除集合失败：%v\n", err)
	} else {
		log.Println("删除集合成功")
	}

	//6.删除固定集合
	err = db.Collection("testCappedCollection").Drop(ctx)
	if err != nil {
		log.Printf("删除固定集合失败：%v\n", err)
	} else {
		log.Println("删除固定集合成功")
	}
}
