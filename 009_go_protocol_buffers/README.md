# ProtoBuf 
전송하기 위한 저장하려고 하는 데이터 구조를 유연하고 효율적으로 작성하고 파싱할 수 있도록 지원하는 자동화된 솔루션이다.  
Json이나 XML을 생각하면 이해가 쉬워지는데, 이것들과는 조금 다르다.    
ProtoBuf는 데이터 구조를 .proto라는 파일로 작성하면, 이를 컴파일해서 C++/C#/Python 등의 언어 코드 형태로 변환이 가능하고,  
변환된 언어 형태로 모듈까지 지원해서 데이터 구조를 전송하고 받아 읽을 수 있게 Serialize / Deserialize도 해준다.  
C언어로 예를 들어 구조체를 정의하고, 이를 소켓을 통해 전송할 수 있게 Serialize(전송할 수 있도록 변환)하고,  
변환된 데이터를 받은 후 Deserialize(파싱해서 다시 구조체의 형태로 저장) 하던 과정이 있다면, 이를 직접 코드로 다 작성해야 한다.  
하지만 ProtoBuf는 데이터 구조만 정의하면 이 과정은 전부 지원해준다.

## ProtoBuf에 대한 컴파일 명령어

- 하나만 지정해서 proto 파일 빌드하기 

```shell
$ protoc -I=./proto/model/ --go_out=./  ./proto/model/data.proto
```

```protobuf
syntax = "proto3";

// package 명은 go package 에 지정된 경로의 제일 마지막 것으로 해야 한다.
package proto.model;

// 위 경로는 실제 배포시 적용될 것으로 해준다.
// protoc -I=./proto/model/ --go_out=./  ./proto/model/data.proto 가 실행되면 실제 컴파일될 경로를 정의한다. 
option go_package = "github.com/keepinmindsh/proto/model";


message MyData {
    string msg = 1;
    int32 number = 2;
}
```

```shell
$ protoc -I=./proto/model/ --go_out=./  ./proto/model/*.proto
```

```protobuf
syntax = "proto3";

// package 명은 go package 에 지정된 경로의 제일 마지막 것으로 해야 한다.
package proto.model;
option go_package = "github.com/keepinmindsh/proto/model";
// 위 경로는 실제 배포시 적용될 것으로 해준다.

enum Corpus {
    CORPUS_UNSPECIFIED = 0;
    CORPUS_UNIVERSAL = 1;
    CORPUS_WEB = 2;
    CORPUS_IMAGES = 3;
    CORPUS_LOCAL = 4;
    CORPUS_NEWS = 5;
    CORPUS_PRODUCTS = 6;
    CORPUS_VIDEO = 7;
}

enum EnumAllowingAlias {
    option allow_alias = true;
    EAA_UNSPECIFIED = 0;
    EAA_STARTED = 1;
    EAA_RUNNING = 1;
    EAA_FINISHED = 2;
}
enum EnumNotAllowingAlias {
    ENAA_UNSPECIFIED = 0;
    ENAA_STARTED = 1;
    // ENAA_RUNNING = 1;  // Uncommenting this line will cause a compile error inside Google and a warning message outside.
    ENAA_FINISHED = 2;
}

message SearchRequest {
    string query = 1;
    int32 page_number = 2;
    int32 result_per_page = 3;
}
```

- 해당 경로의 모든 proto 파일 지정하기
    - * 에 대해서 모두 표기


## Services 

- Service를 사용하기 위해서는 gRPC Plugin을 사용해야함.



> [Protocol Buffers](https://developers.google.com/protocol-buffers/docs/proto3)      
> [Go Generated Code](https://developers.google.com/protocol-buffers/docs/reference/go-generated)