package crawler

import (
	//"github.com/Elderly-AI/ta_eos/internal/pkg/calculations"
	pb "github.com/Elderly-AI/observer/crawler/pkg/proto/crawler"
)

type CrawlerImplementation struct {
	pb.Un
	CalculationsFacade *calculations.Facade
}

func NewCalculationsHandler(calculations *calculations.Facade) CalculationsServer {
	return CalculationsServer{
		CalculationsFacade: calculations,
	}
}