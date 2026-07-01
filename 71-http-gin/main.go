package main

import (
	"flag"
	"http-demo/handlers"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	PORT string
)

func init() {
	slog.Info("fetching APP_PORT from environment")
	PORT = os.Getenv("APP_PORT")
	// the logic to check the port can be written here buy just to demonstrate given in another init func
}

func init() {
	if PORT == "" {
		slog.Info("unable to fetch APP_PORT from environment,checking command line argument if not giving a default value")
		//PORT = "8080"

		flag.StringVar(&PORT, "port", "8080", "--port=8080 | --port 8080 | -port=8080 | -port 8080")
		//PORT = flag.String("port", "8080", "--port=8080 | --port 8080 | -port=8080 | -port 8080")
		flag.Parse()
	}
}

func main() {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	slog.SetDefault(logger)

	router := gin.Default()

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello World!"))
	// })

	// http.HandleFunc("/ping", handlers.Ping)
	// http.HandleFunc("/health", handlers.Health)
	// http.HandleFunc("/user", handlers.CreateUser)

	router.GET("/", func(ctx *gin.Context) {
		slog.Info("root rounte is called")
		ctx.String(200, "Hello World")
	})

	router.GET("/v1/api/public/ping", handlers.Ping)
	router.GET("/v1/api/public/health", handlers.Health)

	//privateRoute := router.Group("/v1/api/private", handlers.Auth)
	privateRoute := router.Group("/v1/api/private")
	user := handlers.Mewuser("users.txt")

	privateRoute.POST("/users", handlers.Auth, user.CreateUser)
	privateRoute.GET("/users", user.GetUsers)

	slog.Info("Server is up and running on the port:" + PORT)

	if err := router.Run("0.0.0.0:" + PORT); err != nil {
		slog.Error(err.Error())
	}

}
