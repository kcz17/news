NAME = kcz17/news
DBNAME = kcz17/news-db

INSTANCE = news

.PHONY: default copy test

default: test

release:
	docker build -t $(NAME) -f ./docker/news/Dockerfile .

#dockertravisbuild: build
#	docker build -t $(NAME):$(TAG) -f docker/news/Dockerfile-release docker/news/
#	docker build -t $(DBNAME):$(TAG) -f docker/news-db/Dockerfile docker/news-db/
#	docker login -u $(DOCKER_USER) -p $(DOCKER_PASS)
#	scripts/push.sh
