package main

import (
	"flag"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

var (
	port string
)

func main() {
	flag.StringVar(&port, "port", "3001", "port of server")
	flag.Parse()
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	store := sessions.NewCookieStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")
	router.GET("/", handleIndex)
	router.GET("/write", func(c *gin.Context) {
		c.HTML(http.StatusOK, "edit.tmpl", MainView{})
	})
	router.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.tmpl", MainView{})
	})
	router.GET("/stories", func(c *gin.Context) {
		c.HTML(http.StatusOK, "write.tmpl", MainView{})
	})
	router.GET("/read", func(c *gin.Context) {
		c.HTML(http.StatusOK, "read.tmpl", MainView{})
	})
	router.GET("/archive", func(c *gin.Context) {
		c.HTML(http.StatusOK, "archive.tmpl", MainView{})
	})
	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.tmpl", MainView{})
	})
	router.GET("/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.tmpl", MainView{})
	})
	router.GET("/admin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin.tmpl", MainView{})
	})
	router.GET("/terms", func(c *gin.Context) {
		c.HTML(http.StatusOK, "terms.tmpl", MainView{})
	})
	router.GET("/privacy", func(c *gin.Context) {
		c.HTML(http.StatusOK, "privacy.tmpl", MainView{})
	})
	router.GET("/favicon.ico", func(c *gin.Context) {
		c.Redirect(302, "/static/img/meta/favicon.ico")
	})
	router.Run(":" + port)
}

type MainView struct {
	Title   string
	Landing bool
}

func handleIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "landing.tmpl", MainView{
		Landing: true,
	})
}
