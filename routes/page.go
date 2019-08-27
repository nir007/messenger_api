package routes

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"html/template"
	"messenger/application"
	"net/http"
	"strings"
)

func InitPages(r *gin.Engine) {
	r.Static("/templates", "./templates")
	r.StaticFS("/static", http.Dir("templates"))

	r.LoadHTMLFiles(
		"templates/layout.html",
		"templates/login.html",
		"templates/registration.html",
		"templates/index.html",
		"templates/admin/layout.html",
		"templates/admin/dashboard.html",
		"templates/admin/create_application.html",
		"templates/admin/applications.html",
	)

	r.GET("/", func(c *gin.Context) {
		ct := "templates/index.html"
		buildSitePages(c, r, ct, gin.H{
			"title": "Messenger API",
		})
	})
	r.GET("/login", func(c *gin.Context) {
		ct := "templates/login.html"
		buildSitePages(c, r, ct, gin.H{
			"title": "Login",
		})
	})
	r.GET("/registration", func(c *gin.Context) {
		ct := "templates/registration.html"
		buildSitePages(c, r, ct, gin.H{
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
		admin.GET("/application/:id", func(c *gin.Context) {
			id := c.Param("id")

			objId, _ := primitive.ObjectIDFromHex(id)
			app := &application.Application{ID: objId}
			err := app.FindOne(c)

			if err != nil {
				buildErrorPage(c, r, err)
				return
			}

			ct := "templates/admin/application.html"
			buildAdminPages(c, r, ct, gin.H{
				"title":       "ApplicationID: " + id,
				"appId":       id,
				"name":        app.Name,
				"description": app.Description,
				"secret":      app.Secret,
				"createdAt":   app.CreatedAt,
				"updatedAt":   app.UpdatedAt,
				"domains":     strings.Join(app.Domains, ","),
			})
		})
	}
}

func buildSitePages(c *gin.Context, r *gin.Engine, contentTemplate string, data gin.H) {
	var baseTemplate = "templates/layout.html"
	r.SetHTMLTemplate(template.Must(template.ParseFiles(baseTemplate, contentTemplate)))
	c.HTML(200, "base", data)
}

func buildAdminPages(c *gin.Context, r *gin.Engine, contentTemplate string, data gin.H) {
	var baseTemplate = "templates/admin/layout.html"
	r.SetHTMLTemplate(template.Must(template.ParseFiles(baseTemplate, contentTemplate)))
	c.HTML(200, "base", data)
}

func buildErrorPage(c *gin.Context, r *gin.Engine, err error) {
	var baseTemplate = "templates/admin/layout.html"
	r.SetHTMLTemplate(template.Must(template.ParseFiles(baseTemplate, "templates/admin/error.html")))
	c.HTML(200, "base", gin.H{"err": err.Error()})
}
