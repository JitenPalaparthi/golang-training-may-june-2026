package main

import (
	"flag"
	"log/slog"
	"net"
	"os"
	"runtime"
	"simple-commerce/internal/database"
	grpchandlers "simple-commerce/internal/grpc_handlers"
	"simple-commerce/internal/models"
	productspb "simple-commerce/internal/protos"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)

var (
	PORT string
	DSN  string // dsn := `host=localhost user=postgres password=postgres dbname=simple-commerce port=5432 sslmode=disable TimeZone=Asia/Shanghai`

)

func init() {
	slog.Info("fetching GRPC_PORT from environment")
	PORT = os.Getenv("GRPC_PORT")
	DSN = os.Getenv("DSN_URL")
	// the logic to check the port can be written here buy just to demonstrate given in another init func
}

func init() {
	if PORT == "" {
		slog.Info("unable to fetch GRPC_PORT from environment,checking command line argument if not giving a default value")
		//PORT = "8080"
		flag.StringVar(&PORT, "port", "58080", "--port=58080 | --port 58080 | -port=58080 | -port 58080")
		//PORT = flag.String("port", "8080", "--port=8080 | --port 8080 | -port=8080 | -port 8080")
	}

	if DSN == "" {
		flag.StringVar(&DSN, "dsn", `host=localhost user=postgres password=postgres dbname=simple-commerce port=5432 sslmode=disable TimeZone=Asia/Shanghai`, "--dsn=host=localhost user=postgres password=postgres dbname=simple-commerce port=5432 sslmode=disable TimeZone=Asia/Shanghai")
	}
	flag.Parse()
}

func main() {
	// logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	// slog.SetDefault(logger)

	db, err := database.Connect(DSN)
	if err != nil {
		slog.Error(err.Error())
		runtime.Goexit()
	}

	if err := Init(db); err != nil {
		slog.Error(err.Error())
		runtime.Goexit()
	}

	lister, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		slog.Error(err.Error())
		runtime.Goexit()
	}

	grpcServer := grpc.NewServer()
	productServer := grpchandlers.NewProductGRPCServer(db)

	productspb.RegisterProductServiceServer(grpcServer, productServer)
	reflection.Register(grpcServer)

	slog.Info("gRPC server listning on port:" + PORT)

	if err := grpcServer.Serve(lister); err != nil {
		slog.Error("Failed to stat gPRC server " + err.Error())
		runtime.Goexit()
	}

}

func Init(db *gorm.DB) error {
	return db.AutoMigrate(&models.Product{})

}
