package main

import (
	"context"
	"datuet/learn-grpc/blog_mongodb/blogpb"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Blog Client")

	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)

	DeleteBlog(c)
}

func CRUBlog(c blogpb.BlogServiceClient) {
	// create blog
	fmt.Println("Create the blog")
	blog := &blogpb.Blog{
		AuthorId: "DatUET",
		Title:    "Test Read Blog",
		Content:  "Content of blog",
	}
	createBlogReq, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	fmt.Printf("Blog has been created: %v", createBlogReq)
	blogID := createBlogReq.GetBlog().GetId()

	// read blog
	fmt.Println("Reading the blog")
	_, err2 := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: "602ab68271583d2eef0f82c"})
	if err2 != nil {
		fmt.Printf("Error happened while reading %v", err2)
	}

	readBlogReq := &blogpb.ReadBlogRequest{
		BlogId: blogID,
	}
	readBlogRes, readBlogErr := c.ReadBlog(context.Background(), readBlogReq)
	if readBlogErr != nil {
		fmt.Printf("Error happened while reading %v\n", readBlogErr)
	}
	fmt.Printf("Blog was read %v", readBlogRes)

	// update blog
	newBlog := &blogpb.Blog{
		Id:       blogID,
		AuthorId: "Author update",
		Title:    "Title Update",
		Content:  "Content Update",
	}
	updateRes, updateErr := c.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{Blog: newBlog})
	if updateErr != nil {
		fmt.Printf("Error while update %v \n", updateErr)
	}
	fmt.Printf("Blog was update: %v \n", updateRes)
}

func DeleteBlog(c blogpb.BlogServiceClient) {
	deleteRes, err := c.DeleteBlog(context.Background(), &blogpb.DeleteBlogRequest{BlogId: "602ab68271583d2eef0f82c7"})
	if err != nil {
		fmt.Printf("Error while deleting %v", err)
	}
	fmt.Printf("Blog was delete: %v", deleteRes.GetBlogId())
}