proto:
	protoc -I . \
		   -I ${GOPATH}/src \
           --go_out ./ --go_opt paths=source_relative \
           protos/model/sample.proto

	   ls protos/model/sample.pb.go | xargs -n1 -IX bash -c 'sed s/,omitempty// X > X.tmp && mv X{.tmp,}'


