package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	//  "strconv"
)

func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("db nil")
	}

	return db
}

// Main function. Starts `gin` http server.
func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	// Global middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	db := InitDB("/home/alex/workspace_go/masstransit.db")
	defer db.Close()

	r.GET("/stops", func(c *gin.Context) {
		stops := ReadStop(db)
		c.JSON(200, stops)
	})

	r.GET("/routes", func(c *gin.Context) {
		routes := ReadRoute(db)
		c.JSON(200, routes)
	})

	r.GET("/calendars", func(c *gin.Context) {
		routes := ReadCalendars(db)
		c.JSON(200, routes)
	})

	r.GET("/stops/:stop/schedules", func(c *gin.Context) {
		//stop, _ := strconv.Atoi(c.Param("stop"))
		stop := c.Param("stop")
		schedules := ReadSchedule(db, stop)
		c.JSON(200, schedules)
	})

	r.GET("/routes/:route", func(c *gin.Context) {
		//route, _ := strconv.Atoi(c.Param("route"))
		route := c.Param("route")
		routeShapes := ReadRouteShape(db, route)
		c.JSON(200, routeShapes)
	})

	r.Run(":8888")
}
