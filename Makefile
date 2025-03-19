DB_CONTAINER_NAME=mypostgres

run-app:
	go run cmd/api/main.go

# Comando para iniciar el contenedor de PostgreSQL
start-db:
	@if [ "`docker inspect -f {{.State.Status}} $(DB_CONTAINER_NAME)`" = "exited" ]; then \
    	docker start $(DB_CONTAINER_NAME); \
	else \
    	echo "El contenedor $(DB_CONTAINER_NAME) ya está en ejecución"; \
	fi

# Comando para iniciar la base de datos y ejecutar la aplicación
start:
	make start-db
	sleep 2
	make run-app

	
	