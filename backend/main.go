package main

import (
	config "github.com/ahmadzakyarifin/echo-pos-admin/internal/infrastructure"
	"github.com/ahmadzakyarifin/echo-pos-admin/internal/router"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	server := echo.New()

	// Middleware
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173", "http://localhost:5174", "http://localhost:3000"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))


	db := config.ConnettDB()
	defer db.Close()

	// Setup Routes
	router.SetupAuthRouter(db, server)

	if err := server.Start(":8080"); err != nil {
		server.Logger.Fatal(err)
	}
}
