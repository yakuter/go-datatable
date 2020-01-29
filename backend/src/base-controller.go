package src

import (
	"go-datatable/backend/inc"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	// Create DB connection
	db = inc.InitDB()
	defer db.Close()
}

type Data struct {
	Items string
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func QueryOffset(c *gin.Context) int {
	offset := c.Query("start")
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		offsetInt = 0
	}
	return offsetInt
}

func QueryLimit(c *gin.Context) int {
	limit := c.Query("length")
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 25
	}
	return limitInt
}

func QueryOrder(c *gin.Context) string {
	columnMap := map[string]string{
		"0": "id",
		"1": "timestamp",
		"2": "title",
		"3": "slug",
		"4": "excerpt",
		"5": "category",
	}

	column := c.DefaultQuery("order[0][column]", "0")
	dir := c.DefaultQuery("order[0][dir]", "desc")
	orderString := columnMap[column] + " " + dir

	return orderString
}

func DateTimeScope(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		begin := c.DefaultQuery("begin", "")
		end := c.DefaultQuery("end", "")

		if begin == "" && end == "" {
			return db
		}

		var tBegin, tEnd time.Time
		layout := "02.01.2006 15:04"

		if begin != "" {
			tBegin, _ = time.Parse(layout, begin)
		} else {
			t := time.Now()
			tBegin = t.AddDate(-20, 0, 0)
		}

		if end != "" {
			tEnd, _ = time.Parse(layout, end)
		} else {
			tEnd = time.Now()
		}

		return db.Where("timestamp BETWEEN ? AND ?", tBegin, tEnd)
	}
}
