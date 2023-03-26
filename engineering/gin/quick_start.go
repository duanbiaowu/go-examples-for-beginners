package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// curl http://127.0.0.1:8080
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello World")
	})

	// curl http://127.0.0.1:8080/user/geektutu
	r.GET("/users/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.String(http.StatusOK, "Hello %s", name)
	})

	// curl "http://127.0.0.1:8080/users?name=Tom&role=student"
	r.GET("/users", func(ctx *gin.Context) {
		name := ctx.Query("name")
		role := ctx.DefaultQuery("role", "doctor")
		ctx.String(http.StatusOK, "%s is a %s", name, role)
	})

	// curl http://127.0.0.1:8080/form  -X POST -d 'username=geektutu&password=1234'
	// {"password":"1234","username":"geektutu"}
	r.POST("/form", func(ctx *gin.Context) {
		username := ctx.PostForm("username")
		password := ctx.DefaultPostForm("password", "000000")

		ctx.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})

	// curl "http://127.0.0.1:8080/posts?id=9876&page=7"  -X POST -d 'username=geektutu&password=1234'
	// {"id":"9876","page":"7","password":"1234","username":"geektutu"}
	r.POST("/posts", func(ctx *gin.Context) {
		id := ctx.Query("id")
		page := ctx.DefaultQuery("page", "0")
		username := ctx.PostForm("username")
		password := ctx.DefaultPostForm("username", "000000")

		ctx.JSON(http.StatusOK, gin.H{
			"id":       id,
			"page":     page,
			"username": username,
			"password": password,
		})
	})

	// curl -g "http://127.0.0.1:8080/post?ids[Jack]=001&ids[Tom]=002" -X POST -d 'names[a]=Sam&names[b]=David'
	// {"ids":{"Jack":"001","Tom":"002"},"names":{"a":"Sam","b":"David"}}
	r.POST("/post", func(ctx *gin.Context) {
		ids := ctx.QueryMap("ids")
		names := ctx.PostFormMap("names")

		ctx.JSON(http.StatusOK, gin.H{
			"ids":   ids,
			"names": names,
		})
	})

	// curl -i http://127.0.0.1:8080/redirect
	r.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/index")
	})

	// curl "http://127.0.0.1:8080/goindex"
	r.GET("/goindex", func(c *gin.Context) {
		c.Request.URL.Path = "/"
		r.HandleContext(c)
	})

	defaultHandler := func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"path": c.FullPath(),
		})
	}

	// curl http://127.0.0.1:8080/v1/posts
	v1 := r.Group("/v1")
	{
		v1.GET("/posts", defaultHandler)
		v1.GET("/series", defaultHandler)
	}

	// curl http://127.0.0.1:8080/v2/posts
	v2 := r.Group("/v2")
	{
		v2.GET("/posts", defaultHandler)
		v2.GET("/series", defaultHandler)
	}

	r.POST("/upload1", func(ctx *gin.Context) {
		file, _ := ctx.FormFile("file")
		// c.SaveUploadedFile(file, dst)
		ctx.String(http.StatusOK, "%s uploaded!", file.Filename)
	})

	r.POST("/upload2", func(ctx *gin.Context) {
		// Multipart form
		form, _ := ctx.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)
			// ctx.SaveUploadedFile(file, dst)
		}
		ctx.String(http.StatusOK, "%d files uploaded!", len(files))
	})

	_ = r.Run("127.0.0.1:8080")
}
