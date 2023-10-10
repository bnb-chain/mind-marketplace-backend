FROM balenalib/amd64-debian-golang:latest as builder

# Set up apk dependencies
ENV PACKAGES libc6-dev make git  bash build-essential curl

# Set working directory for the build
WORKDIR /opt/app

# Add source files
COPY . .

# Install minimum necessary dependencies, remove packages
RUN apt-get -y update && apt-get -y install $PACKAGES

# For Private REPO
ARG GH_TOKEN=""
RUN go env -w GOPRIVATE="github.com/bnb-chain/*"
RUN git config --global url."https://${GH_TOKEN}@github.com".insteadOf "https://github.com"

RUN make build_server


FROM alpine:3.17

ENV CONFIG_FILE_PATH /opt/app/config/config.json

ENV WORKDIR=/app
WORKDIR ${WORKDIR}
COPY --from=builder /opt/app/build/server ${WORKDIR}

# Run the app
ENTRYPOINT /app/server --port, 8080 --config-path "$CONFIG_FILE_PATH"
