package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
	logger.InitLog(consts.LoginApiLogFile)
	config.InitConfigFile()

	if false == viper.GetBool("is_debug") {
		fmt.Print("Read salt from config: ")
		utils.Salt = viper.GetString("salt")
		if utils.Hash1(utils.Salt) != viper.GetString("salt_hashed") {
			panic("salt verification failed!")
		}
	}

	db.InitDb()
	err := db.GetDb(false).
		AutoMigrate(&structs.User{}, &structs.VerificationCode{}, &structs.Post{},
			&structs.Comment{}, &structs.Attention{}, &structs.Report{}, &structs.SystemMessage{}, structs.Ban{})
	utils.FatalErrorHandle(&err, "error migrating database!")

	log.Println("start time: ", time.Now().Format("01-02 15:04:05"))
	if false == viper.GetBool("is_debug") {
		gin.SetMode(gin.ReleaseMode)
	}

	route.LoginApiListenHttp()
}
