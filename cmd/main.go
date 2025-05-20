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
	"go/adv-api/pkg/event"
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
func App() http.Handler {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()
	eventBus := event.NewEventBus()

	// Repository
	linkRepository := link.NewLinkRepository(db)
	userRepository := user.NewUserRepository(db)
	statRepositiry := stat.NewStatRepository(db)
	//Services
	authService := auth.NewAuthService(userRepository)
	statService := stat.NewStatService(&stat.StatServiceDeps{
		EventBus:       eventBus,
		StatRepository: statRepositiry,
	})

	//Handler
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
		Config:         conf,
		EventBus:       eventBus})
	stat.NewStatHandler(router, stat.StatHandlerDeps{
		StatRepository: statRepositiry,
		Config:         conf})
	go statService.AddClick()
	//Middleware
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Looging)
	return stack(router)
}
func main() {
	app := App()
	server := http.Server{
		Addr:    ":8081",
		Handler: app}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
