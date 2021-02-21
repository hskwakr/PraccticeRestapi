FROM mysql:latest
EXPOSE 3306

ADD ./my.cnf /etc/mysql/conf.d/my.cnf
ADD ./sql/ /docker-entrypoint-initdb.d

CMD ["mysqld"]
