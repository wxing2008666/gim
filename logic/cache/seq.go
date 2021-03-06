package cache

import (
	"gim/logic/db"
	"gim/public/logger"
	"strconv"
)

const (
	UserSeqKey  = "user_seq:"
	GroupSeqKey = "group_seq"
)

type seqCache struct{}

var SeqCache = new(seqCache)

func (*seqCache) UserKey(appId, userId int64) string {
	return UserSeqKey + strconv.FormatInt(appId, 10) + ":" + strconv.FormatInt(userId, 10)
}

func (*seqCache) GroupKey(appId, groupId int64) string {
	return GroupSeqKey + strconv.FormatInt(appId, 10) + ":" + strconv.FormatInt(groupId, 10)
}

// Incr 将序列号增加1
func (c *seqCache) Incr(key string) (int64, error) {
	seq, err := db.RedisCli.Incr(key).Result()
	if err != nil {
		logger.Sugar.Error(err)
		return 0, err
	}
	return seq, nil
}
