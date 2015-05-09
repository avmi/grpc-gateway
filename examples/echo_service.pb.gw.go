// Code generated by protoc-gen-grpc-gateway
// source: examples/echo_service.proto
// DO NOT EDIT!

/*
Package main is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gengo/grpc-gateway/runtime"
	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var _ codes.Code
var _ io.Reader
var _ = runtime.String
var _ = json.Marshal

func request_EchoService_Echo_0(ctx context.Context, client EchoServiceClient, req *http.Request, pathParams map[string]string) (msg proto.Message, err error) {
	var protoReq SimpleMessage

	var val string
	var ok bool

	val, ok = pathParams["id"]
	if !ok {
		return nil, grpc.Errorf(codes.InvalidArgument, "missing parameter %s", "id")
	}
	protoReq.Id, err = runtime.String(val)
	if err != nil {
		return nil, err
	}

	return client.Echo(ctx, &protoReq)
}

func request_EchoService_EchoBody_0(ctx context.Context, client EchoServiceClient, req *http.Request, pathParams map[string]string) (msg proto.Message, err error) {
	var protoReq SimpleMessage

	if err = json.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, grpc.Errorf(codes.InvalidArgument, "%v", err)
	}

	return client.EchoBody(ctx, &protoReq)
}

// RegisterEchoServiceHandlerFromEndpoint is same as RegisterEchoServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterEchoServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string) (err error) {
	conn, err := grpc.Dial(endpoint)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				glog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				glog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterEchoServiceHandler(ctx, mux, conn)
}

// RegisterEchoServiceHandler registers the http handlers for service EchoService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterEchoServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	client := NewEchoServiceClient(conn)

	mux.Handle("POST", pattern_EchoService_Echo_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		resp, err := request_EchoService_Echo_0(ctx, client, req, pathParams)
		if err != nil {
			runtime.HTTPError(w, err)
			return
		}

		runtime.ForwardResponseMessage(w, resp)

	})

	mux.Handle("POST", pattern_EchoService_EchoBody_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		resp, err := request_EchoService_EchoBody_0(ctx, client, req, pathParams)
		if err != nil {
			runtime.HTTPError(w, err)
			return
		}

		runtime.ForwardResponseMessage(w, resp)

	})

	return nil
}

var (
	pattern_EchoService_Echo_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 3}, []string{"v1", "example", "echo", "id"}, ""))

	pattern_EchoService_EchoBody_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1", "example", "echo_body"}, ""))
)
