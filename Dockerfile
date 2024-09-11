FROM golang:1.23

COPY . .
RUN go mod download 
RUN apt update && apt install -y dnsutils iputils-ping net-tools
EXPOSE 8080

CMD ["go", "run", "main.go"]