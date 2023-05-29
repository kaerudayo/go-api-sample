FROM redis:7.0

COPY redis.conf /usr/local/etc/redis/redis.conf

RUN mkdir /var/lib/redis && chown redis:redis /var/lib/redis/

CMD [ "redis-server", "/usr/local/etc/redis/redis.conf" ]
