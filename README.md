docker run --name mypostgres -e POSTGRES_USER=test -e POSTGRES_PASSWORD=password -e POSTGRES_DB=eventline -p 5432:5432 -d -v postgres_data:/var/lib/postgresql/data postgres
