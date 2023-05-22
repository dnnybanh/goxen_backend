package main

import (
	"github.com/dnnybanh/goxen_backend/db"
	"github.com/dnnybanh/goxen_backend/handlers/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./utils/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	router := gin.Default()
	router.Use(cors.Default())
	handler := db.Init(dbUrl)

	user.InitRoutes(router, handler)

	router.Run(port)
}