package crawler

import (
	//"github.com/Elderly-AI/ta_eos/internal/pkg/calculations"
	pb "github.com/"
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
