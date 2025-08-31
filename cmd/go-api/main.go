package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ayushmehta03/go-api/internal/config"
	"github.com/ayushmehta03/go-api/internal/http/handlers/student"
	"github.com/ayushmehta03/go-api/internal/storage/sqllite"
)




func main(){

	cfg:=config.MustLoad()



	storage,err:=sqllite.New(cfg)
	if err!=nil{
		log.Fatal(err)
	}

	slog.Info("stroage intialiazed",slog.String("env",cfg.Env))


	router:=http.NewServeMux();
	router.HandleFunc("POST /api/students",student.New(storage))
	router.HandleFunc("GET /api/students/{id}",student.GetById(storage))
	ch:=make(chan os.Signal,1)

	signal.Notify(ch,os.Interrupt,syscall.SIGINT,syscall.SIGTERM)

	server:=http.Server{
		Addr: cfg.Addr,
		Handler: router,
	}

	go func(){
		err:=	server.ListenAndServe()

if err!=nil{
	log.Fatal("Internal Server error")
}

	}()

	<-ch

	slog.Info("shutting down the server")

	ctx,cancel:=context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()
	err=server.Shutdown(ctx)
	if err!=nil{
		slog.Error("FAILED TO SHUTDOWN",slog.String("error",err.Error()))

	}

	slog.Info("ShutDown succesfull")








}