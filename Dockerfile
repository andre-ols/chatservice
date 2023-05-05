FROM golang:1.20.3

WORKDIR /go/src
RUN ln -sf /bin/bash /bin/sh
COPY . .

CMD [ "tail", "-f", "/dev/null" ]