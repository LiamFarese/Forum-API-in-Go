pullpostgres:
	docker pull postgres:15.3

postgres:
	docker run --name postgres15.3 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15.3

createdb:
	docker exec -it postgres15.3 createdb --username=root --owner=root forum

dropdb:
	docker exec -it postgres15.3 dropdb forum

migrateup:
	migrate -path ./db/migrations -database "postgresql://root:secret@localhost:5432/forum?sslmode=disable" -verbose up 

migratedown:
	migrate -path ./db/migrations -database "postgresql://root:secret@localhost:5432/forum?sslmode=disable" -verbose down

createtestdb:
	docker exec -it postgres15.3 createdb --username=root --owner=root forumtest

migrateuptest:
	migrate -path ./db/migrations -database "postgresql://root:secret@localhost:5432/forumtest?sslmode=disable" -verbose up

.PHONY: pullpostgres postgres createdb dropdb migrateup migratedown createtestdb migrateuptest
