package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-chi/chi/v5"
	"github.com/kwantz/ecommerce/services/account/internal/controller"
	"github.com/kwantz/ecommerce/services/account/internal/repository"
	"github.com/kwantz/ecommerce/services/account/internal/usecase"
)

func main() {
	db := setupDatabase()
	defer db.Close()

	accountRepository := repository.NewAccountRepository(db)

	authUsecase := usecase.NewAuthUsecase(accountRepository)
	accountUsecase := usecase.NewAccountUsecase(accountRepository)

	pingController := controller.NewPingController()
	authController := controller.NewAuthController(authUsecase)
	accountController := controller.NewAccountController(accountUsecase)

	router := chi.NewRouter()

	setupRouting(router, Controller{
		Ping:    pingController,
		Auth:    authController,
		Account: accountController,
	})

	log.Println("server starting at :8080")
	http.ListenAndServe(":8080", router)
}

func setupDatabase() *sql.DB {
	db, err := sql.Open("mysql", "user_account:password_account@tcp(mysql-account:3306)/ecommerce_account")
	if err != nil {
		log.Fatal(err.Error())
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}

func setupRouting(router *chi.Mux, controller Controller) {
	router.Get("/ping", controller.Ping.PingHandler)

	router.Post("/login", controller.Auth.AuthenticationHandler)
	router.Post("/authorize", controller.Auth.AuthorizationHandler)

	router.Post("/", controller.Account.CreateAccountHandler)
	router.Get("/", controller.Account.GetAllAccountHandler)
}

type Controller struct {
	Ping    *controller.PingController
	Auth    *controller.AuthController
	Account *controller.AccountController
}
