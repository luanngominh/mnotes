package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-kit/kit/log"

	"github.com/luanngominh/mnotes/config"
	_ "github.com/luanngominh/mnotes/config"
	"github.com/luanngominh/mnotes/db"
	"github.com/luanngominh/mnotes/endpoints"
	mnotesHttp "github.com/luanngominh/mnotes/http"
	"github.com/luanngominh/mnotes/service"
	noteSvc "github.com/luanngominh/mnotes/service/note"
	userSvc "github.com/luanngominh/mnotes/service/user"
)

func main() {
	// Create logger
	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	//Setup location
	local, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}
	time.Local = local

	//Connect to db
	pgDB, pgClose := db.New(fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		config.Cfg.DBHost, config.Cfg.DBPort,
		config.Cfg.DBUser, config.Cfg.DBName,
		config.Cfg.DBPassword, "disable"))

	//Create endpoint service
	s := service.Service{
		UserSerivce: service.Compose(
			userSvc.NewPGService(pgDB),
			userSvc.ValidationMiddleware(),
		).(userSvc.Service),
		NoteService: service.Compose(
			noteSvc.NewPGService(pgDB),
			noteSvc.ValidationMiddleware(),
		).(noteSvc.Service),
	}

	defer pgClose()

	var h http.Handler
	h = mnotesHttp.NewHTTPHandler(
		endpoints.MakeServerEndpoints(s),
		logger,
		//true to use cors
		true,
	)

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	httpAddr := fmt.Sprintf("%s:%s", config.Cfg.Address, config.Cfg.Port)
	fmt.Println(httpAddr)
	go func() {
		logger.Log("transport", "HTTP", "addr", httpAddr)
		errs <- http.ListenAndServe(httpAddr, h)
	}()

	logger.Log("exit", <-errs)
}
