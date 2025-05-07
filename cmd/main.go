package main

import (
	"context"
	"fmt"
	"go/adv-api/configs"
	"go/adv-api/internal/auth"
	"go/adv-api/internal/link"
	"go/adv-api/internal/stat"
	"go/adv-api/internal/user"
	"go/adv-api/pkg/db"
	"go/adv-api/pkg/middleware"
	"net/http"
	"time"
)

func tickOperation(ctx context.Context) {
	tikcer := time.NewTicker(200 * time.Millisecond)
	for {
		select {
		case <-tikcer.C:
			fmt.Println("Tick")
		case <-ctx.Done():
			fmt.Println("Cancel")
			return
		}
	}
}
func main2() {

	ctx, cancel := context.WithCancel(context.Background())
	go tickOperation(ctx)
	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}
func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()

	// Repository
	linkRepository := link.NewLinkRepository(db)
	userRepository := user.NewUserRepository(db)
	statRepositiry := stat.NewStatRepository(db)
	//Services
	authService := auth.NewAuthService(userRepository)
	//Handler
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{Config: conf,
		 AuthService: authService})
	link.NewLinkHandler(router, link.LinkHandlerDeps{LinkRepository: linkRepository,
		 Config: conf,
		  StatRepository: statRepositiry})
	//Middleware
	stack := middleware.Chain(middleware.CORS, middleware.Looging)
	server := http.Server{
		Addr:    ":8081",
		Handler: stack(router)}
	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
