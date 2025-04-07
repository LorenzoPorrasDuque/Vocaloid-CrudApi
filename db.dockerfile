FROM postgres:17.4-alpine3.21

# Set the password for the default PostgreSQL user
ENV POSTGRES_PASSWORD=123

# Create a new database
ENV POSTGRES_DB=test
# Open port 5432
EXPOSE 5432