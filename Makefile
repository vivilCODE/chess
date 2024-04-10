gen: 
	protoc --go_out=./pb --go-grpc_out=./pb ./proto/chess.proto

clean: 
	rm pb/*.go

