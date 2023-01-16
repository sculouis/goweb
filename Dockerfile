#build stage
#final stage
FROM golang
RUN mkdir -p app
WORKDIR /app
COPY pojo/ /app/pojo
COPY service/ /app/service
COPY src/ /app/src
COPY . .
RUN go mod download
RUN go build -o app
ENTRYPOINT ["./app"]
LABEL Name=goweb Version=0.0.1
EXPOSE 8080