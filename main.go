package main

import (
	"OceanLearn/common"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	DB := common.InitDB()
	defer DB.Close()

	r := gin.Default()
	r = CollectRoute(r)
	panic(r.Run())
}
