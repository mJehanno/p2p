FROM golang:1.21.3
ENV GOFLAGS=-mod=vendor
ENV PORT=9500
WORKDIR /app
COPY . .
#RUN go mod vendor
RUN go build -o /peer
EXPOSE 9500
CMD ["/peer"]
