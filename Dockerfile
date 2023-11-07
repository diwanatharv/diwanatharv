#builder stage
FROM golang:1.20-alpine AS builder
ENV APPHOME=/app
WORKDIR $APPHOME
COPY . ./
RUN go mod download && go mod verify && go mod tidy
RUN go build -o /main ./app/service.go

#final stage
FROM alpine:latest
ENV APPHOME=/app
WORKDIR $APPHOME
COPY --from=builder /main ./
COPY ./conf/user-service.yaml ./conf/user-service.yaml
RUN chmod 777 ./main
EXPOSE 6000
CMD ["./main"]