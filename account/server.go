package account

import (
	"context"
	"fmt"
)

type grpcServer struct{
	service Service
}

func ListenGRPC(s Service, port int) error{
	lis, err : net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	serv := grpc.NewServer()
	pb.(serv,)
	reflection.Register(serv)
	return serv.Serve(lis)
}

func(s *grpcServer) PostAccount(ctx context.Context, r *pb.)(*pb.,error){
	a, err := s.service.PostAccount(ctx, r.name)
	if err != nil{
		return nil, err
	}
	return &pb.{}
}

func (s *grpcServer) GetAccount(ctx context.Context, r *pb.)(*pb.,error){
	a, err := s.service.GetAccount(ctx, r.Id)
	if err != nil{
		return nil, err
	}
	return &pb.{}
}


func (s *grpcServer) GetAccounts(ctx context.Context, r *pb.)(*pb.,error){
	a, err := s.service.GetAccounts(ctx, r.Id)
	if err != nil{
		return nil, err
	}
	return &pb.{}
}