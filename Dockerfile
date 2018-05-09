FROM golang:1.10.2-alpine

RUN apk --no-cache update && apk --no-cache add git build-base

VOLUME /go/src/github.com/VG-Tech-Dojo/vg-1day-2018
WORKDIR /go/src/github.com/VG-Tech-Dojo/vg-1day-2018

ENTRYPOINT [ "make" ]
CMD [ "-C", "original", "run" ]
