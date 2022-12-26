package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-chi/chi"
	"github.com/kwantz/ecommerce/services/product/internal/controller"
	"github.com/kwantz/ecommerce/services/product/internal/repository"
	"github.com/kwantz/ecommerce/services/product/internal/usecase"
)

func main() {
	db := setupDatabase()
	defer db.Close()

	productRepository := repository.NewProductRepository(db)

	productUsecase := usecase.NewProductUsecase(productRepository)

	pingController := controller.NewPingController()
	productController := controller.NewProductController(productUsecase)

	router := chi.NewRouter()

	setupRouting(router, Controller{
		Ping:    pingController,
		Product: productController,
	})

	log.Println("server starting at :8080")
	http.ListenAndServe(":8080", router)
}

func setupDatabase() *sql.DB {
	db, err := sql.Open("mysql", "user_product:password_product@tcp(mysql-product:3306)/ecommerce_product")
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

	router.Post("/", controller.Product.CreateProductHandler)
}

type Controller struct {
	Ping    *controller.PingController
	Product *controller.ProductController
}
