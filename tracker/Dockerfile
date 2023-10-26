FROM golang:1.21.3
ENV GOCACHE=/root/.cache/go-build
WORKDIR /go/src/tracker
COPY ./* ./
#RUN go mod download
RUN go build --mod=vendor -o /tracker
EXPOSE 8000
CMD ["/tracker"]
