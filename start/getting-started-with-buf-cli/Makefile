gen-proto:
	buf format proto -w
	buf generate proto
	@echo "proto generate done"
	redocly build-docs -o docs/openapiv2/apidocs.html docs/openapiv2/apidocs.swagger.json
