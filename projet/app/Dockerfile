FROM debian:bullseye

WORKDIR /app

COPY . .

COPY index.php /var/www/html/index.php

RUN apt update 

RUN apt install -y apache2 php libapache2-mod-php php-mysql

EXPOSE 80

CMD ["apache2ctl", "-D", "FOREGROUND"]```