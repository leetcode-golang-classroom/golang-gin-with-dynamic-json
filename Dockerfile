FROM golang:1.22.4 AS build
WORKDIR /app
COPY go.mod go.sum Makefile ./
ADD . ./
RUN go mod download && make build

FROM build AS test
RUN make test

FROM alpine as release
WORKDIR /app
COPY --from=build /app/bin ./bin
COPY --from=build /app/template ./template
CMD ["./bin/main"]