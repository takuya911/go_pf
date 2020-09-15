pb-gen:
	protoc -I protofiles/ protofiles/${service}.proto --go_out=plugins=grpc:services/${service}/proto