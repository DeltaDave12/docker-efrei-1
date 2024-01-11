FROM debian:bullseye-slim

RUN apt-get update && \
    apt-get install -y mysql-server

EXPOSE 3306

CMD ["mysqld", "--user=mysql", "--console"]```