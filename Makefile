run-command-svc:
	cd ./command-service/ && dapr run \
		--app-port 8888 \
		--app-id command-service \
		--app-protocol http \
		--dapr-http-port 3035 \
		--dapr-grpc-port 3036 \
		--resources-path ../components \
		-- go run .