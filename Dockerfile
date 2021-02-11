FROM golang:alpine
WORKDIR /app
ADD main.go /app
ADD go.mod /app 
RUN go build -o clonedemo 
ENTRYPOINT [ "/app/clonedemo" ]
