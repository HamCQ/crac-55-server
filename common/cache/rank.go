package cache

import (
	"crac55/app/server/entities"
	"errors"
	"fmt"
	"math"

	goredis "github.com/go-redis/redis"
)

var (
	UpdateFlag = "CRAC_LOG_UPDATE"
	RunStatus  = "CRAC_LOG_A_RUN"
)

// KeyCnDiffCra 中国电台 - 区分 CRA 电台
func KeyCnDiffCra(year string) string {
	return fmt.Sprintf("CN_DIFF_CRA_%s", year)
}

// KeyCnCra 中国电台 - 不区分 CRA 电台
func KeyCnCra(year string) string {
	return fmt.Sprintf("CN_CRA_%s", year)
}

// KeyGlobleDiffCra 国际电台 - 区分 CRA 电台
func KeyGlobleDiffCra(year string) string {
	return fmt.Sprintf("GLOBLE_DIFF_CRA_%s", year)
}

// KeyGlobleCra 国际电台 - 不区分 CRA 电台
func KeyGlobleCra(year string) string {
	return fmt.Sprintf("GLOBLE_CRA_%s", year)
}

// RankTop5  前5排序
func RankTop5(key string) ([]entities.AnalyseRankContent, error) {
	var (
		res []entities.AnalyseRankContent
	)
	//第0位-更新时间标注
	v, err := RedisClient.ZRevRange(key, 1, 5).Result()
	if err != nil && !errors.Is(err, goredis.Nil) {
		return res, err
	}
	for _, v := range v {
		score, err := RedisClient.ZScore(key, v).Result()
		if err != nil {
			return res, err
		}
		res = append(res, entities.AnalyseRankContent{
			Callsign: v,
			Score:    int(score),
		})
	}
	return res, nil
}

// RankUpdateTime 排序更新时间
func RankUpdateTime(year string) (int, error) {
	key := KeyCnDiffCra(year)
	v, err := RedisClient.ZRevRange(key, 0, 1).Result()
	if err != nil && !errors.Is(err, goredis.Nil) {
		return 0, err
	}
	if len(v) == 0 {
		return 0, nil
	}
	score, err := RedisClient.ZScore(key, v[0]).Result()
	if err != nil {
		return 0, err
	}
	return int(score), nil
}

// RankAll 返回所有
func RankAll(key string, page int) (entities.AnalyseRankRes, error) {
	var (
		res      entities.AnalyseRankRes
		err      error
		pagesize int64 = 100
	)
	allCount := RedisClient.ZCard(key).Val()
	res.Total = allCount
	res.Page = page
	res.PageSize = int(pagesize)
	res.PageCount = int(math.Ceil(float64(allCount) / float64(pagesize)))

	start := (int64(page) - 1) * pagesize
	end := (start + pagesize) - 1
	if page == 0 {
		start = 0
		end = allCount
	}
	v, err := RedisClient.ZRevRange(key, start, end).Result()
	if err != nil && !errors.Is(err, goredis.Nil) {
		return res, err
	}
	for _, v := range v {
		score, err := RedisClient.ZScore(key, v).Result()
		if err != nil {
			return res, err
		}
		index, err := RedisClient.ZRank(key, v).Result()
		if err != nil {
			return res, err
		}
		res.Data = append(res.Data, entities.AnalyseRankContent{
			Number:   int((allCount - index)) - 1,
			Callsign: v,
			Score:    int(score),
		})
	}
	return res, nil
}

// IndexRank 获取排名
func GetRank(key, callsign string) (int64, error) {
	allCount := RedisClient.ZCard(key).Val()
	index, err := RedisClient.ZRank(key, callsign).Result()
	if err != nil && !errors.Is(err, goredis.Nil) {
		return 0, err
	}
	if errors.Is(err, goredis.Nil) {
		return 0, nil
	}
	return (allCount - index) - 1, nil
}
