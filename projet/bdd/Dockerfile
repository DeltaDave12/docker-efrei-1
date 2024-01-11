FROM debian:bullseye

WORKDIR /bdd

COPY . .

RUN apt update  

RUN apt install -y mariadb-server

EXPOSE 3306

CMD ["mysqld", "--user=mysql", "--console"]```