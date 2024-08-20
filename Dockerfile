FROM golang:1.23.0-bookworm AS base
RUN apt-get update && \
  apt-get upgrade -y && \
  apt-get install ca-certificates -y && \
  update-ca-certificates
WORKDIR /app

FROM base AS dev
# Uncomment this if SSH is used to install dependencies
# RUN git config --global --add url."git@github.com:".insteadOf "https://github.com/"
# RUN mkdir -p -m 0700 ~/.ssh && ssh-keyscan github.com >> ~/.ssh/known_hosts \
#   --mount=type=ssh ssh-add -l \
#   --mount=type=ssh go mod download
COPY go.mod .
COPY go.sum .
COPY . .
RUN go generate ./...
RUN go install github.com/cespare/reflex@latest
CMD reflex -r '\.go$$' -s -- go run main.go
EXPOSE 3000

FROM dev AS builder
RUN go build main.go

FROM base AS runner
COPY --from=builder /app/main /app
EXPOSE 3000
ENV PORT 3000
CMD /app/main
