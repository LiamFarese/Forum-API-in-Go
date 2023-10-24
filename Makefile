postgres:
	docker run --name postgres15.3 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15.3

createdb:
	docker exec -it postgres15.3 createdb --username=root --owner=root synoptic

dropdb:
	docker exec -it postgres15.3 dropdb synoptic

migrateup:
	migrate -path ./db/migrations -database "postgresql://root:secret@localhost:5432/synoptic?sslmode=disable" -verbose up 

migratedown:
	migrate -path ./db/migrations -database "postgresql://root:secret@localhost:5432/synoptic?sslmode=disable" -verbose down

createtestdb:
	docker exec -it postgres15.3 createdb --username=root --owner=root synoptictest

migrateuptest:
	migrate -path ./db/migrations -database "postgresql://root:secret@localhost:5432/synoptictest?sslmode=disable" -verbose up

.PHONY: postgres createdb dropdb migrateup migratedown createtestdb migrateuptest