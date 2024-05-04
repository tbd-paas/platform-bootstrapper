gener8s:
	mkdir -p bin
	git clone https://github.com/nukleros/gener8s.git
	cd gener8s && \
	make build && \
	mv gener8s ../bin && \
	chmod +x ../bin/gener8s
	rm -rf gener8s

download:
	vendir sync --file vendor.yaml --lock-file vendor.yaml.lock

resources:
	echo 'package resources\n\nimport "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"\n\n' > ../internal/pkg/resources/generated.go
	bin/gener8s go -m manifests/operator/manifests.yaml >>  ../internal/pkg/resources/generated.go
	bin/gener8s go -m manifests/config/*.yaml >>  ../internal/pkg/resources/generated.go
	cd .. && go mod tidy