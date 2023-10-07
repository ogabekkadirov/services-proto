PWD=$(shell pwd)
PROTO_DIRS=eater \
restaurant_staff \
restaurant_restaurant \
restaurant_payment \
restaurant_support \
logistics_staff \
logistics_vehicle \
logistics_dispatching \
logistics_support

clear:
	rm -rf gen/*

dirs: clear
	mkdir -p gen;
	@for dir in $(PROTO_DIRS); do \
        echo "Creating directory $$dir "; \
		mkdir -p ./gen/$$dir; \
    done

protos:
	@for dir in $(PROTO_DIRS); do \
        echo "Generating $$dir protos"; \
		protoc \
			--go_out=./gen/$$dir \
			--go_opt=paths=source_relative \
    		--go-grpc_out=./gen/$$dir \
			--go-grpc_opt=paths=source_relative \
			-I=$(PWD)/$$dir \
   			$(PWD)/$$dir/*.proto; \
	done