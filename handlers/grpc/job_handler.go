package handler_grpc

import (
	"context"
	"job-application/entity/models"
	"job-application/interfaces"
	pb "job-application/proto"
	"job-application/query"
)

type JobHandler struct {
	pb.UnimplementedJobServiceServer
	usecase interfaces.JobUsecase
}

func NewJobHandler(usecase interfaces.JobUsecase) *JobHandler {
	return &JobHandler{usecase: usecase}
}

func (h *JobHandler) CreateJob(ctx context.Context, req *pb.JobRequest) (*pb.JobResponse, error) {
	job := &models.Job{
		Name:        req.Name,
		Quota:       uint(req.Quota),
		JobPosterID: uint(req.JobPosterId),
	}
	job, err := h.usecase.CreateJob(ctx, job)
	if err != nil {
		return nil, err
	}
	return &pb.JobResponse{
		Id:         uint32(job.ID),
		Name:       job.Name,
		Quota:      uint32(job.Quota),
		ExpiryDate: job.ExpiryDate.Format("2006/01/01"),
	}, nil
}

func (h *JobHandler) FindAllJobs(ctx context.Context, req *pb.WhereClauses) (*pb.AllJobResponses, error) {
	whereClauses := []query.WhereClause{}
	for _, clause := range req.Clauses {
		whereClause := &query.WhereClause{
			Field:     clause.Field,
			Value:     clause.Value,
			Condition: clause.Condition,
		}
		whereClauses = append(whereClauses, *whereClause)
	}
	jobs, err := h.usecase.FindAllJobs(ctx, whereClauses)
	if err != nil {
		return nil, err
	}
	jobResponses := []*pb.JobResponse{}
	for _, job := range jobs {
		jobResponse := &pb.JobResponse{
			Id:          uint32(job.ID),
			Name:        job.Name,
			Quota:       uint32(job.Quota),
			ExpiryDate:  job.ExpiryDate.Format("2006/01/02"),
			JobPosterId: uint32(job.JobPosterID),
		}
		jobResponses = append(jobResponses, jobResponse)
	}
	return &pb.AllJobResponses{
		Message:   "Success get all jobs",
		Responses: jobResponses,
	}, nil
}
