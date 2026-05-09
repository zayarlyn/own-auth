============================================================
# Stage 1 — dependency cache
# Separate layer so `go mod download` only reruns when go.sum
# changes, not on every source file edit.
# ============================================================
FROM golang:1.23-alpine AS deps

WORKDIR /app

# Install git (needed by some go modules) and ca-certificates
RUN apk add --no-cache git ca-certificates

COPY go.mod go.sum ./
RUN go mod download && go mod verify


# ============================================================
# Stage 2 — builder
# Compiles a fully static binary with no CGO dependencies.
# ============================================================
FROM deps AS builder

WORKDIR /app

# Copy all source
COPY . .

# Build args let CI inject version metadata at build time:
#   docker build --build-arg VERSION=1.2.3 --build-arg COMMIT=abc123 .
ARG VERSION=dev
ARG COMMIT=unknown
ARG BUILD_TIME

RUN CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    go build \
      -trimpath \
      -ldflags="-s -w \
        -X main.Version=${VERSION} \
        -X main.Commit=${COMMIT} \
        -X main.BuildTime=${BUILD_TIME}" \
      -o /bin/own-auth \
      ./cmd/server


# ============================================================
# Stage 3 — final image
# scratch = zero OS, zero shell, zero attack surface.
# The binary is fully static so it needs nothing else.
# ============================================================
FROM scratch AS final

# Pull in TLS root certs so GoAuth can make outbound HTTPS
# calls (e.g. email provider, future OAuth upstream).
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy the binary
COPY --from=builder /bin/own-auth /own-auth

# Copy migration files — the server runs migrations on startup
# unless SKIP_MIGRATIONS=true is set.
COPY --from=builder /app/migrations /migrations

# Non-root numeric UID (scratch has no /etc/passwd).
# This satisfies most container security scanners.
USER 65532:65532

# Expose the HTTP port. The actual port is configured via
# SERVER_PORT env var (default 8080).
EXPOSE 8080

# Health check for Docker and compose.
# /healthz returns 200 {"status":"ok"} with no auth required.
HEALTHCHECK --interval=15s --timeout=3s --start-period=10s --retries=3 \
  CMD ["/own-auth", "healthcheck"]

ENTRYPOINT ["/own-auth"]
