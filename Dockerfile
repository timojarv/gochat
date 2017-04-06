FROM alpine
MAINTAINER Timo Järventausta <timo.jaerv@gmail.com>

EXPOSE 8000

COPY main /
COPY public /

CMD "main"