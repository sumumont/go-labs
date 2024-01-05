# note: call scripts from /scripts
image_name =cambricon-pytorch16-python375-ubuntu1804-sdk1-120
harbor_addr=harbor.internal.cn:8443/publicharbor_apulis/${image_name}
tag        =v2-orion-0.1

localpush:
	docker build -t ${image_name} -f ./Dockerfile .
	docker tag ${image_name} ${harbor_addr}:${tag}
	docker push ${harbor_addr}:${tag}