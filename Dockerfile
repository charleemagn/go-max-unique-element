FROM ubuntu:24.04

RUN apt-get update && apt-get install -y golang-go

COPY go.deb /tmp/go.deb

RUN dpkg -i /tmp/go.deb

CMD ["/usr/bin/max_unique_element"]
