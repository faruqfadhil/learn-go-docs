package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"database/sql"

	"github.com/faruqfadhil/learn-go-docs/grpc-go-course/blog/blogpb"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct{}

type blogItem struct {
	ID       string
	AuthorID string
	Content  string
	Title    string
}

var db *sql.DB

func (*server) CreateBlog(ctx context.Context, req *blogpb.CreateBlogRequest) (*blogpb.CreateBlogResponse, error) {
	fmt.Println("Create Request")
	blog := req.GetBlog()
	data := blogItem{
		ID:       blog.GetId(),
		AuthorID: blog.GetAuthorId(),
		Content:  blog.GetContent(),
		Title:    blog.GetTitle(),
	}
	stmtInsert, err := db.Prepare("Insert into blog_item(id,author_id,content,title) values(?,?,?,?)")
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal err %v", err),
		)
	}
	_, err = stmtInsert.Exec(data.ID, data.AuthorID, data.Content, data.Title)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal err %v", err),
		)
	}
	return &blogpb.CreateBlogResponse{
		Blog: &blogpb.Blog{
			Id:       data.ID,
			AuthorId: data.AuthorID,
			Content:  data.Content,
			Title:    data.Title,
		},
	}, err
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Connecting to db...")
	var err error
	db, err = sql.Open("mysql", "frq:psswd@/blog_db")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	fmt.Println("Blog service started")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	blogpb.RegisterBlogServiceServer(s, &server{})

	go func() {
		fmt.Println("Starting server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve :  %v", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch
	fmt.Println("Stopping server...")
	s.Stop()
	fmt.Println("Closing listenner...")
	lis.Close()
	defer db.Close()
	fmt.Println("End of program...")
}
