package global

import (
	"database/sql"

	"github.com/anle/codebase/pkg/logger"
	"github.com/anle/codebase/setting"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
)

var (
	Config    setting.Config
	Logger    *logger.LoggerZap
	Rdb       *redis.Client
	Mdbc      *sql.DB
	Validator *validator.Validate
)
