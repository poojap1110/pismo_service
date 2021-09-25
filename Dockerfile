# Base Docker Image: golang-builder
FROM 325140741718.dkr.ecr.ap-south-1.amazonaws.com/golang:alpine as builder

# NETRC Login User
ARG NETRC_LOGIN

# NETRC Password User
ARG NETRC_PASSWORD

ENV APP_NAME integration-svc-aub
ENV GOPRIVATE bitbucket.org/matchmove/*

ENV APP_EXTERNAL_DOMAIN http://aub-svc.kops.matchmove-beta.com:8080

# Make sure the Code is pulled
ENV /go/src/bitbucket.org/matchmove/{$APP_NAME} "bitbucket.org/matchmove/*"
ENV CODE_DIR /go/src/bitbucket.org/matchmove/{$APP_NAME}

ENV SWAGGER_PATH /docs/

RUN mkdir -p $SWAGGER_PATH

# Make code directory as working directory
WORKDIR ${CODE_DIR}

# COPY go.mod file
ADD go.mod .

# # COPY go.sum file
ADD go.sum .

RUN apk add alpine-sdk && \
    echo "machine bitbucket.org" >> /root/.netrc && \
    echo "login $NETRC_LOGIN" >> /root/.netrc && \
    echo "password $NETRC_PASSWORD" >> /root/.netrc && \
    chmod 600 /root/.netrc && \
    go mod download
# Copy Code
COPY . .

RUN sh swagger.sh

COPY /docs/* /docs/

# Build Code
RUN rm -rf /root/.netrc
RUN go build -o /bin/artifact -mod mod

# #### End of Builder ####

# Base Docker Image: golang
FROM 325140741718.dkr.ecr.ap-south-1.amazonaws.com/golang:alpine

ENV SWAGGER_PATH /docs/aub/

RUN mkdir -p $SWAGGER_PATH

# Copy Artifact
COPY --from=builder /bin/artifact /bin/

COPY --from=builder /docs $SWAGGER_PATH

# Create headless user
RUN addgroup -S littlefinger && adduser -S littlefinger -G littlefinger && \
    # give permissions to user
    chmod +x /bin/artifact

# No Privilege User
USER littlefinger

# setting emtpy as the entrypoint
ENTRYPOINT []

# setting command
CMD ["/bin/artifact"]
