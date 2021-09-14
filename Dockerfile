FROM golang:1.17-alpine3.14 as build
COPY . /go/src/github.com/chesnovsky/fliboobstier
WORKDIR /go/src/github.com/chesnovsky/fliboobstier
RUN apk add --no-cache git make gcc libc-dev ca-certificates sqlite \
  && make db \
  && make deps \
  && make

FROM library/alpine:3.14
RUN apk add --no-cache ca-certificates
COPY config.yml /config.yml
COPY --from=build /go/src/github.com/chesnovsky/fliboobstier/bin/fliboobstier /fliboobstier
COPY --from=build /go/src/github.com/chesnovsky/fliboobstier/bin/fliboobstier.db /fliboobstier.db
CMD ["/fliboobstier"]
