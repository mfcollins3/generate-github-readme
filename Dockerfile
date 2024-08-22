# Copyright 2024 Michael F. Collins, III
#
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in
# all copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
# SOFTWARE.

FROM golang:1.23.0-bookworm AS build

WORKDIR /go/src/github.com/mfcollins3/generate-github-readme
COPY go.mod ./
# COPY go.mod go.sum ./
# RUN go mod download && go mod verify

COPY . .
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -v -o /opt/generate-github-readme/bin/generate-readme ./cmd/generate-readme

FROM alpine:3.20.2

COPY --from=build /opt/generate-github-readme/bin/generate-readme /opt/generate-github-readme/bin/generate-readme

ENTRYPOINT ["/opt/generate-github-readme/bin/generate-readme"]
