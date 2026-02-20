FROM golang AS build

WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server

FROM scratch

WORKDIR /app
COPY --from=build /build/server .
COPY templates.json .
COPY config ./config
COPY public ./public

ENTRYPOINT ["./server"]