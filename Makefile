restart:
	sudo docker rm -f FM
	sudo docker run --name FM -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=postgres -p 5432:5432 -d postgres:latest


down:
	sudo docker stop postgres

prod:
	./scripts/run.sh prod

dev:
	./scripts/run.sh dev

exec:
	sudo docker exec -it FM psql -U postgres

push:
	git add .
	git commit -m "$(m)"
	git push
	./scripts/push.sh
build:
	sudo docker push mujag/fm-fiber:prod 

