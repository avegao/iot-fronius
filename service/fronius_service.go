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
	"github.com/avegao/iot-fronius/entity/fronius/current_data/meter"
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

func (service Fronius) InsertCurrentDataMeter(ctx context.Context, request *pb.CurrenDataMeterRequest) (*pb.SuccessResponse, error) {
	inverters := currentDataMeterToEntity(request)

	for _, inverter := range inverters {
		if err := inverter.Persist(); err != nil {
			return nil, status.New(codes.Internal, err.Error()).Err()
		}
	}

	return &pb.SuccessResponse{Success: true}, nil
}
func currentDataMeterToEntity(request *pb.CurrenDataMeterRequest) []froniusCurrentDataMeter.CurrentDataMeter {
	length := len(request.GetElements())
	currentDataMeters := make([]froniusCurrentDataMeter.CurrentDataMeter, length)

	for index, requestElement := range request.GetElements() {
		currentDataMeters[index] = froniusCurrentDataMeter.CurrentDataMeter{
			CurrentACPhase1:                   requestElement.GetCurrentAcPhase1(),
			CurrentACPhase2:                   requestElement.GetCurrentAcPhase2(),
			CurrentACPhase3:                   requestElement.GetCurrentAcPhase3(),
			CurrentACSum:                      requestElement.GetCurrentAcSum(),
			Enable:                            requestElement.GetEnable(),
			EnergyReactiveVArACPhase1Consumed: requestElement.GetEnergyReactiveVArAcPhase1Consumed(),
			EnergyReactiveVArACPhase1Produced: requestElement.GetEnergyReactiveVArAcPhase1Produced(),
			EnergyReactiveVArACPhase2Consumed: requestElement.GetEnergyReactiveVArAcPhase2Consumed(),
			EnergyReactiveVArACPhase2Produced: requestElement.GetEnergyReactiveVArAcPhase2Produced(),
			EnergyReactiveVArACPhase3Consumed: requestElement.GetEnergyReactiveVArAcPhase3Consumed(),
			EnergyReactiveVArACPhase3Produced: requestElement.GetEnergyReactiveVArAcPhase3Produced(),
			EnergyReactiveVArACSumConsumed:    requestElement.GetEnergyReactiveVArAcSumConsumed(),
			EnergyReactiveVArACSumProduced:    requestElement.GetEnergyReactiveVArAcSumProduced(),
			EnergyRealWACMinusAbsolute:        requestElement.GetEnergyRealWAcMinusAbsolute(),
			EnergyRealWACPhase1Consumed:       requestElement.GetEnergyRealWAcPhase1Consumed(),
			EnergyRealWACPhase1Produced:       requestElement.GetEnergyRealWAcPhase1Produced(),
			EnergyRealWACPhase2Consumed:       requestElement.GetEnergyRealWAcPhase2Consumed(),
			EnergyRealWACPhase2Produced:       requestElement.GetEnergyRealWAcPhase2Produced(),
			EnergyRealWACPhase3Consumed:       requestElement.GetEnergyRealWAcPhase3Consumed(),
			EnergyRealWACPhase3Produced:       requestElement.GetEnergyRealWAcPhase3Produced(),
			EnergyRealWACPlusAbsolute:         requestElement.GetEnergyRealWAcPlusAbsolute(),
			EnergyRealWACSumConsumed:          requestElement.GetEnergyRealWAcSumConsumed(),
			EnergyRealWACSumProduced:          requestElement.GetEnergyRealWAcSumProduced(),
			FrequencyPhaseAverage:             requestElement.GetFrequencyPhaseAverage(),
			MeterLocationCurrent:              requestElement.GetMeterLocationCurrent(),
			PowerApparentSPhase1:              requestElement.GetPowerApparentSPhase1(),
			PowerApparentSPhase2:              requestElement.GetPowerApparentSPhase2(),
			PowerApparentSPhase3:              requestElement.GetPowerApparentSPhase3(),
			PowerApparentSSum:                 requestElement.GetPowerApparentSSum(),
			PowerFactorPhase1:                 requestElement.GetPowerFactorPhase1(),
			PowerFactorPhase2:                 requestElement.GetPowerFactorPhase2(),
			PowerFactorPhase3:                 requestElement.GetPowerFactorPhase3(),
			PowerFactorSum:                    requestElement.GetPowerFactorSum(),
			PowerReactiveQPhase1:              requestElement.GetPowerReactiveQPhase1(),
			PowerReactiveQPhase2:              requestElement.GetPowerReactiveQPhase2(),
			PowerReactiveQPhase3:              requestElement.GetPowerReactiveQPhase3(),
			PowerReactiveQSum:                 requestElement.GetPowerReactiveQSum(),
			PowerRealPPhase1:                  requestElement.GetPowerRealPPhase1(),
			PowerRealPPhase2:                  requestElement.GetPowerRealPPhase2(),
			PowerRealPPhase3:                  requestElement.GetPowerRealPPhase3(),
			PowerRealPSum:                     requestElement.GetPowerRealPSum(),
			TimeStamp:                         requestElement.GetTimestamp(),
			Visible:                           requestElement.GetVisible(),
			VoltageACPhase1:                   requestElement.GetVoltageAcPhase1(),
			VoltageACPhase2:                   requestElement.GetVoltageAcPhase2(),
			VoltageACPhase3:                   requestElement.GetVoltageAcPhase3(),
		}
	}

	return currentDataMeters
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
			State:        fronius.OhmpilotState(requestOhmpilot.GetState()),
			Temperature:  requestOhmpilot.GetTemperature(),
		}
	}

	return
}

func currentDataInverterToEntity(request *pb.CurrenDataInverterRequest) []froniusCurrentDataInverter.CurrentDataInverter {
	inverters := make([]froniusCurrentDataInverter.CurrentDataInverter, 0)

	for index, requestData := range request.GetDayEnergy() {
		inverter := froniusCurrentDataInverter.CurrentDataInverter{
			DayEnergy:   requestData,
			Pac:         request.GetPac()[index],
			YearEnergy:  request.GetYearEnergy()[index],
			TotalEnergy: request.GetTotalEnergy()[index],
			Timestamp:   time.Unix(request.GetTimestamp(), 0),
			CreatedAt:   time.Now(),
		}

		inverters = append(inverters, inverter)
	}

	return inverters
}
