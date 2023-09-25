build-app:
	docker-compose build --no-cache app
start:
	docker-compose up -d
restart:
	docker restart api-searching
stop:
	docker-compose down
logs:
	docker logs -f api-searching
ssh-app:
	docker exec -it api-searching bash

unit-test:
	go test tests/unit/user.service_test.go -v

integration-test:
	go test tests/unit/user.service_test.go -v
