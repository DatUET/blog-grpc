gen:
	protoc blogpb/blog.proto --go_out=plugins=grpc:.
run-server:
	go run blog_server/server.go
run-client:
	go run blog_client/client.go