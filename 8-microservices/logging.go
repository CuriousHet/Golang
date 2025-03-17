package main

import (
	"context"
	"fmt"
	"time"
)

type LoggingService struct {
	next Service
}

func NewLoggingService(next Service) Service {
	return &LoggingService{next: next}
}

func (s *LoggingService) GetCatFact(ctx context.Context) (*CatFact, error) {
	start := time.Now()
	fact, err := s.next.GetCatFact(ctx)
	fmt.Printf("fact: %s, err: %v, took: %v\n", fact.Fact, err, time.Since(start))
	return fact, err
}
