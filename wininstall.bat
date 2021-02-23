@ECHO OFF

SET SHA_CMD="git rev-parse --short=10 HEAD"
FOR /F "tokens=* USEBACKQ" %%F IN (`%SHA_CMD%`) DO (
	SET SHA=%%F
)

SET HEAD_CMD="git rev-parse HEAD"
FOR /F "tokens=* USEBACKQ" %%F IN (`%HEAD_CMD%`) DO (
	SET HEAD_COMMIT=%%F
)

SET HEAD_DATE_CMD="git show -s --format=%%ct %HEAD_COMMIT%"
FOR /F "tokens=* USEBACKQ" %%F IN (`%HEAD_DATE_CMD%`) DO (
	SET GIT_COMMIT_EPOC=%%F
)

SET DATE_FMT_CMD="go-datefmt -ts %GIT_COMMIT_EPOC% -fmt UnixDate -utc"
FOR /F "tokens=* USEBACKQ" %%F IN (`%DATE_FMT_CMD%`) DO (
	SET HEAD_DATE=%%F
)

@ECHO ON
go get github.com/pauln/go-datefmt

:: go get github.com/gogo/protobuf/protoc-gen-gogo
go get -u github.com/gogo/protobuf/protoc-gen-gogo@21df5aa0e680850681b8643f0024f92d3b09930c
go get -u github.com/gogo/protobuf/protoc-gen-gogofaster@21df5aa0e680850681b8643f0024f92d3b09930c

:: go get github.com/gogo/protobuf/proto
go get -u github.com/gogo/protobuf/proto@21df5aa0e680850681b8643f0024f92d3b09930c

:: go get github.com/jteeuwen/go-bindata/...
go get -u github.com/kevinburke/go-bindata/go-bindata

:: go get  github.com/grpc-ecosystem/grpc-gateway/v2/...
go get -u github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2

:: go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

:: go get google.golang.org/protobuf/cmd/protoc-gen-go
go get -u google.golang.org/protobuf/cmd/protoc-gen-go

:: go get github.com/envoyproxy/protoc-gen-validate
go get -u github.com/envoyproxy/protoc-gen-validate

go generate github.com/Reasno/trs/gengokit/template
go install -ldflags "-X 'main.version=%SHA%' -X 'main.date=%HEAD_DATE%'" github.com/Reasno/trs/cmd/tr
@ECHO OFF
