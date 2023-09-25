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
unit-test-controller:
	go test -v ./tests/unit/user.controller_test.go \
 		-covermode=count \
 		-coverpkg=./controllers -coverprofile ./tests/report/user.controller.coverage.out
	go tool cover -html ./tests/report/user.controller.coverage.out \
		-o ./tests/report/user.controller.coverage.html

unit-test-service:
	go test -v ./tests/unit/user.service_test.go \
 		-covermode=count \
 		-coverpkg=./services -coverprofile ./tests/report/user.service.coverage.out
	go tool cover -html ./tests/report/user.service.coverage.out \
		-o ./tests/report/user.service.coverage.html
