SED=$(shell which sed)
.DEFAULT_GOAL := help
background_option=-d
nickname=
repository_name=$(shell basename $(PWD))

setup/mac: $(nickname)
	$(MAKE) setup/bsd

setup/bsd: $(nickname) ## for mac
	$(SED) -i '' -e 's/original/$(nickname)/g' ./$(nickname)/*.go
	$(SED) -i '' -e 's/original/$(nickname)/g' ./$(nickname)/**/*.go
	$(SED) -i '' -e 's/vg-1day-2018/$(repository_name)/g' ./$(nickname)/*.go
	$(SED) -i '' -e 's/vg-1day-2018/$(repository_name)/g' ./$(nickname)/**/*.go

setup/gnu: $(nickname) ## for linux
	$(SED) --in-place 's/original/$(nickname)/g' ./$(nickname)/*.go
	$(SED) --in-place 's/original/$(nickname)/g' ./$(nickname)/**/*.go
	$(SED) --in-place 's/vg-1day-2018/$(repository_name)/g' ./$(nickname)/*.go
	$(SED) --in-place 's/vg-1day-2018/$(repository_name)/g' ./$(nickname)/**/*.go

$(nickname):
	cp -rf original $(nickname)

help:
	@echo docker_build:	Build the docker container
	@echo docker_up:	Start the docker container
	@echo docker_stop:	Stop the docker container
	@echo docker_rm:	Remove the docker container
	@echo docker_ssh:	Execute an interactive bash shell on the container

docker/build:
	docker build -t $(repository_name) .

docker/run:
	docker run --rm --name $(repository_name) -p 8080:8080 -v $(CURDIR):/go/src/github.com/VG-Tech-Dojo/vg-1day-2018 -it $(repository_name)

docker/run/%: $(@F)
	docker run --rm --name $(repository_name) -p 8080:8080 -v $(CURDIR):/go/src/github.com/VG-Tech-Dojo/vg-1day-2018 -it $(repository_name) -C $(@F) run
