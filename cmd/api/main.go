package main

import (
	"clean_architecture"
	"fmt"
	"os"
	"runtime/debug"
	"time"

	"clean_architecture/pkg/logrus"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiber_logger "github.com/gofiber/fiber/v2/middleware/logger"
	fiber_recover "github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	logrus.Init()
	logrus := logrus.GetLogger()

	app, cleanup, err := clean_architecture.InitializeApp()
	if err != nil {
		logrus.Fatal("Failed to initialize app:", err)
	}
	defer cleanup()

	// Create Fiber instance
	server := fiber.New(fiber.Config{
		AppName:           "Golang ORM API",
		EnablePrintRoutes: true,
		Prefork:           true,
	})

	var fiberLogConfig = fiber_logger.Config{
		Next:          nil,
		Done:          nil,
		Format:        "\u001b[36m${time}\u001b[0m | \u001b[32m${status}\u001b[0m | \u001b[33m${latency}\u001b[0m | \u001b[35m${ip}\u001b[0m | \u001b[34m${method}\u001b[0m | \u001b[37m${path}\u001b[0m | \u001b[31m${error}\u001b[0m\n",
		TimeFormat:    "15:04:05",
		TimeZone:      "Asia/Jakarta",
		TimeInterval:  500 * time.Millisecond,
		Output:        os.Stdout,
		DisableColors: false,
	}

	// Middleware
	// server.Use(recover.New(recover.Config{
	// 	EnableStackTrace: true,
	// 	StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
	// 		logrus.Errorf("PANIC: %v\n%s", e, debug.Stack())
	// 	},
	// }))

	server.Use(fiber_recover.New(fiber_recover.Config{
		EnableStackTrace: true,
		StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
			fmt.Fprintf(os.Stderr, "\033[31mPANIC: %v\n%s\033[0m\n", e, debug.Stack())
		},
	}))
	server.Use(fiber_logger.New(fiberLogConfig))
	server.Use(cors.New())

	// Health check
	server.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "Server is running",
		})
	})

	// API routes
	api := server.Group("/api/v1")
	users := api.Group("/users")
	{
		users.Get("/", app.Handlers.User.ListUser)
		users.Post("/create", app.Handlers.User.CreateUser)
		users.Get("/:id", app.Handlers.User.GetUser)
		users.Put("/:id", app.Handlers.User.UpdateUser)
		users.Delete("/:id", app.Handlers.User.DeleteUser)
	}

	// Start server
	port := app.Config.Server.Port
	logrus.Infof("Server starting on port %s", port)
	if err := server.Listen(fmt.Sprintf(":%s", port)); err != nil {
		logrus.Fatal("Failed to start server:", err)
	}
}
