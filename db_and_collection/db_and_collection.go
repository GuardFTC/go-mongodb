// Package db_and_collection @Author:冯铁城 [17615007230@163.com] 2025-10-13 14:46:41
package db_and_collection

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateDb 创建数据库
func CreateDb(client *mongo.Client, ctx context.Context) *mongo.Database {
	db := client.Database("testDb")
	log.Printf("创建数据库成功：%s\n", db.Name())
	return db
}

// DropDb 删除数据库
func DropDb(client *mongo.Client, ctx context.Context) {
	err := client.Database("testDb").Drop(ctx)
	if err != nil {
		log.Printf("删除数据库失败：%v\n", err)
	} else {
		log.Println("删除数据库成功")
	}
}

// CreateCollection 创建集合
func CreateCollection(db *mongo.Database, ctx context.Context) {

	//1.创建普通集合
	err := db.CreateCollection(ctx, "testCollection")
	if err != nil {
		log.Printf("创建集合失败：%v\n", err)
	} else {
		log.Println("创建普通集合成功")
	}

	//2.创建一个固定大小的集合
	err = db.CreateCollection(ctx, "testCappedCollection", options.CreateCollection().
		SetCapped(true).
		SetSizeInBytes(1024).
		SetMaxDocuments(100))
	if err != nil {
		log.Printf("创建固定集合失败：%v\n", err)
	} else {
		log.Println("创建固定集合成功")
	}

	//3.列出所有集合
	collections, err := db.ListCollectionNames(ctx, bson.D{})
	if err == nil {
		log.Println("所有集合：")
		for _, collection := range collections {
			log.Println(collection)
		}
	}
}

// DropCollection 删除集合
func DropCollection(db *mongo.Database, ctx context.Context) {

	//1.删除普通集合
	err := db.Collection("testCollection").Drop(ctx)
	if err != nil {
		log.Printf("删除普通集合失败：%v\n", err)
	} else {
		log.Println("删除普通集合成功")
	}

	//2.删除固定大小集合
	err = db.Collection("testCappedCollection").Drop(ctx)
	if err != nil {
		log.Printf("删除固定集合失败：%v\n", err)
	} else {
		log.Println("删除固定集合成功")
	}
}
