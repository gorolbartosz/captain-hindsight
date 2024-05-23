build:
	go build -o bin/captain-hindsight cmd/main.go

build-image:
	docker build --tag gorolbartosz/captain-hindsight . 

run-dependencies:
	docker run \
		--name mattermost \
		--env MM_SERVICESETTINGS_SITEURL=http://127.0.0.1:8065 \
		--env MM_SERVICESETTINGS_ALLOWEDUNTRUSTEDINTERNALCONNECTIONS=127.0.0.1 \
		--net=host \
		--detach \
		mattermost/mattermost-preview

run-local: build
	./bin/captain-hindsight

run-local-container: build-image
	docker run \
	--name captain-hindsight \
	--hostname captain-hindsight \
	--net=host \
	--detach \
	gorolbartosz/captain-hindsight

stop-dependencies:
	docker rm --force mattermost

stop-local-container:
	docker rm --force captain-hindsight

logs-dependencies:
	docker logs --follow mattermost

logs:
	docker logs --follow captain-hindsight

open:
	xdg-open http://127.0.0.1:8065
