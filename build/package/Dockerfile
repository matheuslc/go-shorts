FROM golang:latest as builder

LABEL maintainer = "Matheus Carmo <mematheuslc@gmail.com>"

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Final image
FROM alpine:latest  
RUN apk --no-cache add ca-certificates

WORKDIR /go-shorts
COPY --from=builder /app/main .

EXPOSE 3000

# Command to run the executable
CMD ["./main"] 