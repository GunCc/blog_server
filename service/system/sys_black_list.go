package system

import (
	"blog_server/global"
	"blog_server/utils"
	"context"
)

type JwtService struct {
}

func (jwtService *JwtService) GetRedisJWT(userName string) (redisJWT string, err error) {
	redisJWT, err = global.BLOG_REDIS.Get(context.Background(), userName).Result()
	return redisJWT, err
}

func (jwtService *JwtService) SetRedisJWT(jwt string, userName string) (err error) {
	dr, err := utils.ParseDuration(global.BLOG_CONFIG.JWT.ExpiresTime)
	if err != nil {
		return err
	}
	timer := dr
	err = global.BLOG_REDIS.Set(context.Background(), userName, jwt, timer).Err()
	return err
}

// func (jwtService *JwtService) JsonInBlacklist(jwtList system.JwtBlackList) (err error) {
// 	err = global.BLOG_DB.Create(&jwtList).Error
// 	if err != nil {
// 		return
// 	}
// 	// global.BlackCache
// }
