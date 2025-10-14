// Package document @Author:冯铁城 [17615007230@163.com] 2025-10-14 10:28:44
package document

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SelectOneAndMany 查询数据一个或多个
func SelectOneAndMany(coll *mongo.Collection, ctx context.Context) {

	//1.查询一个数据
	var one bson.M
	err := coll.FindOne(ctx, bson.D{{"name", "ftc"}}).Decode(&one)
	if err != nil {
		log.Printf("FindOne error: %v", err)
	} else {
		log.Printf("FindOne success: %v", one)
	}

	//2.查询多个数据
	var many bson.A
	cur, err := coll.Find(ctx, bson.D{{"age", 18}})
	parseFindManyResult(err, cur, ctx, many, "FindMany")
}

// SelectByCondition 条件查询
func SelectByCondition(coll *mongo.Collection, ctx context.Context) {

	//1.等值查询
	var many bson.A
	cur, err := coll.Find(ctx, bson.D{{"age", 18}})
	parseFindManyResult(err, cur, ctx, many, "Equal Find")

	//2.小于或小于等于查询
	cur, err = coll.Find(ctx, bson.D{{"age", bson.D{{"$lt", 18}}}})
	parseFindManyResult(err, cur, ctx, many, "Less Than Find")
	cur, err = coll.Find(ctx, bson.D{{"age", bson.D{{"$lte", 18}}}})
	parseFindManyResult(err, cur, ctx, many, "Less Than or Equal Find")

	//3.大于或大于等于查询
	cur, err = coll.Find(ctx, bson.D{{"age", bson.D{{"$gt", 18}}}})
	parseFindManyResult(err, cur, ctx, many, "Greater Than Find")
	cur, err = coll.Find(ctx, bson.D{{"age", bson.D{{"$gte", 18}}}})
	parseFindManyResult(err, cur, ctx, many, "Greater Than or Equal Find")

	//4.不等值查询
	cur, err = coll.Find(ctx, bson.D{{"age", bson.D{{"$ne", 18}}}})
	parseFindManyResult(err, cur, ctx, many, "Not Equal Find")

	//5.In查询
	cur, err = coll.Find(ctx, bson.D{{"age", bson.D{{"$in", []int{15, 19}}}}})
	parseFindManyResult(err, cur, ctx, many, "In Find")

	//6.Not In查询
	cur, err = coll.Find(ctx, bson.D{{"age", bson.D{{"$nin", []int{15, 19}}}}})
	parseFindManyResult(err, cur, ctx, many, "Not In Find")

	//7.And查询
	cur, err = coll.Find(ctx, bson.D{
		{"age", bson.D{{"$nin", []int{15, 19}}}},
		{"name", "wqw"},
	})
	parseFindManyResult(err, cur, ctx, many, "And Find")

	//8.Or查询
	cur, err = coll.Find(ctx, bson.D{{"$or", bson.A{
		bson.D{{"age", 15}},
		bson.D{{"name", "wqw"}},
	}}})
	parseFindManyResult(err, cur, ctx, many, "Or Find")

	//9.And与Or查询
	cur, err = coll.Find(ctx, bson.D{
		{"age", bson.D{{"$nin", []int{15, 19}}}},
		{"$or", bson.A{
			bson.D{{"name", "skx"}},
			bson.D{{"name", "wqw"}},
		}},
	})
	parseFindManyResult(err, cur, ctx, many, "And With Or Find")

	//10.查询部分字段
	cur, err = coll.Find(
		ctx,
		bson.D{{"age", 18}},
		options.Find().SetProjection(bson.D{
			{"name", 1},
			{"age", 1},
			{"_id", 0},
		}),
	)
	parseFindManyResult(err, cur, ctx, many, "Partial Field Find")
}

// parseFindManyResult 解析批量查询结果
func parseFindManyResult(err error, cur *mongo.Cursor, ctx context.Context, many bson.A, title string) {
	if err != nil {
		log.Printf("%v error: %v", title, err)
	} else {
		err = cur.All(ctx, &many)
		if err != nil {
			log.Printf("%v error: %v", title, err)
		} else {
			log.Printf("%v success: %v", title, many)
		}
	}
}
