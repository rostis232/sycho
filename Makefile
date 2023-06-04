makemigration:
	migrate create -ext sql -dir ./schema -seq init

up:
	migrate -path ./schema -database 'postgres://postgres:psychopass2023@localhost:5436/postgres?sslmode=disable' up

down:
	migrate -path ./schema -database 'postgres://postgres:psychopass2023@localhost:5436/postgres?sslmode=disable' down

startdb:
	sudo docker start pssdb