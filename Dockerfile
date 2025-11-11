FROM golang:1.22
WORKDIR /app
COPY . .
RUN go build -o orchestrator
CMD ["./orchestrator"]
EXPOSE 8080
