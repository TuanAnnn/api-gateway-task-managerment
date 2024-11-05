FROM golang:1.23

WORKDIR /go/src/task_management_api_gatway

# Copy go.mod và go.sum trước để cache các dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy toàn bộ mã nguồn và build
COPY . .
RUN go build -o bin/server cmd/main.go

CMD ["./bin/server"]
