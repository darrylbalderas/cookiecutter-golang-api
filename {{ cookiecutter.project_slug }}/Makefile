VERSION := 0.0.1
APP_NAME := {{cookiecutter.project_slug}}
IMAGE_NAME := ${APP_NAME}:${VERSION}

build:
	@docker build -t ${IMAGE_NAME} .
run:
	@docker run -p 8000:8000 ${IMAGE_NAME}