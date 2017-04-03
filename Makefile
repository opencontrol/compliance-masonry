build:
	docker build -t masonry .

build_test:
	docker build -f Dockerfile.test -t masonry-test .

test: build_test
	docker run -it -v `pwd`:/go/src/github.com/opencontrol/compliance-masonry masonry-test
