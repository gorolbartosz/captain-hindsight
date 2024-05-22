build:
	go build -o bin/captain-hindsight cmd/main.go

build-image: build
	docker build --tag gorolbartosz/captain-hindsight . 

run-dependencies:
	docker run \
		--name mattermost \
		-e MM_SERVICESETTINGS_SITEURL=http://127.0.0.1:8065 \
		-e MM_SERVICESETTINGS_ALLOWEDUNTRUSTEDINTERNALCONNECTIONS=127.0.0.1 \
		--net=host \
		--detach \
		mattermost/mattermost-preview

run-local: build
	./bin/captain-hindsight

run-container: build-image
	docker run \
	--name captain-hindsight \
	--hostname captain-hindsight \
	--net=host \
	gorolbartosz/captain-hindsight

stop-dependencies:
	docker rm --force mattermost

open:
	xdg-open http://127.0.0.1:8065