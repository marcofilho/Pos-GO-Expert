package main

import (
	"database/sql"
	"net"

	"github.com/marcofilho/Pos-GO-Expert/gRPC/internal/database"
	"github.com/marcofilho/Pos-GO-Expert/gRPC/internal/pb"
	"github.com/marcofilho/Pos-GO-Expert/gRPC/internal/service"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS categories (
        id TEXT,
        name TEXT,
        description TEXT
    )`)
	if err != nil {
		panic(err)
	}

	categoryDB := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDB)

	grpcService := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcService, categoryService)
	reflection.Register(grpcService)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcService.Serve(listener); err != nil {
		panic(err)
	}
}
