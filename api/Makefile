SWAGGER_PORT=8088

define gen_proto
	@echo "Regenerate $(1) GO client from .proto"
	@mkdir -p langs/go/$(1)
	@cd langs/go/$(1) && rm -rf *
	@protoc -I. -I.. grpc/$1.proto \
		--go_out=./langs/go/$(1) --go_opt=paths=source_relative \
		--go-grpc_out=./langs/go/$(1) --go-grpc_opt=paths=source_relative \
		--go-grpc_opt=require_unimplemented_servers=false
	@echo "Regenerate $(1) GO client completed"
endef

generate: gen-spaceship

gen-spaceship:
	$(call gen_proto,spaceship)
