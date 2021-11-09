package main

import (
	"context"

	pb "github.com/lreimer/from-rest-to-grpc/grpc-beer-service/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	pb.UnimplementedBeerServiceServer
}

// Get the list of all beers
func (s *Server) AllBeers(ctx context.Context, r *emptypb.Empty) (*pb.GetBeersResponse, error) {
	return &pb.GetBeersResponse{Beers: allBeers()}, nil
}

// Get a single beer by Asin
func (s *Server) GetBeer(ctx context.Context, r *pb.GetBeerRequest) (*pb.GetBeerResponse, error) {
	beer, found := getBeer(r.Asin)
	if found {
		return &pb.GetBeerResponse{Beer: beer}, nil
	} else {
		return nil, status.Error(codes.NotFound, "Beer not found")
	}
}

// Create a new beer
func (s *Server) CreateBeer(ctx context.Context, r *pb.CreateBeerRequest) (*emptypb.Empty, error) {
	_, created := createBeer(*r.Beer)
	if created {
		return &emptypb.Empty{}, nil
	} else {
		return nil, status.Error(codes.AlreadyExists, "Beer already exists")
	}
}

// Update an existing beer
func (s *Server) UpdateBeer(ctx context.Context, r *pb.UpdateBeerRequest) (*emptypb.Empty, error) {
	updated := updateBeer(r.Asin, *r.Beer)
	if updated {
		return &emptypb.Empty{}, nil
	} else {
		return nil, status.Error(codes.NotFound, "Beer not updated")
	}
}

// Delete an existing beeer
func (s *Server) DeleteBeer(ctx context.Context, r *pb.DeleteBeerRequest) (*emptypb.Empty, error) {
	deleteBeer(r.Asin)
	return &emptypb.Empty{}, nil
}
