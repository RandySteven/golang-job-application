package handler_grpc

import (
	"context"
	"job-application/entity/models"
	"job-application/interfaces"
	pb "job-application/proto"
	"job-application/query"
)

type ApplicationHandler struct {
	pb.UnimplementedApplicationServiceServer
	usecase interfaces.ApplicationUsecase
}

func (h *ApplicationHandler) ApplyApplication(ctx context.Context, req *pb.ApplicationRequest) (*pb.ApplicationResponse, error) {
	application := &models.Application{
		UserID: uint(req.UserId),
		JobID:  uint(req.JobId),
	}
	application, err := h.usecase.ApplyApplication(ctx, application)
	if err != nil {
		return nil, err
	}
	return &pb.ApplicationResponse{
		Id:     uint32(application.ID),
		UserId: uint32(application.JobID),
		JobId:  uint32(application.JobID),
		Status: application.Status,
	}, nil
}

func (h *ApplicationHandler) GetAllApplications(ctx context.Context, req *pb.WhereClauses) (*pb.AllApplicationResponses, error) {
	whereClauses := []query.WhereClause{}
	for _, clause := range req.Clauses {
		whereClause := &query.WhereClause{
			Field:     clause.Field,
			Value:     clause.Value,
			Condition: clause.Condition,
		}
		whereClauses = append(whereClauses, *whereClause)
	}
	applications, err := h.usecase.GetAllApplications(ctx, whereClauses)
	if err != nil {
		return nil, err
	}
	appResponses := []*pb.ApplicationResponse{}
	for _, application := range applications {
		resp := &pb.ApplicationResponse{
			Id:     uint32(application.ID),
			UserId: uint32(application.JobID),
			JobId:  uint32(application.JobID),
			Status: application.Status,
		}
		appResponses = append(appResponses, resp)
	}
	return &pb.AllApplicationResponses{
		Message:   "Success get all applications",
		Responses: appResponses,
	}, nil
}

func NewApplicationHandler(usecase interfaces.ApplicationUsecase) *ApplicationHandler {
	return &ApplicationHandler{usecase: usecase}
}
