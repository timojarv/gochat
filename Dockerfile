FROM scratch
MAINTAINER Timo Järventausta <timo.jaerv@gmail.com>

EXPOSE 8000

COPY main /
COPY public /public/

CMD ["/main"]
