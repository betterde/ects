# This image is used to create bleeding edge docker image and is not compatible with any other image
FROM golang:1.13.0-alpine3.10

# Copy sources
WORKDIR /go/src/ects
COPY . .

# Build development version
ENV BUILD_PLATFORMS -osarch=linux/amd64
RUN make

# Install runner
RUN packaging/root/usr/share/gitlab-runner/post-install

# Preserve runner's data
VOLUME ["/etc/gitlab-runner", "/home/gitlab-runner"]

# init sets up the environment and launches gitlab-runner
CMD ["run", "--user=gitlab-runner", "--working-directory=/home/gitlab-runner"]
ENTRYPOINT ["/usr/bin/gitlab-runner"]