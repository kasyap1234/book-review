package userservice



import (
	pb "github.com/kasyap1234/book-review/proto"
	"log"
	"net"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	
	"gorm.io/gorm"
)


func StartServer3() {
	 // Initialize database connection first
     dsn := "host=database user=pg password=pass dbname=crud port=5432 sslmode=disable"



    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect database")
    }

    // Auto migrate the schema
    db.AutoMigrate(&User{})

    // Create server instance with DB
    server := &Server{
        db: db,
    }
    s := grpc.NewServer()
    listener, err := net.Listen("tcp", ":50053")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

   
    pb.RegisterUserServiceServer(s,server)

    log.Printf("gRPC server starting on port %d", 50053)
    if err := s.Serve(listener); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
