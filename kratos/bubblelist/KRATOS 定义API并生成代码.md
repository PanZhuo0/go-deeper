#### KRATOS 定义API并生成代码（以清单项目为例）

- 新建一个bubble list 项目

```
kratos new bubblelist
```

- 增加.proto文件

```
cd bubblelist
kratos proto add xxx.proto

eg:
	kratos proto add api/bubblelist/v1/todo.proto
```

将会生成todo. proto  (底层基于go. template实现)

```protobuf
syntax = "proto3";

package api.bubblelist.v1;

option go_package = "bubblelist/api/bubblelist/v1;v1";
option java_multiple_files = true;
option java_package = "api.bubblelist.v1";

service Todo {
	rpc CreateTodo (CreateTodoRequest) returns (CreateTodoReply);
	rpc UpdateTodo (UpdateTodoRequest) returns (UpdateTodoReply);
	rpc DeleteTodo (DeleteTodoRequest) returns (DeleteTodoReply);
	rpc GetTodo (GetTodoRequest) returns (GetTodoReply);
	rpc ListTodo (ListTodoRequest) returns (ListTodoReply);
}

message CreateTodoRequest {}
message CreateTodoReply {}

message UpdateTodoRequest {}
message UpdateTodoReply {}

message DeleteTodoRequest {}
message DeleteTodoReply {}

message GetTodoRequest {}
message GetTodoReply {}

message ListTodoRequest {}
message ListTodoReply {}
```

- 对生成的todo . proto进行修改，改成自己想要的

```protobuf
syntax = "proto3";

package api.bubblelist.v1;

import "google/api/annotations.proto";

option go_package = "bubblelist/api/bubblelist/v1;v1";
option java_multiple_files = true;
option java_package = "api.bubblelist.v1";

service Todo {
	rpc CreateTodo (CreateTodoRequest) returns (CreateTodoReply){
		/* 基于grpc-gateway 实现提供HTTP服务: 需要提前了解grpc-gateway HTTP与GRPC的映射规则*/
		/* 当然，别忘了import google的注解扩展.proto
			import "google/api/annotations.proto"	
		*/
		option (google.api.http)={
			post: "v1/todo", //改grpc将被映射为创造一个一个http服务(post 请求，请求路径为: 域名/v1/todo)
			body: "*", //接收全部数据,映射到CreateTodoRequest中
		};
	};

	rpc UpdateTodo (UpdateTodoRequest) returns (UpdateTodoReply){
		option (google.api.http)={
			put:"/v1/todo/{id}",
			body:"*"
		};
	};

	rpc DeleteTodo (DeleteTodoRequest) returns (DeleteTodoReply){
		option (google.api.http)={
			delete: "/v1/todo/{id}",
		};
	};


	rpc GetTodo (GetTodoRequest) returns (GetTodoReply){
		option (google.api.http)={
			get:"/v1/todo/{id}",
		};
	};


	rpc ListTodo (ListTodoRequest) returns (ListTodoReply){
		option (google.api.http)={
			get:"/v1/todos",
		};
	};

}

/* 这样只会提供GRPC服务，如何提供HTTP服务?	-> 使用注解,基于grpc-gateway */
message todo{ //注意这里是小写的todo,和service Todo不同
	int64 id=1;	
	string title=2;
	bool status=3;
}

message CreateTodoRequest {
	string title =1;
}

message CreateTodoReply {
	int64 id=1;
	string title=2;
	bool status=3;
}

message UpdateTodoRequest {
	int64 id = 1;
	string title =2;
	bool status=3;
}

message UpdateTodoReply {}

message DeleteTodoRequest {
	int64 id = 1;
}

message DeleteTodoReply {}

message GetTodoRequest {
	int64 id =1;
}

message GetTodoReply {
	int64 id =1;
	string title=2;
	bool status =3;
}

/* 后期可以增加分页列举功能 */
message ListTodoRequest {}

message ListTodoReply {
	repeated todo todos = 1;
}
```

- 生成代码

```bash
# 支持 make
make api

# 不支持 make
kratos proto client api/bubblelist/v1/todo.proto  # 将会在对应目录下生成todo_grpc.pb.go todo_http.pb.go todo.pb.go文件
```

> KRATOS 框架将proto文件作为生成代码的依据