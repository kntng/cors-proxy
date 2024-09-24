FROM golang:alpine AS build
WORKDIR /app
COPY . .
RUN go build -ldflags "-s -w" -o /bin/main ./main.go
RUN apk add --no-cache upx
RUN upx /bin/main

FROM scratch
COPY --from=build /bin/main /bin/main
CMD ["/bin/main"]
