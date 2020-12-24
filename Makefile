all: shell

build:
	docker build -t golearn .

shell:
	docker run -v $$(pwd):/golearn -ti golearn bash
