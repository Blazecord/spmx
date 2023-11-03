FROM golang:1.21.3

WORKDIR /spmx/

COPY . .

EXPOSE 4000

RUN go install

CMD ["go", "run", "main.go"]