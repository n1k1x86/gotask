package main

import (
	"backend/pkg/common/db"
	"backend/pkg/tasks"
	"backend/pkg/tasksgroup"
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

	port := viper.Get("port").(string)
	host := viper.Get("host").(string)
	user := viper.Get("user").(string)
	password := viper.Get("password").(string)
	db_name := viper.Get("dbname").(string)
	sslmode := viper.Get("sslmode").(string)

	r := gin.Default()
	d := db.Init(host, user, password, db_name, sslmode)

	defer d.Close()

	tasks.RegisterRouters(r, d)
	tasksgroup.RegisterRoutes(r, d)

	r.Run(port)
}
