package main

import (
	"flag"
	"http-demo/handlers"
	"log/slog"
	"net/http"
	"os"
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

	// file, err := os.OpenFile("data.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	// if err != nil {
	// 	slog.Error(err.Error())
	// 	//os.Exit(2)
	// 	runtime.Goexit()
	// }

	// defer file.Close()
	//logger := slog.New(slog.NewJSONHandler(file, &slog.HandlerOptions{Level: slog.LevelDebug}))
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	slog.SetDefault(logger)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	http.HandleFunc("/ping", handlers.Ping)
	http.HandleFunc("/health", handlers.Health)
	http.HandleFunc("/user", handlers.CreateUser)

	slog.Info("Server is up and running on the port:" + PORT)
	if err := http.ListenAndServe("0.0.0.0:"+PORT, nil); err != nil {
		slog.Error(err.Error())
	}

}
