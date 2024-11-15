gen:
	go run gen/*.go

.PHONY: gen

update-webservice:
	curl -s https://sonarcloud.io/api/webservices/list | jq > gen/webservices.json 