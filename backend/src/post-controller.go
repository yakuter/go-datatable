package src

import (
	"fmt"
	"go-datatable/backend/inc"
	"go-datatable/backend/model"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetPosts(c *gin.Context) {
	db = inc.GetDB()

	table := "posts"
	var total, filtered int
	var posts []model.Post

	query := db.Table(table)
	query = query.Offset(QueryOffset(c))
	query = query.Limit(QueryLimit(c))
	query = query.Order(QueryOrder(c))
	query = query.Scopes(SearchScope(c), DateTimeScope(c))

	if err := query.Find(&posts).Error; err != nil {
		c.AbortWithStatus(404)
		log.Println(err)
		return
	}

	// Filtered data count
	query = query.Offset(0)
	query.Table(table).Count(&filtered)

	// Total data count
	db.Table(table).Count(&total)

	result := model.DataResult{
		total,
		filtered,
		posts,
	}

	c.JSON(200, result)
}

func SearchScope(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		query := db
		search := c.QueryMap("search")
		fmt.Println(search)
		if search["value"] != "" {
			query = query.Where("title ILIKE ? ", "%"+search["value"]+"%")
			query = query.Or("slug ILIKE ? ", "%"+search["value"]+"%")
			query = query.Or("excerpt ILIKE ? ", "%"+search["value"]+"%")
			query = query.Or("category ILIKE ? ", "%"+search["value"]+"%")
		}
		return query
	}
}
