pb-gen:
	protoc -I protofiles/ protofiles/${service}.proto --go_out=plugins=grpc:services/${service}/pb && \
	protoc -I protofiles/ protofiles/${service}.proto --go_out=plugins=grpc:services/graphql/pb