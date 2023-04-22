FROM golang:1.20-alpine as deps

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

FROM golang:1.20-alpine as runner

WORKDIR /usr/src/app

ENV CGO_ENABLED=0

COPY --from=deps /usr/src/app /usr/src/app

RUN cd /usr/src/app/cmd

CMD [ "go", "run", "/usr/src/app/cmd/main.go" ]