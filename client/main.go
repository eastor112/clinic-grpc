package main

import (
	"context"
	"log"
	"net/http"

	"buf.build/gen/go/eastor112/clinic-grpc/connectrpc/go/patients/v1/patientsv1connect"
	patientsV1 "buf.build/gen/go/eastor112/clinic-grpc/protocolbuffers/go/patients/v1"
	connect "connectrpc.com/connect"
)

func main() {
	client := patientsv1connect.NewPatientServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
	)
	res, err := client.GetPatientByID(
		context.Background(),
		connect.NewRequest(&patientsV1.PatientRequest{
			PatientId: 2,
		}),
	)
	if err != nil {
		log.Println("error!!", err)
		return
	}
	log.Println(res.Msg)
}
