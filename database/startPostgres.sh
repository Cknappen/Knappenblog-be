docker pull postgres:14.5
docker run --name knappenblog-pg -p 5432:5432 -e POSTGRES_PASSWORD=mysecretpassword --rm -d postgres
