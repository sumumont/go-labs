# note: call scripts from /scripts

image_name =go-labs
harbor_addr=harbor.apulis.cn:8443/aistudio/app/${image_name}
tag        =$(tag)
arch       =$(shell arch)

testarch:
ifeq (${arch}, x86_64)
	@echo "current build host is amd64 ..."
	$(eval arch=amd64)
else ifeq (${arch},aarch64)
	@echo "current build host is arm64 ..."
	$(eval arch=arm64)
else
	echo "cannot judge host arch:${arch}"
	exit -1
endif
	@echo "arch type:$(arch)"





get-deps:
	#git submodule sync
	#git submodule update --init --recursive
	go mod tidy
	go mod download

vet-check-all: get-deps
	go vet ./...

gosec-check-all: get-deps
	gosec ./...

bin: get-deps
	go build -o ${image_name} cmd/${image_name}.go

run: bin
	./ai-arts

docker:
	docker build -f Dockerfile . -t ${image_name}:v0.1.0

gen-swagger:
	swag init -g cmd/${image_name}.go -o api

builder:
	docker build -t ${image_name} -f build/Dockerfile .

push:
	docker tag ${image_name} ${harbor_addr}:${tag}
	docker push ${harbor_addr}:${tag}

dist: testarch
	docker build -t ${image_name} -f build/Dockerfile .
	docker tag ${image_name} ${harbor_addr}/${arch}:${tag}
	docker push ${harbor_addr}/${arch}:${tag}

manifest:
	./docker_manifest.sh ${harbor_addr}:${tag}

localpush:
	docker build -t ${image_name} -f buildlocal/Dockerfile .
	docker tag ${image_name} ${harbor_addr}:${tag}
	docker push ${harbor_addr}:${tag}

localbuild: get-deps bin

dkpush: localbuild localpush