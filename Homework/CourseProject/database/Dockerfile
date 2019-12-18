FROM mysql:latest
COPY create_table.sql /docker-entrypoint-initdb.d
ENV MYSQL_ROOT_PASSWORD=123456
EXPOSE 3306