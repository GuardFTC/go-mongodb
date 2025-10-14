// Package index @Author:冯铁城 [17615007230@163.com] 2025-10-14 16:28:24
package index

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Index 索引
func Index(coll *mongo.Collection, ctx context.Context) {

	//1.获取全部索引
	listIndexes(coll, ctx)

	//2.创建单列索引
	oneRes, err := coll.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{{"name", 1}},
	})
	if err != nil {
		log.Printf("CreateOne Index error: %v", err)
	} else {
		log.Printf("CreateOne Index success: %v", oneRes)
	}
	listIndexes(coll, ctx)

	//3.创建多列索引
	manyRes, err := coll.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{
			Keys: bson.D{{"name", 1}, {"age", -1}},
		},
	})
	if err != nil {
		log.Printf("CreateMany Index error: %v", err)
	} else {
		log.Printf("CreateMany Index success: %v", manyRes)
	}
	listIndexes(coll, ctx)
}

// listIndexes 遍历全部索引
func listIndexes(coll *mongo.Collection, ctx context.Context) {
	indexView, err := coll.Indexes().List(ctx)
	if err != nil {
		log.Printf("List Index error: %v", err)
	} else {
		var many bson.A
		for indexView.Next(ctx) {
			var index bson.M
			if err := indexView.Decode(&index); err != nil {
				log.Printf("Decode Index error: %v", err)
			} else {
				many = append(many, index)
			}
		}
		log.Printf("List Index success: %v", many)
	}
}
