.PHONY: start down restart logs

start:
	docker-compose up -d
	docker-compose logs -f --tail 16

stop:
	docker-compose down

restart:
	docker-compose restart
	docker-compose logs -f --tail 16

logs:
	docker-compose logs -f --tail 16
