package controller

import (
	"OceanLearn/common"
	"OceanLearn/model"
	"OceanLearn/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
)

type ICategoryController interface {
	RestController
}

type CategoryController struct {
	DB *gorm.DB
}

func NewCategoryController() ICategoryController {
	db := common.GetDB()
	db.AutoMigrate(model.Category{})

	return CategoryController{DB: db}
}

func (c CategoryController) Create(ctx *gin.Context) {
	var requestCategory model.Category
	ctx.Bind(&requestCategory)
	fmt.Printf("requestCategory: %s\n", &requestCategory)

	if requestCategory.Name == "" {
		response.Fail(ctx, gin.H{}, "数据验证错误，分类名称必须填写")
		return
	}

	c.DB.Create(&requestCategory)

	response.Success(ctx, gin.H{"category": requestCategory}, "")
}

func (c CategoryController) Update(ctx *gin.Context) {
	// 绑定body中的参数
	var requestCategory model.Category
	ctx.Bind(&requestCategory)

	if requestCategory.Name == "" {
		response.Fail(ctx, gin.H{}, "数据验证错误，分类名称必须填写")
		return
	}

	// 获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	var updateCategory model.Category
	if c.DB.First(&updateCategory, categoryId).RecordNotFound() {
		response.Fail(ctx, gin.H{}, "更新出错：分类不存在")
		return
	}

	// 更新分类
	// map
	// struct
	// name value
	c.DB.Model(&updateCategory).Update("name", requestCategory.Name)

	response.Success(ctx, gin.H{"category": updateCategory}, "修改成功")
}

func (c CategoryController) Show(ctx *gin.Context) {
	// 获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	var category model.Category
	if c.DB.First(&category, categoryId).RecordNotFound() {
		response.Fail(ctx, gin.H{}, "读取错误：分类不存在")
		return
	}

	response.Success(ctx, gin.H{"category": category}, "")
}

func (c CategoryController) Delete(ctx *gin.Context) {
	// 获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	if err := c.DB.Delete(model.Category{}, categoryId).Error; err != nil {
		response.Fail(ctx, gin.H{}, "删除失败")
		return
	}

	response.Success(ctx, gin.H{}, "删除成功")

}
