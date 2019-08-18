package routes

import (
	"net/http"
	"html/template"
	"github.com/gin-gonic/gin"
)

func InitPages(r *gin.Engine) {
	r.Static("/templates", "./templates")
	r.StaticFS("/static", http.Dir("templates"))

	r.LoadHTMLFiles(
		"templates/login.html",
		"templates/registration.html",
		"templates/index.html",
		"templates/admin/layout.html",
		"templates/admin/dashboard.html",
		"templates/admin/create_application.html",
		"templates/admin/applications.html",
	)

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Sitius inc",
		})
	})
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title": "Login",
		})
	})
	r.GET("/registration", func(c *gin.Context) {
		c.HTML(http.StatusOK, "registration.html", gin.H{
			"title": "Registration",
		})
	})

	admin := r.Group("admin")
	admin.Use(ginAuthJWT.MiddlewareFunc())
	{
		admin.GET("/dashboard", func(c *gin.Context) {
			ct := "templates/admin/dashboard.html"
			buildAdminPages(c, r, ct, gin.H{
				"title": "Dashboard",
			})
		})
		admin.GET("/applications", func(c *gin.Context) {
			ct := "templates/admin/applications.html"
			buildAdminPages(c, r, ct, gin.H{
				"title": "Applications",
			})
		})
		admin.GET("/create_application", func(c *gin.Context) {
			ct := "templates/admin/create_application.html"
			buildAdminPages(c, r, ct, gin.H{
				"title": "Create application",
			})
		})
	}
}

func buildAdminPages(c *gin.Context, r *gin.Engine, contentTemplate string, data gin.H) {
	var baseTemplate = "templates/admin/layout.html"
	r.SetHTMLTemplate(template.Must(template.ParseFiles(baseTemplate, contentTemplate)))
	c.HTML(200, "base", data)
}
