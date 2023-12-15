FROM golang:alpine
WORKDIR /recipe_giver
COPY . .
EXPOSE 8081
RUN go mod download
ENTRYPOINT ["go", "run", "main.go"]
