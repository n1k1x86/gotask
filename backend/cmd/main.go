package main

import (
	"backend/pkg/common/db"
	"backend/pkg/tasks"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("../pkg/common/envs/.env")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal(err)
	}

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	r := gin.Default()
	d := db.Init(dbUrl)

	tasks.RegisterRouters(r, d)

	r.Run(port)
}
