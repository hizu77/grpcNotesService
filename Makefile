CURDIR=$(shell pwd)
GENERATEDDIR=${CURDIR}/pkg/generated
PROTODIR=${GENERATEDDIR}/proto

generate_proto:
	rm -rf ${PROTODIR}
	mkdir -p ${PROTODIR}
	protoc --proto_path=${CURDIR}/internal/api/proto \
	--go_out=${PROTODIR} \
	--go-grpc_out=${PROTODIR} \
	${CURDIR}/internal/api/proto/note.proto
	go mod tidy
run: 
	go run cmd/notes/main.go
