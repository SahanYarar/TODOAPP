FROM golang:latest
COPY . .
RUN go get -v 
RUN go build -o main .
EXPOSE 9092
CMD ["./main"]