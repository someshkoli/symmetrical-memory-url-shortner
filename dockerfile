FROM golang:1.16.4-alpine3.13 AS build

WORKDIR /shortner

COPY . .

RUN go build .

from alpine

WORKDIR /shortner
COPY --from=build /shortner/symmetrical-memory-url-shorner .
CMD ["./symmetrical-memory-url-shorner"]

EXPOSE 8888
