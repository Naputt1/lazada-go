package golazada

import (
	"context"
	"encoding/json"
	"fmt"
)

type LogisticsService interface {
	// AddOrUpdatePickupStop 3PL call TPS to update pickup stops
	// Path: /logistics/tps/runsheets/stops
	AddOrUpdatePickupStop(ctx context.Context) (*AddOrUpdatePickupStopResponse, error)
	// Create3PLStation TPS_CREATE_STATION_API
	// External partner call TPS to create station
	// Path: /logistics/tps/stations/create
	Create3PLStation(ctx context.Context) (*Create3PLStationResponse, error)
	// CreateConsolidationService create Consolidation Service
	// Path: /logistics/ldp/createConsolidationService
	CreateConsolidationService(ctx context.Context) (*CreateConsolidationServiceResponse, error)
	// GetOrderTrace Query logistic detail for seller erp with seller id, order id and locale info. This api is only available in the state after ready to ship.
	// Path: /logistic/order/trace
	GetOrderTrace(ctx context.Context) (*GetOrderTraceResponse, error)
	// ScanParcel DOP Scan Parcel
	// Path: /dop/scan
	ScanParcel(ctx context.Context) (*ScanParcelResponse, error)
	// StationDopScan StationDopScan
	// Path: /stations/dop/scan
	StationDopScan(ctx context.Context) (*StationDopScanResponse, error)
	// Update3PLStation TPS_UPDATE_STATION_API
	// External partner call TPS to update station
	// Path: /logistics/tps/stations/update
	Update3PLStation(ctx context.Context) (*Update3PLStationResponse, error)
	// UpdateLastMile 跨境场景，物流末端预报信息
	// Path: /logistics/ldp/updateLastmile
	UpdateLastMile(ctx context.Context) (*UpdateLastMileResponse, error)
	// UpdatePickupTimeSlot 3PL call TPS to update pickup timeslot
	// Path: /logistics/tps/sellers/pickup_timeslot
	UpdatePickupTimeSlot(ctx context.Context) (*UpdatePickupTimeSlotResponse, error)
}

type LogisticsServiceOp[T any] struct {
	client *Client[T]
}

// AddOrUpdatePickupStop 3PL call TPS to update pickup stops
// Path: /logistics/tps/runsheets/stops
func (s *LogisticsServiceOp[T]) AddOrUpdatePickupStop(ctx context.Context) (*AddOrUpdatePickupStopResponse, error) {
	path := "/logistics/tps/runsheets/stops"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(AddOrUpdatePickupStopResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// Create3PLStation TPS_CREATE_STATION_API
// External partner call TPS to create station
// Path: /logistics/tps/stations/create
func (s *LogisticsServiceOp[T]) Create3PLStation(ctx context.Context) (*Create3PLStationResponse, error) {
	path := "/logistics/tps/stations/create"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(Create3PLStationResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// CreateConsolidationService create Consolidation Service
// Path: /logistics/ldp/createConsolidationService
func (s *LogisticsServiceOp[T]) CreateConsolidationService(ctx context.Context) (*CreateConsolidationServiceResponse, error) {
	path := "/logistics/ldp/createConsolidationService"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(CreateConsolidationServiceResponse)
	if string(wrapper.Data) != "null" && len(wrapper.Data) > 0 {
		if err := json.Unmarshal(wrapper.Data, &resp.Response); err != nil {
			return nil, fmt.Errorf("failed to decode response: %w", err)
		}
	}
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetOrderTrace Query logistic detail for seller erp with seller id, order id and locale info. This api is only available in the state after ready to ship.
// Path: /logistic/order/trace
func (s *LogisticsServiceOp[T]) GetOrderTrace(ctx context.Context) (*GetOrderTraceResponse, error) {
	path := "/logistic/order/trace"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetOrderTraceResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// ScanParcel DOP Scan Parcel
// Path: /dop/scan
func (s *LogisticsServiceOp[T]) ScanParcel(ctx context.Context) (*ScanParcelResponse, error) {
	path := "/dop/scan"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(ScanParcelResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// StationDopScan StationDopScan
// Path: /stations/dop/scan
func (s *LogisticsServiceOp[T]) StationDopScan(ctx context.Context) (*StationDopScanResponse, error) {
	path := "/stations/dop/scan"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(StationDopScanResponse)
	if string(wrapper.Data) != "null" && len(wrapper.Data) > 0 {
		if err := json.Unmarshal(wrapper.Data, &resp.Response); err != nil {
			return nil, fmt.Errorf("failed to decode response: %w", err)
		}
	}
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// Update3PLStation TPS_UPDATE_STATION_API
// External partner call TPS to update station
// Path: /logistics/tps/stations/update
func (s *LogisticsServiceOp[T]) Update3PLStation(ctx context.Context) (*Update3PLStationResponse, error) {
	path := "/logistics/tps/stations/update"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(Update3PLStationResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// UpdateLastMile 跨境场景，物流末端预报信息
// Path: /logistics/ldp/updateLastmile
func (s *LogisticsServiceOp[T]) UpdateLastMile(ctx context.Context) (*UpdateLastMileResponse, error) {
	path := "/logistics/ldp/updateLastmile"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(UpdateLastMileResponse)
	if string(wrapper.Data) != "null" && len(wrapper.Data) > 0 {
		if err := json.Unmarshal(wrapper.Data, &resp.Response); err != nil {
			return nil, fmt.Errorf("failed to decode response: %w", err)
		}
	}
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// UpdatePickupTimeSlot 3PL call TPS to update pickup timeslot
// Path: /logistics/tps/sellers/pickup_timeslot
func (s *LogisticsServiceOp[T]) UpdatePickupTimeSlot(ctx context.Context) (*UpdatePickupTimeSlotResponse, error) {
	path := "/logistics/tps/sellers/pickup_timeslot"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(UpdatePickupTimeSlotResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}
