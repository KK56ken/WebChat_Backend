goContainerName=go
DBContainerName=mysql_host


build:
	docker-compose build

up:
	docker-compose up

up/d:
	docker-compose up -d

ps:
	docker-compose ps

stop:
	docker stop $(goContainerName)
	docker stop  $(DBContainerName)

stop/go:
	docker stop $(goContainerName)
stop/db:
	docker stop $(goContainerName)

go:
	docker exec -it $(goContainerName) sh
go/bash:
	docker exec -it $(goContainerName) bin/bash
sql:
	docker exec -it $(DBContainerName) sh -c 'mysql -u $$MYSQL_USER -p$$MYSQL_PASSWORD $$MYSQL_DATABASE'
sql/bash:
	docker exec -it $(DBContainerName) bin/bash
