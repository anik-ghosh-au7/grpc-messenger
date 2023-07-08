generate_grpc_code:
	rm -rf gen && mkdir gen
	export PATH="$${PATH}:$$(go env GOPATH)/bin";\
	protoc --proto_path=proto proto/*.proto --go_out=gen/ --go-grpc_out=gen/

generate_grpc_code_windows:
	rm -rf gen && mkdir gen
	set PATH=%PATH%;%GOPATH%\bin &\
	protoc --proto_path=proto proto\*.proto --go_out=gen\ --go-grpc_out=gen\\