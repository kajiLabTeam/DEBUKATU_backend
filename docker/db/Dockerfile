FROM mysql:8.0

COPY ./db.cnf /etc/mysql/conf.d/my.cnf

# 初期化SQL
COPY ./init.sql /docker-entrypoint-initdb.d/
