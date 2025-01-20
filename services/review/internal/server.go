package reviewservice 


import (
	pb "github.com/kasyap1234/book-review/proto"
	"log"
	"net"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartServer2(){
	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalln("Failed to listen to port")

	}
	dsn :="host=localhost user=pg password=pass dbname=crud port=5432 sslmode=disable"


    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect database")
    }

    // Auto migrate the schema
    db.AutoMigrate(&Review{})

    // Create server instance with DB
    server := &server{
        db: db,
    }

	log.Println("server started on port 50052")
	grpcServer := grpc.NewServer()
	pb.RegisterReviewServiceServer(grpcServer,server)
	log.Println("Review service running on port 50052")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("err listening ")
	}

}
