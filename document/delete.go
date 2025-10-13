// Package document @Author:冯铁城 [17615007230@163.com] 2025-10-13 17:37:58
package document

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Delete 删除数据
func Delete(coll *mongo.Collection, ctx context.Context) {

	//1.删除名字叫做zyl的文档
	deleteOneResult, err := coll.DeleteOne(ctx, bson.M{"name": "zyl"})
	if err != nil {
		log.Printf("删除数据失败：%v\n", err)
	} else {
		log.Printf("删除数据成功，删除的文档数量为：%v\n", deleteOneResult.DeletedCount)
	}

	//2.删除年纪为18的文档
	deleteManyResult, err := coll.DeleteMany(ctx, bson.M{"age": 18})
	if err != nil {
		log.Printf("删除数据失败：%v\n", err)
	} else {
		log.Printf("删除数据成功，删除的文档数量为：%v\n", deleteManyResult.DeletedCount)
	}

	//3.全量删除
	deleteResult, err := coll.DeleteMany(ctx, bson.D{{}})
	if err != nil {
		log.Printf("删除数据失败：%v\n", err)
	} else {
		log.Printf("删除数据成功，删除的文档数量为：%v\n", deleteResult.DeletedCount)
	}
}
