.PHONY:	docker run

all:	build

build:
	docker build -t my-todo .

run:
	docker run -it -d --rm --name my-todo -p 8080:8080 my-todo
