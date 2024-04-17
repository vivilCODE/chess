protogen: 
	protoc --go_out=./chessapi/pb --go-grpc_out=./chessapi/pb ./proto/chess.proto 
	protoc --js_out=import_style=commonjs,binary:./frontend/src --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./frontend/src ./proto/chess.proto

protoclean: 
	rm ./chessapi/pb/*.go
	rm ./frontend/src/proto/*.js

buildenvoy: 
	docker build -t chessapp-envoyimage:latest -f ./build/envoy/Dockerfile .

runenvoy: 
	docker run -d -p 8081:8000 --name chessapp-envoy chessapp-envoyimage

buildfrontend: 
	docker build -t chessapp-frontendimage:latest -f ./build/frontend/Dockerfile .

runfrontend:
	docker run -d -p 8000:3000 --name chessapp-frontend chessapp-frontendimage:latest

buildchessapi: 
	docker build -t chessapp-chessapiimage:latest -f ./build/chessapi/Dockerfile .

runchessapi:
	docker run -d -p 8080:8080 --name chessapp-chessapi chessapp-chessapiimage:latest


pingserver: 
	grpcurl --plaintext localhost:8080 pb.ChessApi/Ping

pingenvoy: 
	grpcurl --plaintext localhost:8081 pb.ChessApi/Ping


	

startserver: 
	go run ./chessapi/server.go

startfrontend: 
	cd ./frontend && npm start


compose-up: 
	docker-compose up