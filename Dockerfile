
FROM golang:1.15.11-alpine3.13 AS build

# Set necessary environmet variables needed for our image
# ENV GO111MODULE=on \
#     CGO_ENABLED=0 \
#     GOOS=linux \
#     GOARCH=amd64

WORKDIR /build

COPY src .

RUN ls -lha 

RUN go version

RUN go build -o /bin/my-web-app


FROM alpine:3.13
COPY --from=build /bin/my-web-app  /bin/my-web-app
ENTRYPOINT ["/bin/my-web-app"]