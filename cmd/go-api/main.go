package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/ayushmehta03/go-api/internal/config"
)




func main(){

	cfg:=config.MustLoad()


	router:=http.NewServeMux();
	router.HandleFunc("GET /",func(w http.ResponseWriter, r *http.Request){
		w.Write([] byte("Home Page"))
	})
	
	ch:=make(chan os.Signal,1)

	signal.Notify(ch,os.Interrupt,syscall.SIGINT,syscall.SIGTERM)

	server:=http.Server{
		Addr: cfg.Addr,
		Handler: router,
	}

	go func(){
		err:=	server.ListenAndServe()

if err!=nil{
	log.Fatal("nternal Server error")
}

	}()

	<-ch








}