// Package document @Author:冯铁城 [17615007230@163.com] 2025-10-13 15:24:41
package document

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Insert 插入数据
func Insert(coll *mongo.Collection, ctx context.Context) {

	//1.创建文档
	doc1 := bson.D{{"name", "ftc"}, {"age", 18}}
	doc2 := bson.D{{"name", "zyl"}, {"age", 15}}
	doc3 := bson.D{{"name", "skx"}, {"age", 18}}

	//2.insertOne
	insertOneResult, err := coll.InsertOne(ctx, doc1)
	if err != nil {
		log.Printf("insertOne error: %v", err)
	} else {
		log.Printf("insertOne success: %v", insertOneResult)
	}

	//3.insertMany
	insertManyResult, err := coll.InsertMany(ctx, []interface{}{doc2, doc3})
	if err != nil {
		log.Printf("insertMany error: %v", err)
	} else {
		log.Printf("insertMany success: %v", insertManyResult)
	}

	//4.insert重复数据，报错
	doc1 = append(doc1, bson.E{Key: "_id", Value: insertOneResult.InsertedID})
	_, err = coll.InsertOne(ctx, doc1)
	if err != nil {
		log.Printf("repeat insertOne error: %v", err)
	} else {
		log.Println("repeat insertOne success")
	}
}
