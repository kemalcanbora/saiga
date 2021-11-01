FROM golang:1.16-alpine


RUN mkdir app
COPY . /app

WORKDIR /app
ENV PORT=8080

RUN go get


RUN go build
CMD ["./saiga"]