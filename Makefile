generate_grpc_code:
	rm -rf gen && mkdir gen
	export PATH="$${PATH}:$$(go env GOPATH)/bin";\
	protoc --proto_path=proto proto/*.proto --go_out=gen/ --go-grpc_out=gen/ -I . \
	--grpc-gateway_out ./gen \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt generate_unbound_methods=true

generate_grpc_code_windows:
	rm -rf gen && mkdir gen
	set PATH=%PATH%;%GOPATH%\bin &\
	protoc --proto_path=proto proto\*.proto --go_out=gen\ --go-grpc_out=gen\ -I . \
	--grpc-gateway_out .\gen \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt generate_unbound_methods=true

download_protos:
	mkdir -p proto/google/api
	curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto > proto/google/api/annotations.proto
	curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto > proto/google/api/http.proto