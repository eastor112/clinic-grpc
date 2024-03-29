// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: patients/v1/patient.proto

package patientsv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/eastor112/clinic-grpc/gen/patients/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// PatientServiceName is the fully-qualified name of the PatientService service.
	PatientServiceName = "patients.v1.PatientService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// PatientServiceGetPatientByIDProcedure is the fully-qualified name of the PatientService's
	// GetPatientByID RPC.
	PatientServiceGetPatientByIDProcedure = "/patients.v1.PatientService/GetPatientByID"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	patientServiceServiceDescriptor              = v1.File_patients_v1_patient_proto.Services().ByName("PatientService")
	patientServiceGetPatientByIDMethodDescriptor = patientServiceServiceDescriptor.Methods().ByName("GetPatientByID")
)

// PatientServiceClient is a client for the patients.v1.PatientService service.
type PatientServiceClient interface {
	GetPatientByID(context.Context, *connect.Request[v1.PatientRequest]) (*connect.Response[v1.PatientResponse], error)
}

// NewPatientServiceClient constructs a client for the patients.v1.PatientService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewPatientServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) PatientServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &patientServiceClient{
		getPatientByID: connect.NewClient[v1.PatientRequest, v1.PatientResponse](
			httpClient,
			baseURL+PatientServiceGetPatientByIDProcedure,
			connect.WithSchema(patientServiceGetPatientByIDMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// patientServiceClient implements PatientServiceClient.
type patientServiceClient struct {
	getPatientByID *connect.Client[v1.PatientRequest, v1.PatientResponse]
}

// GetPatientByID calls patients.v1.PatientService.GetPatientByID.
func (c *patientServiceClient) GetPatientByID(ctx context.Context, req *connect.Request[v1.PatientRequest]) (*connect.Response[v1.PatientResponse], error) {
	return c.getPatientByID.CallUnary(ctx, req)
}

// PatientServiceHandler is an implementation of the patients.v1.PatientService service.
type PatientServiceHandler interface {
	GetPatientByID(context.Context, *connect.Request[v1.PatientRequest]) (*connect.Response[v1.PatientResponse], error)
}

// NewPatientServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewPatientServiceHandler(svc PatientServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	patientServiceGetPatientByIDHandler := connect.NewUnaryHandler(
		PatientServiceGetPatientByIDProcedure,
		svc.GetPatientByID,
		connect.WithSchema(patientServiceGetPatientByIDMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/patients.v1.PatientService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case PatientServiceGetPatientByIDProcedure:
			patientServiceGetPatientByIDHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedPatientServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedPatientServiceHandler struct{}

func (UnimplementedPatientServiceHandler) GetPatientByID(context.Context, *connect.Request[v1.PatientRequest]) (*connect.Response[v1.PatientResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("patients.v1.PatientService.GetPatientByID is not implemented"))
}
