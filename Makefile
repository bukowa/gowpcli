
dump:
	docker build --tag=test . && \
	docker run --volume=${PWD}/generate:/app --rm -it test