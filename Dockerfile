FROM scratch
MAINTAINER Timo Järventausta <timo.jaerv@gmail.com>

EXPOSE 8000

COPY gochat /
COPY public /public/

CMD ["/gochat"]
