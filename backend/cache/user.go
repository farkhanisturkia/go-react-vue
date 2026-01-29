package cache

import (
	"strconv"
	"time"

	"go-react/backend/database"
	"go-react/backend/models"
	"go-react/backend/pkg/redis"
)

const (
	UserListCachePrefix  = "users:list:page:"
	UserTotalCacheKey    = "users:total"
	UserListKeysSet      = "users:list:keys"
	UserCachePrefix      = "user:id:"
	CacheTTL             = 5 * time.Minute
	UserTTL              = 10 * time.Minute
)

func GetUsersTotalCount() int64 {
	val, err := redis.Client.Get(redis.Ctx, userTotalCacheKey).Result()
	if err == nil {
		if total, err := strconv.ParseInt(val, 10, 64); err == nil {
			return total
		}
	}

	var total int64
	database.DB.Model(&models.User{}).Count(&total)

	redis.Client.Set(redis.Ctx, userTotalCacheKey, strconv.FormatInt(total, 10), cacheTTL)

	return total
}

func InvalidateUserListCache() {
	redis.Client.Del(redis.Ctx, userTotalCacheKey)

	keys, err := redis.Client.SMembers(redis.Ctx, userListKeysSet).Result()
	if err == nil && len(keys) > 0 {
		redis.Client.Del(redis.Ctx, keys...)
	}

	redis.Client.Del(redis.Ctx, userListKeysSet)
}