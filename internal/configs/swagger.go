package configs

import "github.com/gofiber/contrib/swagger"

var SwgCfg = swagger.Config{
	BasePath: "/",
	FilePath: "internal/server/docs/swagger.json",
	CacheAge: 60,
}
