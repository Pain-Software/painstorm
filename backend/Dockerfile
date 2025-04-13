FROM golang:1.23-bullseye as base

RUN adduser \
  --disabled-password \
  --gecos "" \
  --home "/nonexistent" \
  --shell "/sbin/nologin" \
  --no-create-home \
  --uid 65532 \
  small-user

COPY .env /app/.env     

WORKDIR /go/src/go_project/

COPY backend/. .           

RUN go mod download
RUN go mod verify

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /main .

FROM gcr.io/distroless/static-debian11

COPY --from=base /etc/group /etc/group
COPY --from=base /etc/passwd /etc/passwd

COPY --from=base /main .
COPY --from=base /go/src/go_project/config.yaml .

USER small-user:small-user

ENTRYPOINT ["/main"]
