package main

import (
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
	"log"
	"time"
	"treehollow-backend/pkg/config"
	"treehollow-backend/pkg/consts"
	"treehollow-backend/pkg/db"
	"treehollow-backend/pkg/logger"
	"treehollow-backend/pkg/route"
	"treehollow-backend/pkg/structs"
	"treehollow-backend/pkg/utils"
)

func main() {
	logger.InitLog(consts.ServicesApiLogFile)
	config.InitConfigFile()

	db.InitDb()
	err := db.GetDb(false).
		AutoMigrate(&structs.User{}, &structs.VerificationCode{}, &structs.Post{},
			&structs.Comment{}, &structs.Attention{}, &structs.Report{}, &structs.SystemMessage{}, structs.Ban{})
	utils.FatalErrorHandle(&err, "error migrating database!")

	log.Println("start time: ", time.Now().Format("01-02 15:04:05"))
	if false == viper.GetBool("is_debug") {
		gin.SetMode(gin.ReleaseMode)
	}

	route.HotPosts, _ = db.GetHotPosts()
	c := cron.New()
	_, _ = c.AddFunc("*/1 * * * *", func() {
		route.HotPosts, _ = db.GetHotPosts()
		//log.Println("refreshed hotPosts ,err=", err)
	})
	c.Start()

	route.ServicesApiListenHttp()
}
