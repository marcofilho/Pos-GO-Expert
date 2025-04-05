package main

import (
	"database/sql"
	"fmt"
	"net"
	"net/http"

	graphql_handler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/marcofilho/Pos-GO-Expert/CleanArchitecture/configs"
	"github.com/marcofilho/Pos-GO-Expert/CleanArchitecture/internal/event/handler"
	"github.com/marcofilho/Pos-GO-Expert/CleanArchitecture/internal/infra/gRPC/pb"
	"github.com/marcofilho/Pos-GO-Expert/CleanArchitecture/internal/infra/gRPC/service"
	"github.com/marcofilho/Pos-GO-Expert/CleanArchitecture/internal/infra/graphQL/graph"
	"github.com/marcofilho/Pos-GO-Expert/CleanArchitecture/internal/web/webserver"
	"github.com/marcofilho/Pos-GO-Expert/CleanArchitecture/pkg/events"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS orders (
		id varchar(255) NOT NULL,
		price float NOT NULL, 
		tax float NOT NULL, 
		final_price float NOT NULL,
		PRIMARY KEY (id))`)
	defer db.Close()

	rabbitMQChannel := getRabbitMQChannel()

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("OrderCreated", &handler.OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})

	createOrderUseCase := NewCreateOrderUseCase(db, eventDispatcher)
	getOrderByIdUseCase := NewGetOrderByIdUseCase(db)
	getOrdersUseCase := NewGetOrdersUseCase(db)

	webServer := webserver.NewWebServer(configs.WebServerPort)

	webOrderHandler := NewWebOrderHandler(db, eventDispatcher)

	webServer.AddHandler("/order", webOrderHandler.Create)
	webServer.AddHandler("/orders/{id}", webOrderHandler.GetOrderById)
	webServer.AddHandler("/orders", webOrderHandler.GetOrders)

	fmt.Println("Starting web server on port", configs.WebServerPort)
	go webServer.Start()

	grpcServer := grpc.NewServer()
	createOrderService := service.NewOrderService(*createOrderUseCase, *getOrderByIdUseCase, *getOrdersUseCase)
	pb.RegisterOrderServiceServer(grpcServer, createOrderService)

	reflection.Register(grpcServer)
	fmt.Println("Starting gRPC server on port", configs.GRPCServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", configs.GRPCServerPort))
	if err != nil {
		panic(err)
	}
	go grpcServer.Serve(lis)

	srv := graphql_handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CreateOrderUseCase:  *createOrderUseCase,
		GetOrderByIdUseCase: *getOrderByIdUseCase,
		GetOrdersUseCase:    *getOrdersUseCase,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	fmt.Println("Starting GraphQL server on port", configs.GraphQLServerPort)
	http.ListenAndServe(":"+configs.GraphQLServerPort, nil)
}

func getRabbitMQChannel() *amqp.Channel {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	return ch
}
