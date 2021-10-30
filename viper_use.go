package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config") //配置文件名称
	viper.SetConfigType("yaml")   //配置文件类型
	viper.AddConfigPath(".")      //配置文件路径
	err := viper.ReadInConfig()   //读取配置文件信息
	if err != nil {
		panic(fmt.Errorf("Fatal error config file :%s \n", err))
	}
	r := gin.Default()
	if err := r.Run(fmt.Sprintf(":%d", viper.Get("port"))); err != nil {
		panic(err)
	}
}
