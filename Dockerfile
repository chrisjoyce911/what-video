FROM chrisjoyce911/go-base AS build

ARG GOPRIVATE
ARG GOPROXY

WORKDIR /workdir

# Let's cache modules retrieval - those don't change so often
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code necessary to build the application
COPY . .
RUN richgo test -cover -v ./...

# Maybe won't pass linter yet
#RUN golangci-lint run ./... || true

RUN golangci-lint run ./...


RUN go build -o /workdir/myapp /workdir/src/.

# Build the application
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /workdir/myapp .

# Build the binary
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /workdir/myapp .

# # Create the minimal runtime image
# FROM scratch as production
# COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
# COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip
# COPY --from=builder /workdir/myapp /myapp

# Set up the app to run as a non-root user inside the /data folder
# User ID 65534 is usually user 'nobody'.
# COPY --chown=65534:0 --from=builder /data /data
# USER 65534
# WORKDIR /data

# COPY /media /media
# COPY /templates /templates
EXPOSE 80
ENTRYPOINT ["/workdir/myapp"]