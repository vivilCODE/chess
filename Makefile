protogen: 
	protoc --go_out=./chessapi/pb --go-grpc_out=./chessapi/pb ./proto/chess.proto 
	protoc --js_out=import_style=commonjs,binary:./frontend/src --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./frontend/src ./proto/chess.proto

protoclean: 
	rm ./chessapi/pb/*.go
	rm ./frontend/src/proto/*.js

buildenvoy: 
	docker build -t envoyimage -f ./build/envoy/Dockerfile .

runenvoy: 
	docker run -d -p 8081:8000 envoyimage

pingserver: 
	grpcurl --plaintext localhost:8080 pb.ChessApi/Ping

pingenvoy: 
	grpcurl --plaintext localhost:8081 pb.ChessApi/Ping


