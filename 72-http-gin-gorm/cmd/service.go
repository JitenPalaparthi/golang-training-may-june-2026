package main

import (
	"flag"
	"log/slog"
	"os"
	"runtime"
	"simple-commerce/internal/database"
	"simple-commerce/internal/handlers"
	"simple-commerce/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	PORT string
	DSN  string // dsn := `host=localhost user=postgres password=postgres dbname=simple-commerce port=5432 sslmode=disable TimeZone=Asia/Shanghai`

)

func init() {
	slog.Info("fetching APP_PORT from environment")
	PORT = os.Getenv("APP_PORT")
	DSN = os.Getenv("DSN_URL")
	// the logic to check the port can be written here buy just to demonstrate given in another init func
}

func init() {
	if PORT == "" {
		slog.Info("unable to fetch APP_PORT from environment,checking command line argument if not giving a default value")
		//PORT = "8080"
		flag.StringVar(&PORT, "port", "8080", "--port=8080 | --port 8080 | -port=8080 | -port 8080")
		//PORT = flag.String("port", "8080", "--port=8080 | --port 8080 | -port=8080 | -port 8080")
	}

	if DSN == "" {
		flag.StringVar(&DSN, "dsn", `host=localhost user=postgres password=postgres dbname=simple-commerce port=5432 sslmode=disable TimeZone=Asia/Shanghai`, "--dsn=host=localhost user=postgres password=postgres dbname=simple-commerce port=5432 sslmode=disable TimeZone=Asia/Shanghai")
	}
	flag.Parse()
}

func main() {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	slog.SetDefault(logger)

	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		slog.Info("root rounte is called")
		ctx.String(200, "Hello World")
	})

	router.GET("/v1/api/public/ping", handlers.Ping)
	router.GET("/v1/api/public/health", handlers.Health)

	db, err := database.Connect(DSN)
	if err != nil {
		slog.Error(err.Error())
		runtime.Goexit()
	}

	if err := Init(db); err != nil {
		slog.Error(err.Error())
		runtime.Goexit()
	}

	iproductDB := database.NewProductDB(db)
	productHandler := handlers.NewProductHandler(iproductDB)

	productRoute := router.Group("/v1/api/product")

	productRoute.POST("/", productHandler.CreateProduct)

	slog.Info("Server is up and running on the port:" + PORT)

	if err := router.Run("0.0.0.0:" + PORT); err != nil {
		slog.Error(err.Error())
	}

}

func Init(db *gorm.DB) error {
	return db.AutoMigrate(&models.Product{})

}
