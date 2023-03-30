mysql:
	docker run --name mysqlimg --network emall-network -p 6633:3306 -e MYSQL_ROOT_PASSWORD=238600aA

createdb:
	docker exec -it mysqlimg mysql -uroot -p238600aA -e "create database `simple_bank`;"

dropdb:
	docker exec -it mysqlimg mysql -uroot -p238600aA -e "drop database `simple_bank`;"

migrateup:
	migrate -path db/migration -database "mysql://root:238600aA@tcp(127.0.0.1:6633)/simple_bank" -verbose up

migrateup1:
	migrate -path db/migration -database "mysql://root:238600aA@tcp(127.0.0.1:6633)/simple_bank" -verbose up 1

migratedown:
	migrate -path db/migration -database "mysql://root:238600aA@tcp(127.0.0.1:6633)/simple_bank" -verbose down

migratedown1:
	migrate -path db/migration -database "mysql://root:238600aA@tcp(127.0.0.1:6633)/simple_bank" -verbose down 1

sqlc:
	sqlc generate

server:
	go run main.go

proto:
	@protoc --proto_path=. --micro_out=. --go_out=:. proto/hipstershop.proto
	@protoc --proto_path=. --micro_out=. --go_out=:. proto/health.proto

consul:
# 	docker run --name consul1 -p 8500:8500 -p 8300:8300 -p 8301:8301 -p 8302:8302  -d consul:latest

.PHONY: mysql createdb dropdb migratedown migrateup migrateup1 sqlc server test mock migratedown1 proto