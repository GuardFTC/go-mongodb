// Package index @Author:冯铁城 [17615007230@163.com] 2025-10-14 16:28:24
package index

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	//4.创建全文索引
	fullTextRes, err := coll.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{{"name", "text"}},
	})
	if err != nil {
		log.Printf("CreateOne FullText Index error: %v", err)
	} else {
		log.Printf("CreateOne FullText Index success: %v", fullTextRes)
	}
	listIndexes(coll, ctx)

	//5.创建Hash索引
	hashRes, err := coll.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{{"age", "hashed"}},
	})
	if err != nil {
		log.Printf("CreateOne Hash Index error: %v", err)
	} else {
		log.Printf("CreateOne Hash Index success: %v", hashRes)
	}
	listIndexes(coll, ctx)

	//6.删除全部索引
	all, err := coll.Indexes().DropAll(ctx)
	if err != nil {
		log.Printf("DropAll Index error: %v", err)
	} else {
		log.Printf("DropAll Index success: %v", all)
	}

	//6.创建唯一索引(名称唯一)
	uniqueRes, err := coll.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{"name", 1}},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		log.Printf("CreateOne Unique Index error: %v", err)
	} else {
		log.Printf("CreateOne Unique Index success: %v", uniqueRes)
	}
	listIndexes(coll, ctx)

	//7.创建稀疏索引(别名唯一且文档必须包含别名)
	sparseRes, err := coll.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{"nikeName", 1}},
		Options: options.Index().SetSparse(true).SetUnique(true),
	})
	if err != nil {
		log.Printf("CreateOne Sparse Index error: %v", err)
	} else {
		log.Printf("CreateOne Sparse Index success: %v", sparseRes)
	}
	listIndexes(coll, ctx)

	//8.创建部分索引(年龄>=18岁的文档创建该索引)
	partialRes, err := coll.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{"age", 1}},
		Options: options.Index().SetPartialFilterExpression(bson.D{{"age", bson.D{{"$gt", 18}}}}),
	})
	if err != nil {
		log.Printf("CreateOne Partial Index error: %v", err)
	} else {
		log.Printf("CreateOne Partial Index success: %v", partialRes)
	}
	listIndexes(coll, ctx)

	//9.创建TTL索引(存在登录时间的文档，超过规定时间则自动删除)
	ttlRes, err := coll.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{"loginTime", 1}},
		Options: options.Index().SetExpireAfterSeconds(10),
	})
	if err != nil {
		log.Printf("CreateOne TTL Index error: %v", err)
	} else {
		log.Printf("CreateOne TTL Index success: %v", ttlRes)
	}
	listIndexes(coll, ctx)
}

// listIndexes 遍历全部索引
func listIndexes(coll *mongo.Collection, ctx context.Context) {

	//1.获取集合列表
	indexView, err := coll.Indexes().List(ctx)
	if err != nil {
		log.Printf("List Index error: %v", err)
		return
	}

	//2.解析出数据
	var many bson.A
	for indexView.Next(ctx) {
		var index bson.M
		if err := indexView.Decode(&index); err != nil {
			log.Printf("Decode Index error: %v", err)
		} else {
			many = append(many, index)
		}
	}

	//3.遍历many
	for _, index := range many {
		log.Printf("Index: %v", index)
	}
	log.Println("----------------------------------------------------------------------------")
}
