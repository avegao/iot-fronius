package service

import (
	pb "github.com/avegao/iot-fronius/resource/grpc"
	"context"
	"github.com/avegao/iot-fronius/entity/fronius/current_powerflow"
	"github.com/avegao/iot-fronius/entity/fronius"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Fronius struct {
	pb.FroniusServer
}

func (froniusService Fronius) InsertCurrentDataPowerflow(ctx context.Context, request *pb.Powerflow) (*pb.SuccessResponse, error) {
	powerflow := powerflowFromRequestToEntity(request)
	err := powerflow.Persist()

	if err != nil {
		return nil, status.New(codes.Internal, err.Error()).Err()
	}

	return &pb.SuccessResponse{Success: true}, nil
}

func powerflowFromRequestToEntity(request *pb.Powerflow) (entity froniusCurrentPowerflow.CurrentPowerflow) {
	entity.Site = froniusCurrentPowerflow.Site{
		Mode:                    fronius.NewSiteModeFromString(request.Site.Mode),
		BatteryStandby:          request.Site.GetBatteryStandby(),
		BackupMode:              request.Site.GetBackupMode(),
		PowerFromGrid:           request.Site.GetPowerFromGrid(),
		PowerLoad:               request.Site.GetPowerLoad(),
		PowerAkku:               request.Site.GetPowerAkku(),
		PowerFromPV:             request.Site.GetPowerFromPv(),
		RelativeSelfConsumption: uint8(request.Site.GetRelativeSelfConsumption()),
		RelativeAutonomy:        uint8(request.Site.GetRelativeAutonomy()),
		MeterLocation:           fronius.NewMeterLocationFromString(request.Site.GetMeterLocation()),
		EnergyDay:               request.Site.GetEnergyDay(),
		EnergyYear:              request.Site.GetEnergyYear(),
		EnergyTotal:             request.Site.GetEnergyTotal(),
	}

	return
}
