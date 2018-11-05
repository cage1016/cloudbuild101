run:
	@go run main.go

test:
	@go test -ldflags -s -v

docker-build:
	docker build -t cage1016/bonego:latest -f Dockerfile.local .

docker-build-debug:
	docker build -t cage1016/bonego-debug:latest -f Dockerfile.debug .	

build-local:
	cloud-build-local --config cloudbuild.yaml --dryrun=false --write-workspace=./tmp .

build-gcloud:
	gcloud builds submit --config cloudbuild.yaml .

.PHONY: run test build-gcloud docker-build build-local
