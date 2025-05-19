# Create database

https://medium.com/@roystatham3003/database-connection-golang-docker-dfff9e958e47
docker run --name pg-demo -e POSTGRES_PASSWORD=password -p 5432:5432 -d postgres
docker exec -ti pg-demo createdb -U postgres calculatron
docker exec -ti go-postgres-demo psql -U postgres