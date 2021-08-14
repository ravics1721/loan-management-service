FROM golang:latest
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod download
ENV key=value
RUN go build -o main .
CMD [ "/app/main" ]
# FROM golang:1.12.0-alpine3.9
# WORKDIR /app
# # COPY go.mod .
# # COPY go.sum .
# RUN apk update
# RUN apk add git
# RUN go mod download
# COPY . .
# ENV PORT=8080
# RUN go build
# CMD [ "./loan-management-service" ]