FROM golang:1.21.3
WORKDIR /app
COPY ./* ./
#RUN go mod tidy
RUN go build -o /peer
EXPOSE 9500
CMD ["/peer"]
