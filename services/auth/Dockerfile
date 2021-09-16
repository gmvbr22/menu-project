FROM golang:alpine3.14 as builder
LABEL stage=builder
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
RUN go build app/main.go

FROM alpine:latest as production
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
RUN mkdir -p /app
RUN chown -R appuser /app
WORKDIR /app
COPY --from=builder /go/src/app/main .
USER appuser
CMD /app/main