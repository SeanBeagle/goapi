FROM golang:1.16

RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go build -o server . 

CMD [ "/app/server" ]