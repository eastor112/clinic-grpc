package main

import (
	"context"
	"fmt"
	"net/http"

	"connectrpc.com/connect"
	"github.com/eastor112/clinic-grpc/configs"
	"github.com/eastor112/clinic-grpc/models"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/protobuf/types/known/timestamppb"

	patientsv1 "github.com/eastor112/clinic-grpc/gen/patients/v1"
	"github.com/eastor112/clinic-grpc/gen/patients/v1/patientsv1connect"
)

func init() {
	configs.ConnectToDB()
}

// func main() {

// 	r := gin.Default()
// 	r.Use(middlewares.SecretMiddleware())

// 	routes.PersonRouter(r)
// 	routes.PatientRouter(r)

// 	r.GET("/", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "Hello world from server Go.",
// 		})
// 	})
// 	r.Run()
// }

const address = "localhost:8080"

func main() {
	mux := http.NewServeMux()
	path, handler := patientsv1connect.NewPatientServiceHandler(&patientServiceServer{})
	mux.Handle(path, handler)
	fmt.Println("... Listening on", address)

	http.ListenAndServe(
		address,
		h2c.NewHandler(mux, &http2.Server{}),
	)
}

type patientServiceServer struct {
	patientsv1connect.UnimplementedPatientServiceHandler
}

func (s *patientServiceServer) GetPatientByID(
	c context.Context,
	req *connect.Request[patientsv1.PatientRequest],
) (*connect.Response[patientsv1.PatientResponse], error) {
	id := req.Msg.GetPatientId()
	var patient models.Patient
	result := configs.DB.First(&patient, id)

	if result.Error != nil {
		return nil, result.Error
	}
	fmt.Println("HEREEEEEE!!!")

	return connect.NewResponse(&patientsv1.PatientResponse{
		Patient: &patientsv1.Patient{
			Id:        patient.ID,
			FirstName: patient.FirstName,
			LastName:  patient.LastName,
			Phone:     patient.Phone,
			Ssn:       patient.SSN,
			CreatedAt: timestamppb.New(patient.CreatedAt),
			UpdatedAt: timestamppb.New(patient.UpdatedAt),
		},
	}), nil
}
