
dump:
	docker build --tag=test . && \
	docker run --volume=${PWD}/generate:/app --rm -it test

generate:
	python -m generate

.PHONY: dump generate
