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
	"github.com/avegao/iot-fronius/entity/fronius/current_io_state"
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

func (service Fronius) InsertCurrentIoState(ctx context.Context, request *pb.CurrentIoState) (*pb.SuccessResponse, error) {
	states := currentIoStateToEntity(request)

	for _, state := range states {
		if err := state.Persist(); err != nil {
			return nil, status.New(codes.Internal, err.Error()).Err()
		}
	}

	return &pb.SuccessResponse{Success: true}, nil
}

func (service Fronius) FindDataMeterByDate(ctx context.Context, request *pb.FindByDateParameters) (*pb.GetDataMeterResponseArray, error) {
	startDate := time.Unix(request.StartDate, 0)
	endDate := time.Unix(request.EndDate, 0)

	results, err := froniusCurrentDataMeter.FindByDate(startDate, endDate)

	if err != nil {
		return nil, status.New(codes.Internal, err.Error()).Err()
	}

	response := make([]*pb.GetDataMeterResponse, len(results))

	for index, result := range results {
		response[index] = result.ToGrpc()
	}

	return &pb.GetDataMeterResponseArray{Data: response}, nil
}

func currentDataMeterToEntity(request *pb.CurrenDataMeterRequest) []froniusCurrentDataMeter.CurrentDataMeter {
	length := len(request.GetElements())
	currentDataMeters := make([]froniusCurrentDataMeter.CurrentDataMeter, length)

	for index, requestElement := range request.GetElements() {
		currentDataMeters[index] = froniusCurrentDataMeter.CurrentDataMeter{
			CurrentAcPhase1:                   requestElement.GetCurrentAcPhase1(),
			CurrentAcPhase2:                   requestElement.GetCurrentAcPhase2(),
			CurrentAcPhase3:                   requestElement.GetCurrentAcPhase3(),
			CurrentAcSum:                      requestElement.GetCurrentAcSum(),
			Enable:                            requestElement.GetEnable(),
			EnergyReactiveVArAcPhase1Consumed: requestElement.GetEnergyReactiveVArAcPhase1Consumed(),
			EnergyReactiveVArAcPhase1Produced: requestElement.GetEnergyReactiveVArAcPhase1Produced(),
			EnergyReactiveVArAcPhase2Consumed: requestElement.GetEnergyReactiveVArAcPhase2Consumed(),
			EnergyReactiveVArAcPhase2Produced: requestElement.GetEnergyReactiveVArAcPhase2Produced(),
			EnergyReactiveVArAcPhase3Consumed: requestElement.GetEnergyReactiveVArAcPhase3Consumed(),
			EnergyReactiveVArAcPhase3Produced: requestElement.GetEnergyReactiveVArAcPhase3Produced(),
			EnergyReactiveVArAcSumConsumed:    requestElement.GetEnergyReactiveVArAcSumConsumed(),
			EnergyReactiveVArAcSumProduced:    requestElement.GetEnergyReactiveVArAcSumProduced(),
			EnergyRealWAcMinusAbsolute:        requestElement.GetEnergyRealWAcMinusAbsolute(),
			EnergyRealWAcPhase1Consumed:       requestElement.GetEnergyRealWAcPhase1Consumed(),
			EnergyRealWAcPhase1Produced:       requestElement.GetEnergyRealWAcPhase1Produced(),
			EnergyRealWAcPhase2Consumed:       requestElement.GetEnergyRealWAcPhase2Consumed(),
			EnergyRealWAcPhase2Produced:       requestElement.GetEnergyRealWAcPhase2Produced(),
			EnergyRealWAcPhase3Consumed:       requestElement.GetEnergyRealWAcPhase3Consumed(),
			EnergyRealWAcPhase3Produced:       requestElement.GetEnergyRealWAcPhase3Produced(),
			EnergyRealWAcPlusAbsolute:         requestElement.GetEnergyRealWAcPlusAbsolute(),
			EnergyRealWAcSumConsumed:          requestElement.GetEnergyRealWAcSumConsumed(),
			EnergyRealWAcSumProduced:          requestElement.GetEnergyRealWAcSumProduced(),
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
			Timestamp:                         time.Unix(requestElement.GetTimestamp(), 0),
			VoltageAcPhase1:                   requestElement.GetVoltageAcPhase1(),
			VoltageAcPhase2:                   requestElement.GetVoltageAcPhase2(),
			VoltageAcPhase3:                   requestElement.GetVoltageAcPhase3(),
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

func currentIoStateToEntity(request *pb.CurrentIoState) ([]froniusCurrentIoState.CurrentIoState) {
	entities := make([]froniusCurrentIoState.CurrentIoState, len(request.GetPins()))

	for index, pin := range request.GetPins() {
		entities[index] = froniusCurrentIoState.CurrentIoState{
			PinNumber: pin.GetPinNumber(),
			Function:  pin.GetFunction(),
			Type:      pin.GetType(),
			Direction: pin.GetDirection(),
			Set:       pin.GetSet(),
		}
	}

	return entities
}
