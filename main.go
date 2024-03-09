package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/adity37/task/repository/auth"
	"github.com/adity37/task/repository/mysql"
	"github.com/adity37/task/repository/redis"
	"github.com/adity37/task/router"
	"github.com/adity37/task/service"
	"github.com/adity37/task/transport"
	getenv "github.com/aditya37/get-env"
)

type config struct {
	servicePort int
	redisConfig redis.Config
	mysqlConfig mysql.MYSQLConfig
}

func main() {
	ctx := context.Background()

	// get env config
	config := GetConfig()

	// redis repo
	redisRepo, err := redis.NewRedis(ctx, config.redisConfig)
	if err != nil {
		log.Fatal(err)
		return
	}

	db, err := mysql.NewMYSQL(config.mysqlConfig)
	if err != nil {
		log.Fatal(err)
		return
	}

	// oauth...
	scope := []string{"https://www.googleapis.com/auth/userinfo.email"}
	redirectURL := getenv.GetString("OAUTH_REDIRECT", "http://localhost:2024/users/callback")
	oauth := auth.NewOauth(redirectURL, scope)

	// service
	svc := service.NewService(db, redisRepo, oauth)

	// transport
	tp := transport.NewTransport(svc)

	// router...
	router := router.NewRouter(tp)

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
		errs <- fmt.Errorf("%s", <-c)
	}()
	go func() {
		errs <- serve(router, config.servicePort)
	}()

	log.Fatalf("Stop server with error detail: %v", <-errs)

}

// GetConfig
func GetConfig() config {
	return config{
		servicePort: getenv.GetInt("SERVICE_PORT", 2024),
		redisConfig: redis.Config{
			Address:  getenv.GetString("REDIS_HOST", "127.0.0.1:6379"),
			Password: getenv.GetString("REDIS_PASSWORD", ""),
			DB:       getenv.GetInt("REDIS_DB", 0),
		},
		mysqlConfig: mysql.MYSQLConfig{
			Port:     getenv.GetInt("DB_PORT", 3306),
			Host:     getenv.GetString("DB_HOST", "127.0.0.1"),
			Name:     getenv.GetString("DB_NAME", "db_task"),
			Password: getenv.GetString("DB_PASSWORD", "root"),
			User:     getenv.GetString("DB_USER", "root"),
		},
	}
}

func serve(handler http.Handler, port int) error {
	addr := fmt.Sprintf(":%d", port)
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	h := &http.Server{
		Handler: handler,
	}
	log.Printf("service running on %s", addr)
	return h.Serve(listen)
}
