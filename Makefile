stop:
	docker-compose down
remove_volumes:
	docker volume rm cbot_cbot
run:
	docker-compose up
build:
	docker-compose build
reset: stop remove_volumes