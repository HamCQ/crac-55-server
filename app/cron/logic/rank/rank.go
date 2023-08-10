package rank

import (
	"crac55/common/cache"
	"crac55/common/clog"
	"crac55/common/db"

	goredis "github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	CN_DIFF_CRA     = "CN_DIFF_CRA" + "_2023"
	GLOBLE_DIFF_CRA = "GLOBLE_DIFF_CRA" + "_2023"
	CN_CRA          = "CN_CRA" + "_2023"
	GLOBLE_CRA      = "GLOBLE_CRA" + "_2023"
	UpdateFlag      = "CRAC_LOG_UPDATE"
	RunStatus       = "CRAC_LOG_A_RUN"
)

// CracLogRankCron 呼号排名定时脚本
type CracLogRankCron struct {
	log *logrus.Logger
	db  *gorm.DB
	rds *goredis.Client
}

// NewCracLogRankCron new
func NewCracLogRankCron() *CracLogRankCron {
	return &CracLogRankCron{
		log: clog.Log(),
		db:  db.DB,
		rds: cache.RedisClient,
	}
}
