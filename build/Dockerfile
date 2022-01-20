# Copyright 2019 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
FROM golang:alpine
WORKDIR /go/src/github.com/go-labs/app/go-labs

ENV GOSUMDB sum.golang.google.cn
ENV GOPROXY http://192.168.2.33:8081/repository/go-group/
ENV GO111MODULE=on
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk --no-cache add git pkgconfig build-base

# Cache go modules
ADD go.mod .
ADD go.sum .
RUN go mod download
ADD . .


#RUN swag init --parseDependency --parseInternal && GO111MODULE=${GO111MODULE} go build cmd/ai_lab.go -o /go/bin/ai_lab
RUN  GO111MODULE=${GO111MODULE} go build -buildmode=pie  -o /go/bin/go-labs cmd/go-labs.go
RUN  strip /go/bin/go-labs


FROM alpine:3.11
RUN apk --no-cache add ca-certificates libdrm
RUN apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata
WORKDIR /app/
COPY --from=0 /go/bin/go-labs .
CMD ["./go-labs"]
