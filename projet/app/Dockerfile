FROM debian:bullseye-slim

RUN apt-get update && \
    apt-get install -y apache2

EXPOSE 80

CMD ["apache2ctl", "-D", "FOREGROUND"]```