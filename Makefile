ln-proto:
	# Copy proto_go/ inside the projects directory
	cp -r proto_go/ microservices/command_service/proto_go/
	cp -r proto_go/ microservices/query_service/proto_go/
	cp -r proto_go/ bff/proto_go/


build-bff:
	ln-proto
	
