package bookservice 

import (
	pb "github.com/kasyap1234/book-review/proto"

	"log"
	"net"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartServer1() {
	listener, err := net.Listen("tcp", ":50051")
if err !=nil {
	log.Fatalln("unable to listen to port"); 

}
 // Use environment variable or fallback to local development settings



dsn :="host=localhost user=pg password=pass dbname=crud port=5432 sslmode=disable"


    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect database")
    }

    // Auto migrate the schema
    db.AutoMigrate(&Book{})

    // Create server instance with DB
    server := &server{
        db: db,
    }

log.Println("server started on port :50051")


grpcServer :=grpc.NewServer()
pb.RegisterBookServiceServer(grpcServer,server)
if err :=grpcServer.Serve(listener); err !=nil {
	log.Fatalf("Failed to listen to port"); 
}

}
