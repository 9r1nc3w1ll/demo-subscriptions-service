proto_generate:
	mkdir -p pb && protoc --proto_path=proto --go_out=:. --go-grpc_out=. proto/*.proto

proto_clean:
	rm pb/*.pb.go

run:
	reflex -r '\.go$$' -s -- go run main.go