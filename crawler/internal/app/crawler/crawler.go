package crawler

import (
	pb "github.com/Elderly-AI/observer/crawler/pkg/proto/crawler"
)

type Implementation struct {
	pb.UnimplementedCrawlerServer
	//CalculationsFacade *calculations.Facade
}

func New() Implementation {
	return Implementation{
		//CalculationsFacade: calculations,
	}
}
