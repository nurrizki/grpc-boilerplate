FROM golang:1.17.5 as builder

# Install Protocol Buffers
RUN apt-get update && \
    apt-get install -y protobuf-compiler

# Install protoc-gen-go and protoc-gen-go-grpc
RUN go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
# Install protoc-gen-go
RUN go get -u github.com/golang/protobuf/protoc-gen-go
# Install protoc-gen-go, protoc-gen-go-grpc, and protoc-gen-govalidators
RUN go get -u github.com/mwitkow/go-proto-validators/protoc-gen-govalidators

ENV PATH $PATH:$GOPATH/bin

RUN mkdir /app
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN make build

WORKDIR /app

# RUN mkdir /app/config \
# 	&& mkdir /app/logs \
	# && mkdir /app/internal \
	# && mkdir /app/internal/config \
	# && mkdir /app/internal/config/db \
	# && mkdir /app/internal/config/server \
	# && mkdir /app/pkg \
	# && mkdir /app/pkg/shared 

# Stage final untuk container runtime
FROM golang:1.17.5 as final

WORKDIR /app

COPY --from=builder /app/main /app
COPY --from=builder /app/internal/config/db/database.yml /app/internal/config/db/
COPY --from=builder /app/internal/config/server/server.yml /app/internal/config/server/

EXPOSE 2202

CMD ["./main"]