# DOCKER (postgres)

# Create database
docker run --name pg-container -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres

# Create DB with user postgres
docker exec -ti pg-container createdb -U postgres gopgtest

# Enter postgres to change default db to gopgtest
docker exec -ti pg-container -U postgres
\c gopgtest -- (change db)
\q -- (quit)
\dt --(displays the tables)
