FROM golang:1.21.3
ENV GOCACHE=/root/.cache/go-build
WORKDIR /go/src/tracker
COPY ./* ./
RUN go mod vendor
RUN go build  -o /tracker
EXPOSE 8000
CMD ["/tracker"]
