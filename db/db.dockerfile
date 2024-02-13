FROM postgres:latest

# Set the password for the default PostgreSQL user
ENV POSTGRES_PASSWORD=123

# Create a new database
ENV POSTGRES_DB=test
# Open port 5432

EXPOSE 5432

#create table cosa
COPY init.sql /docker-entrypoint-initdb.d/


