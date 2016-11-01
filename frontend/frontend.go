package main

import (
	"github.com/goldbattle/raydart/frontend/app/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"html/template"
	"io"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}


func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Static files
	e.Static("/", "public")
	//e.Use(middleware.Static("/public/"))

	// Setup our templates
	t := &Template{
		templates: template.Must(template.ParseGlob("app/view/*.html")),
	}
	e.SetRenderer(t)

	// Routes
	e.GET("/", controller.HomeGET)
	e.GET("/origin/:stream/", controller.StreamGET)
	//e.GET("/repo/:hash/tree/*", controller.RepoGET)
	//e.GET("/repo/:hash/blob/*", controller.RepoGET)

	// Start server
	e.Run(standard.New(":6969"))
}
