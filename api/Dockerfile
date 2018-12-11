FROM golang AS builder

RUN mkdir -p /go/src/github.com/learning/project/api
WORKDIR /go/src/github.com/learning/project/api

COPY ./ ./
RUN make install
RUN go build -o bin/api main.go

FROM nginx
COPY --from=builder /go/src/github.com/learning/project/api/bin /app
CMD ["/app/api"]