# Project description

Over-engineered calculator:

Must have:

1. Create a basic calculator with history.
2. There should be a clear separation of logic.
3. All logic should be covered by unit tests.
4. The calculator should expose a RESTful interface deployed via a backend technology fx Firebase, but a docker container on eg. Heroku will also suffice or something else you are familar with.
5. The interface should be documented with an accompanying Postman package for easy testing.
6. All source code should be available via your personal Github account.

Nice to have:

1. Best effort micro-service architecture (given the timeframe).
2. Auth for using the service (email/password will suffice).
3. A small webpage utilizing the calculator.

As the title says, it is an over-engineered calculator, which means:
You should go all-in on design patterns and best practices as well as making sure to fulfill the "must have" requirements.
The coding language is also up to you, but GO lang would be preferred, since it is part of our tech stack.

# Create database

https://medium.com/@roystatham3003/database-connection-golang-docker-dfff9e958e47
https://cloud.google.com/sql/docs/postgres/connect-run#public-ip-default

docker run --name pg -e POSTGRES_PASSWORD=lunar -p 5432:5432 -d postgres
# postgres will probably already exist
docker exec -ti pg createdb -U postgres postgres
docker exec -ti pg psql -U postgres

swagger generate server --spec api/swagger.yaml --name calculatron-service --exclude-main --target ./generated