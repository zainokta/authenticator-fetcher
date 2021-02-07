FROM postgres:alpine

ENV POSTGRES_DB authenticator_fetcher

ENV POSTGRES_USER postgres

ENV POSTGRES_PASSWORD Master@123

COPY init.sql /docker-entrypoint-initdb.d/

EXPOSE 5432