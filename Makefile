restart:
	sudo docker rm -f FM
	sudo docker run --name FM -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=postgres -p 5432:5432 -d postgres:latest

start:
	sudo docker start FM

prod:
	go run cmd/main.go

down:
	sudo docker stop postgres

setupdev:
	./scripts/setup.sh dev

setupprod:
	./scripts/setup.sh prod

ipfs:
	sudo docker run -d --name ipfs_host -v $ipfs_staging:/export -v $ipfs_data:/data/ipfs -p 4001:4001 -p 4001:4001/udp -p 127.0.0.1:8080:8080 -p 127.0.0.1:5001:5001 ipfs/kubo:latest

run:
	./scripts/run.sh

exec:
	sudo docker exec -it FM psql -U postgres

push:
	git add .
	git commit -m "$(m)"
	git push

build:
	go build