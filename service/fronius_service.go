package service

import (
	pb "github.com/avegao/iot-fronius/resource/grpc"
	"context"
	"github.com/avegao/iot-fronius/entity/fronius/current_powerflow"
	"github.com/avegao/iot-fronius/entity/fronius"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"github.com/avegao/iot-fronius/entity/fronius/current_data/inverter"
	"time"
)

type Fronius struct {
	pb.FroniusServer
}

func (service Fronius) InsertCurrentDataPowerflow(ctx context.Context, request *pb.Powerflow) (*pb.SuccessResponse, error) {
	powerflow := powerflowFromRequestToEntity(request)

	if err := powerflow.Persist(); err != nil {
		return nil, status.New(codes.Internal, err.Error()).Err()
	}

	return &pb.SuccessResponse{Success: true}, nil
}

func (service Fronius) InsertCurrentDataInverter(ctx context.Context, request *pb.CurrenDataInverterRequest) (*pb.SuccessResponse, error) {
	inverters := currentDataInverterToEntity(request)

	for _, inverter := range inverters {
		if err := inverter.Persist(); err != nil {
			return nil, status.New(codes.Internal, err.Error()).Err()
		}
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

	entity.Inverters = make([]froniusCurrentPowerflow.Inverter, len(request.GetInverter()))
	entity.Ohmpilots = make([]froniusCurrentPowerflow.Ohmpilot, len(request.GetOhmpilot()))

	for index, requestInverter := range request.GetInverter() {
		entity.Inverters[index] = froniusCurrentPowerflow.Inverter{
			BatteryMode: fronius.BatteryMode(requestInverter.GetBatteryMode()),
			DeviceType:  uint16(requestInverter.GetDeviceType()),
			EnergyDay:   requestInverter.GetEnergyDay(),
			EnergyYear:  requestInverter.GetEnergyYear(),
			EnergyTotal: requestInverter.GetEnergyTotal(),
			Soc:         uint8(requestInverter.GetSoc()),
		}
	}

	for index, requestOhmpilot := range request.GetOhmpilot() {
		entity.Ohmpilots[index] = froniusCurrentPowerflow.Ohmpilot{
			PowerAcTotal: requestOhmpilot.GetPowerAcTotal(),
			State: fronius.OhmpilotState(requestOhmpilot.GetState()),
			Temperature: requestOhmpilot.GetTemperature(),
		}
	}

	return
}

func currentDataInverterToEntity(request *pb.CurrenDataInverterRequest) []froniusCurrentDataInverter.CurrentDataInverter {
	inverters := make([]froniusCurrentDataInverter.CurrentDataInverter, 0)

	for index, requestData := range request.GetDayEnergy() {
		inverter := froniusCurrentDataInverter.CurrentDataInverter{
			DayEnergy: requestData,
			Pac: request.GetPac()[index],
			YearEnergy: request.GetYearEnergy()[index],
			TotalEnergy: request.GetTotalEnergy()[index],
			Timestamp: time.Unix(request.GetTimestamp(), 0),
			CreatedAt: time.Now(),
		}

		inverters = append(inverters, inverter)
	}

	return inverters
}
