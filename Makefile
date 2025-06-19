include vendor.proto.mk

# Используем bin в текущей директории для установки плагинов protoc
LOCAL_BIN := $(CURDIR)/bin

# Добавляем bin в текущей директории в PATH при запуске protoc
PROTOC = PATH="$$PATH:$(LOCAL_BIN)" protoc

# Путь до protobuf файлов
PROTO_SRC := $(CURDIR)/protobuf/api

# Путь до сгенеренных .pb.go файлов
PROTO_OUT := $(CURDIR)/pkg

# Путь до завендореных protobuf файлов
VENDOR_PROTO := $(CURDIR)/vendor.protobuf

# устанавливаем необходимые плагины
.bin-deps: export GOBIN := $(LOCAL_BIN)
.bin-deps:
	$(info Installing binary dependencies...)

	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/envoyproxy/protoc-gen-validate@latest

# генерация .go файлов с помощью protoc
.protoc-generate:
	$(info Generate .go files from protobuf...)

	mkdir -p $(PROTO_OUT)
	$(PROTOC) -I=$(CURDIR) \
	-I=$(CURDIR)/vendor.protobuf \
	-I=$(CURDIR)/vendor.protobuf/envoyproxy-validate \
 	  --proto_path=$(PROTO_OUT) \
	  --go_out=$(PROTO_OUT) --go_opt paths=source_relative \
	  --go-grpc_out=$(PROTO_OUT) --go-grpc_opt paths=source_relative \
	  --grpc-gateway_out=$(PROTO_OUT) --grpc-gateway_opt paths=source_relative \
	  --validate_out="lang=go,paths=source_relative:$(PROTO_OUT)" \
	  $(PROTO_SRC)/salary/messages.proto \
	  $(PROTO_SRC)/salary/service.proto \
	  $(PROTO_SRC)/user/message.proto \
	  $(PROTO_SRC)/user/service.proto\
	  $(PROTO_SRC)/vacation/messages.proto \
	  $(PROTO_SRC)/vacation/service.proto \
	  $(PROTO_SRC)/salary/v2/messages.proto \
	  $(PROTO_SRC)/salary/v2/service.proto

# генерация .go файлов с помощью protoc
generate: .bin-deps .protoc-generate

# Генерация Swagger (OpenAPIv2)
.swagger-generate:
	$(info Generate OpenAPIv2 swagger json...)
	$(PROTOC) \
	  -I=$(CURDIR) -I=$(VENDOR_PROTO) \
	  --openapiv2_out=. --openapiv2_opt logtostderr=true \
	  $(PROTO_SRC)/user/service.proto

swagger: .swagger-generate

.PHONY: vendor .protoc-generate generate .bin-deps .swagger-generate swagger