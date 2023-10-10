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

RUN make build_monitor


FROM gcr.io/distroless/cc-debian11:latest

ENV DATA_MARKETPLACE_HOME /opt/app
ENV CONFIG_FILE_PATH $DATA_MARKETPLACE_HOME/config/config.json

ENV WORKDIR=/app
WORKDIR ${WORKDIR}
COPY --from=builder /opt/app/build/monitor ${WORKDIR}

# Run the app
ENTRYPOINT ["/app/monitor", "--config-path",  "$CONFIG_FILE_PATH"]
