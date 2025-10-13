// Package document @Author:冯铁城 [17615007230@163.com] 2025-10-13 17:57:45
package document

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Update 更新数据
func Update(coll *mongo.Collection, ctx context.Context) {

	//1.更新单个数据
	updateOneResult, err := coll.UpdateOne(ctx, bson.M{"name": "ftc"}, bson.M{"$set": bson.M{"name": "ftcftcftc"}})
	if err != nil {
		log.Printf("更新数据失败：%v\n", err)
	} else {
		log.Printf("更新数据成功，更新的文档数量为：%v\n", updateOneResult.ModifiedCount)
	}

	//2.更新多个数据
	updateManyResult, err := coll.UpdateMany(ctx, bson.M{"age": 18}, bson.M{"$set": bson.M{"age": 19}})
	if err != nil {
		log.Printf("更新数据失败：%v\n", err)
	} else {
		log.Printf("更新数据成功，更新的文档数量为：%v\n", updateManyResult.ModifiedCount)
	}

	//3.增量更新
	updateManyResult, err = coll.UpdateMany(ctx, bson.M{"age": 19}, bson.M{"$inc": bson.M{"age": -1}})
	if err != nil {
		log.Printf("更新数据失败：%v\n", err)
	} else {
		log.Printf("更新数据成功，更新的文档数量为：%v\n", updateManyResult.ModifiedCount)
	}

	//4.按照ID更新
	doc := bson.D{{"name", "wkf"}, {"age", 19}}
	oneResult, err := coll.InsertOne(ctx, doc)
	if err == nil {
		updateByIDResult, err := coll.UpdateByID(ctx, oneResult.InsertedID, bson.M{"$set": bson.M{"name": "wkfwkfwkf"}})
		if err != nil {
			log.Printf("更新数据失败：%v\n", err)
		} else {
			log.Printf("更新数据成功，更新的文档数量为：%v\n", updateByIDResult.ModifiedCount)
		}
	}
}
