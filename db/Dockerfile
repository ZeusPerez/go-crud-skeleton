FROM mysql/mysql-server:5.7

EXPOSE 3306

ARG BUILD_TAG=unknown
LABEL BUILD_TAG=$BUILD_TAG

COPY motd.sh /docker-entrypoint-initdb.d/

ENV MYSQL_USER=root \
    MYSQL_PASSWORD=pass \
    MYSQL_DATABASE=test_devs_crud \
    MYSQL_ALLOW_EMPTY_PASSWORD=yes \
    MYSQL_LOG_CONSOLE=true

COPY schema.sql /docker-entrypoint-initdb.d/
