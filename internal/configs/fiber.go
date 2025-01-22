package configs

import (
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
)

var FbrCfg = fiber.New(fiber.Config{
	ServerHeader:  "bookshelf-api",
	AppName:       "bookshelf-api",
	JSONEncoder:   sonic.Marshal,
	JSONDecoder:   sonic.Unmarshal,
	Prefork:       true,
	StrictRouting: true,
})
