proto:
	protoc -I . \
		   -I ${GOPATH}/src \
           --go_out ./ --go_opt paths=source_relative \
           --go-grpc_out=./ --go-grpc_opt=paths=source_relative \
           protos/model/route_guide.proto

	   ls protos/model/route_guide.pb.go | xargs -n1 -IX bash -c 'sed s/,omitempty// X > X.tmp && mv X{.tmp,}'


