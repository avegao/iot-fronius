TAG ?= master
VCS_REF = $(shell git rev-parse --verify HEAD)
BUILD_DATE = $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

build-develop:
	docker build \
		-t avegao/iot-fronius\:develop \
		--build-arg="VCS_REF=$(VCS_REF)" \
		--build-arg="BUILD_DATE=$(BUILD_DATE)" \
		.

push-develop:
	docker push avegao/iot-fronius\:develop

publish-develop: build-develop push-develop

tag-develop-version:
	docker tag avegao/iot-fronius\:develop avegao/iot-fronius\:$(VERSION)
	docker push avegao/iot-fronius\:$(VERSION)

promote-develop-to-master:
	docker pull avegao/iot-fronius\:develop
	docker tag avegao/iot-fronius\:develop avegao/iot-fronius\:master
	docker tag avegao/iot-fronius\:develop avegao/iot-fronius\:latest
	docker push avegao/iot-fronius\:master
	docker push avegao/iot-fronius\:latest

build-master:
	docker build \
		-t avegao/iot-fronius\:master \
		-t avegao/iot-fronius\:latest \
		--build-arg="VCS_REF=$(VCS_REF)" \
		--build-arg="BUILD_DATE=$(BUILD_DATE)" \
		.

push-master:
	docker push avegao/iot-fronius\:master
	docker push avegao/iot-fronius\:latest

publish-master: build-master push-master

tag-version:
	docker pull avegao/iot-fronius\:master
	docker tag avegao/iot-fronius\:master avegao/iot-fronius\:$(VERSION)
	docker push avegao/iot-fronius\:$(VERSION)
