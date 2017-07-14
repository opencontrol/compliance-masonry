build:
	docker build -t masonry .

build_test:
	docker build -t masonry-test - < Dockerfile.test

test: build_test
	docker run -it -v `pwd`:/go/src/github.com/opencontrol/compliance-masonry masonry-test
