FROM ubuntu:16.04
LABEL owner="Jason Lee"

RUN apt-get update -y && apt-get install -y curl && apt-get clean && rm -rf /var/lib/apt/lists && mkdir /opt/test-server
ADD test-server /opt/test-server/test-server
WORKDIR /opt/test-server
CMD ./test-server

HEALTHCHECK --interval=5s --timeout=10s CMD curl -f http://localhost/health || exit 1