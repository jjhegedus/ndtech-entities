module github.com/jjhegedus/ndtech-entities

go 1.15

require (
	github.com/jjhegedus/ndtech-entities/protos v0.0.0-00010101000000-000000000000 // indirect
	google.golang.org/grpc v1.36.0 // indirect
)

replace github.com/jjhegedus/ndtech-entities/protos => ./bazel-bin/protos/entities_go_proto_/github.com/jjhegedus/ndtech-entities/protos/
