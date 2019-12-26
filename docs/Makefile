PWD := ${CURDIR}
DOCKER_IMAGE_NAME := "www9tjdoe"
CONTAINER_REMOVE_FLAG := "--rm"
BASE_URL := "https://tamada.github.io/tjdoe"
HUGO_THEME := "cayman-hugo-theme"

build:
	docker run ${CONTAINER_REMOVE_FLAG} --name ${DOCKER_IMAGE_NAME} -v ${PWD}:/src -v ${PWD}/public:/output -e HUGO_THEME=${HUGO_THEME} -e HUGO_BASEURL=${BASE_URL} jojomi/hugo:latest

start:
	docker run ${CONTAINER_REMOVE_FLAG} -d --name ${DOCKER_IMAGE_NAME} -p 1313:1313 -v "${PWD}":/src -v "$(PWD)"/public:/output -e HUGO_THEME=${HUGO_THEME} -e HUGO_WATCH="true" -e HUGO_BASEURL=${BASE_URL} jojomi/hugo:latest

stop:
	docker stop ${DOCKER_IMAGE_NAME}

