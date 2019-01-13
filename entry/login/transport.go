package login

import (
	"encoding/json"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/ratelimit"
	"github.com/go-kit/kit/tracing/opentracing"
	"github.com/go-kit/kit/tracing/zipkin"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	httptransport "github.com/go-kit/kit/transport/http"
	stdopentracing "github.com/opentracing/opentracing-go"
	stdzipkin "github.com/openzipkin/zipkin-go"
	"github.com/sony/gobreaker"
	"golang.org/x/net/context"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"net/http"
	"time"
	"zskparker.com/foundation/entry/login/pb"
	"zskparker.com/foundation/pkg/format"
	"zskparker.com/foundation/pkg/functions"
	"zskparker.com/foundation/pkg/transport"
)

type GRPCServer struct {
	entrybyap           grpctransport.Handler
	entrybyoauth        grpctransport.Handler
	entrybyvalidatecode grpctransport.Handler
	entrybyqrcode       grpctransport.Handler
	entrybyface         grpctransport.Handler
	entrybyinvite       grpctransport.Handler
}

func (g *GRPCServer) EntryByInvite(ctx context.Context, in *fs_entry_login.EntryByInviteRequest) (*fs_entry_login.EntryResponse, error) {
	_, resp, err := g.entrybyinvite.ServeGRPC(ctx, in)
	if err != nil {
		return &fs_entry_login.EntryResponse{State: fs_metadata_transport.GetResponseState(err, resp)}, nil
	}
	return resp.(*fs_entry_login.EntryResponse), nil
}

func (g *GRPCServer) EntryByAP(ctx context.Context, in *fs_entry_login.EntryByAPRequest) (*fs_entry_login.EntryResponse, error) {
	_, resp, err := g.entrybyap.ServeGRPC(ctx, in)
	if err != nil {
		return &fs_entry_login.EntryResponse{State: fs_metadata_transport.GetResponseState(err, resp)}, nil
	}
	return resp.(*fs_entry_login.EntryResponse), nil
}

func (g *GRPCServer) EntryByOAuth(ctx context.Context, in *fs_entry_login.EntryByOAuthRequest) (*fs_entry_login.EntryResponse, error) {
	_, resp, err := g.entrybyoauth.ServeGRPC(ctx, in)
	if err != nil {
		return &fs_entry_login.EntryResponse{State: fs_metadata_transport.GetResponseState(err, resp)}, nil
	}
	return resp.(*fs_entry_login.EntryResponse), nil
}

func (g *GRPCServer) EntryByValidateCode(ctx context.Context, in *fs_entry_login.EntryByValidateCodeRequest) (*fs_entry_login.EntryResponse, error) {
	_, resp, err := g.entrybyvalidatecode.ServeGRPC(ctx, in)
	if err != nil {
		return &fs_entry_login.EntryResponse{State: fs_metadata_transport.GetResponseState(err, resp)}, nil
	}
	return resp.(*fs_entry_login.EntryResponse), nil
}

func (g *GRPCServer) EntryByQRCode(ctx context.Context, in *fs_entry_login.EntryByQRCodeRequest) (*fs_entry_login.EntryResponse, error) {
	_, resp, err := g.entrybyqrcode.ServeGRPC(ctx, in)
	if err != nil {
		return &fs_entry_login.EntryResponse{State: fs_metadata_transport.GetResponseState(err, resp)}, nil
	}
	return resp.(*fs_entry_login.EntryResponse), nil
}

func (g *GRPCServer) EntryByFace(ctx context.Context, in *fs_entry_login.EntryByFaceRequest) (*fs_entry_login.EntryResponse, error) {
	_, resp, err := g.entrybyface.ServeGRPC(ctx, in)
	if err != nil {
		return &fs_entry_login.EntryResponse{State: fs_metadata_transport.GetResponseState(err, resp)}, nil
	}
	return resp.(*fs_entry_login.EntryResponse), nil
}

func MakeHTTPHandler(endpoints Endpoints, otTracer stdopentracing.Tracer, zipkinTracer *stdzipkin.Tracer, logger log.Logger) http.Handler {
	zipkinServer := zipkin.HTTPServerTrace(zipkinTracer)

	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerBefore(fs_metadata_transport.HTTPToContext()),
		zipkinServer,
	}

	m := http.NewServeMux()
	m.Handle(fs_functions.GetEntryByAPFunc().Infix, httptransport.NewServer(
		endpoints.EntryByAPEndpoint,
		decodeEntryByAP,
		format.EncodeHTTPGenericResponse,
		append(options, httptransport.ServerBefore(opentracing.HTTPToContext(otTracer, "EntryByAP", logger)))...,
	))

	m.Handle(fs_functions.GetEntryByQRCodeFunc().Infix, httptransport.NewServer(
		endpoints.EntryByQRCodeEndpoint,
		decodeEntryByQRCode,
		format.EncodeHTTPGenericResponse,
		append(options, httptransport.ServerBefore(opentracing.HTTPToContext(otTracer, "EntryByQRCode", logger)))...,
	))

	m.Handle(fs_functions.GetEntryByOAuthFunc().Infix, httptransport.NewServer(
		endpoints.EntryByOAuthEndpoint,
		decodeEntryByOAuth,
		format.EncodeHTTPGenericResponse,
		append(options, httptransport.ServerBefore(opentracing.HTTPToContext(otTracer, "EntryByOAuth", logger)))...,
	))

	m.Handle(fs_functions.GetEntryByFaceFunc().Infix, httptransport.NewServer(
		endpoints.EntryByFaceEndpoint,
		decodeEntryByFace,
		format.EncodeHTTPGenericResponse,
		append(options, httptransport.ServerBefore(opentracing.HTTPToContext(otTracer, "EntryByFace", logger)))...,
	))

	m.Handle(fs_functions.GetEntryByValidateCodeFunc().Infix, httptransport.NewServer(
		endpoints.EntryByValidateCodeEndpoint,
		decodeEntryByValidateCode,
		format.EncodeHTTPGenericResponse,
		append(options, httptransport.ServerBefore(opentracing.HTTPToContext(otTracer, "EntryByValidateCode", logger)))...,
	))

	m.Handle(fs_functions.GetEntryByInviteFunc().Infix, httptransport.NewServer(
		endpoints.EntryByInviteEndpoint,
		decodeEntryByInvite,
		format.EncodeHTTPGenericResponse,
		append(options, httptransport.ServerBefore(opentracing.HTTPToContext(otTracer, "EntryByInvite", logger)))...,
	))

	return m
}

func decodeEntryByInvite(_ context.Context, r *http.Request) (interface{}, error) {
	var req *fs_entry_login.EntryByInviteRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func decodeEntryByAP(_ context.Context, r *http.Request) (interface{}, error) {
	var req *fs_entry_login.EntryByAPRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func decodeEntryByValidateCode(_ context.Context, r *http.Request) (interface{}, error) {
	var req *fs_entry_login.EntryByValidateCodeRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func decodeEntryByQRCode(_ context.Context, r *http.Request) (interface{}, error) {
	var req *fs_entry_login.EntryByQRCodeRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func decodeEntryByOAuth(_ context.Context, r *http.Request) (interface{}, error) {
	var req *fs_entry_login.EntryByOAuthRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func decodeEntryByFace(_ context.Context, r *http.Request) (interface{}, error) {
	var req *fs_entry_login.EntryByFaceRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func MakeGRPCServer(endpoints Endpoints, otTracer stdopentracing.Tracer, tracer *stdzipkin.Tracer, logger log.Logger) fs_entry_login.LoginServer {
	zipkinServer := zipkin.GRPCServerTrace(tracer)

	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
		grpctransport.ServerBefore(fs_metadata_transport.GRPCToContext()),
		zipkinServer,
	}

	return &GRPCServer{
		entrybyap: grpctransport.NewServer(
			endpoints.EntryByAPEndpoint,
			format.GrpcMessage,
			format.GrpcMessage,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer, "EntryByAP", logger)))...),
		entrybyface: grpctransport.NewServer(
			endpoints.EntryByFaceEndpoint,
			format.GrpcMessage,
			format.GrpcMessage,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer, "EntryByFace", logger)))...),
		entrybyoauth: grpctransport.NewServer(
			endpoints.EntryByOAuthEndpoint,
			format.GrpcMessage,
			format.GrpcMessage,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer, "EntryByOAuth", logger)))...),
		entrybyqrcode: grpctransport.NewServer(
			endpoints.EntryByQRCodeEndpoint,
			format.GrpcMessage,
			format.GrpcMessage,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer, "EntryByQRCode", logger)))...),
		entrybyvalidatecode: grpctransport.NewServer(
			endpoints.EntryByValidateCodeEndpoint,
			format.GrpcMessage,
			format.GrpcMessage,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer, "EntryByValidateCode", logger)))...),
		entrybyinvite: grpctransport.NewServer(
			endpoints.EntryByInviteEndpoint,
			format.GrpcMessage,
			format.GrpcMessage,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer, "EntryByInvite", logger)))...),
	}
}

func MakeGRPCClient(conn *grpc.ClientConn, otTracer stdopentracing.Tracer, zipkinTracer *stdzipkin.Tracer, logger log.Logger) Service {
	limiter := ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 100))
	zipkinClient := zipkin.GRPCClientTrace(zipkinTracer)

	options := []grpctransport.ClientOption{
		zipkinClient,
		grpctransport.ClientBefore(fs_metadata_transport.ContextToGRPC()),
	}

	var entryByAPEndpoint endpoint.Endpoint
	{
		entryByAPEndpoint = grpctransport.NewClient(conn,
			"fs.entry.login.Login",
			"EntryByAP",
			format.GrpcMessage,
			format.GrpcMessage,
			fs_entry_login.EntryResponse{},
			append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer, logger)))...).Endpoint()
		entryByAPEndpoint = limiter(entryByAPEndpoint)
		entryByAPEndpoint = opentracing.TraceClient(otTracer, "EntryByAP")(entryByAPEndpoint)
		entryByAPEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "EntryByAP",
			Timeout: 5 * time.Second,
		}))(entryByAPEndpoint)
	}

	var entryByOAuthEndpoint endpoint.Endpoint
	{
		entryByOAuthEndpoint = grpctransport.NewClient(conn,
			"fs.entry.login.Login",
			"EntryByOAuth",
			format.GrpcMessage,
			format.GrpcMessage,
			fs_entry_login.EntryResponse{},
			append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer, logger)))...).Endpoint()
		entryByOAuthEndpoint = limiter(entryByOAuthEndpoint)
		entryByOAuthEndpoint = opentracing.TraceClient(otTracer, "EntryByOAuth")(entryByOAuthEndpoint)
		entryByOAuthEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "EntryByOAuth",
			Timeout: 5 * time.Second,
		}))(entryByOAuthEndpoint)
	}

	var entryByValidateCodeEndpoint endpoint.Endpoint
	{
		entryByValidateCodeEndpoint = grpctransport.NewClient(conn,
			"fs.entry.login.Login",
			"EntryByValidateCode",
			format.GrpcMessage,
			format.GrpcMessage,
			fs_entry_login.EntryResponse{},
			append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer, logger)))...).Endpoint()
		entryByValidateCodeEndpoint = limiter(entryByValidateCodeEndpoint)
		entryByValidateCodeEndpoint = opentracing.TraceClient(otTracer, "EntryByValidateCode")(entryByValidateCodeEndpoint)
		entryByValidateCodeEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "EntryByValidateCode",
			Timeout: 5 * time.Second,
		}))(entryByValidateCodeEndpoint)
	}

	var entryByQRCodeEndpoint endpoint.Endpoint
	{
		entryByQRCodeEndpoint = grpctransport.NewClient(conn,
			"fs.entry.login.Login",
			"EntryByQRCode",
			format.GrpcMessage,
			format.GrpcMessage,
			fs_entry_login.EntryResponse{},
			append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer, logger)))...).Endpoint()
		entryByQRCodeEndpoint = limiter(entryByQRCodeEndpoint)
		entryByQRCodeEndpoint = opentracing.TraceClient(otTracer, "EntryByQRCode")(entryByQRCodeEndpoint)
		entryByQRCodeEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "EntryByQRCode",
			Timeout: 5 * time.Second,
		}))(entryByQRCodeEndpoint)
	}

	var entryByFaceEndpoint endpoint.Endpoint
	{
		entryByFaceEndpoint = grpctransport.NewClient(conn,
			"fs.entry.login.Login",
			"EntryByFace",
			format.GrpcMessage,
			format.GrpcMessage,
			fs_entry_login.EntryResponse{},
			append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer, logger)))...).Endpoint()
		entryByFaceEndpoint = limiter(entryByFaceEndpoint)
		entryByFaceEndpoint = opentracing.TraceClient(otTracer, "EntryByFace")(entryByFaceEndpoint)
		entryByFaceEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "EntryByFace",
			Timeout: 5 * time.Second,
		}))(entryByFaceEndpoint)
	}

	var entryByInviteEndpoint endpoint.Endpoint
	{
		entryByInviteEndpoint = grpctransport.NewClient(conn,
			"fs.entry.login.Login",
			"EntryByInvite",
			format.GrpcMessage,
			format.GrpcMessage,
			fs_entry_login.EntryResponse{},
			append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer, logger)))...).Endpoint()
		entryByInviteEndpoint = limiter(entryByInviteEndpoint)
		entryByInviteEndpoint = opentracing.TraceClient(otTracer, "EntryByInvite")(entryByInviteEndpoint)
		entryByInviteEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "EntryByInvite",
			Timeout: 5 * time.Second,
		}))(entryByInviteEndpoint)
	}

	return Endpoints{
		EntryByQRCodeEndpoint:       entryByQRCodeEndpoint,
		EntryByOAuthEndpoint:        entryByOAuthEndpoint,
		EntryByFaceEndpoint:         entryByFaceEndpoint,
		EntryByAPEndpoint:           entryByAPEndpoint,
		EntryByValidateCodeEndpoint: entryByValidateCodeEndpoint,
		EntryByInviteEndpoint:       entryByInviteEndpoint,
	}

}
