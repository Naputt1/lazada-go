package golazada

import (
	"context"
)

type RedMartService interface {
	// RssGetOnePickupJob Get details of a pickup job
	// Path: /rss/pickup-job/get
	RssGetOnePickupJob(ctx context.Context) (*RssGetOnePickupJobResponse, error)
	// RssGetPickupJobs Retrieve RSS pickup jobs based on time range and status filter.
	// Path: /rss/pickup-jobs/get
	RssGetPickupJobs(ctx context.Context) (*RssGetPickupJobsResponse, error)
	// RssGetPickupLocations rss get pickupLocations by storeId
	// Path: /rss/pickupLocations/get
	RssGetPickupLocations(ctx context.Context) (*RssGetPickupLocationsResponse, error)
	// RssGetProduct get rss product by storeId and productId
	// Path: /rss/product/get
	RssGetProduct(ctx context.Context) (*RssGetProductResponse, error)
	// RssGetProducts rss get products paged by storeId and pickupLocationId
	// Path: /rss/products/get
	RssGetProducts(ctx context.Context) (*RssGetProductsResponse, error)
	// RssGetStockLot rss get stockLot
	// Path: /rss/stockLot/get
	RssGetStockLot(ctx context.Context) (*RssGetStockLotResponse, error)
	// RssGetStockLots rss get stockLots
	// Path: /rss/stockLots/get
	RssGetStockLots(ctx context.Context) (*RssGetStockLotsResponse, error)
	// RssUpdateStockLot rss update stockLot
	// Path: /rss/stockLot/update
	RssUpdateStockLot(ctx context.Context) (*RssUpdateStockLotResponse, error)
}

type RedMartServiceOp[T any] struct {
	client *Client[T]
}

// RssGetOnePickupJob Get details of a pickup job
// Path: /rss/pickup-job/get
func (s *RedMartServiceOp[T]) RssGetOnePickupJob(ctx context.Context) (*RssGetOnePickupJobResponse, error) {
	path := "/rss/pickup-job/get"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(RssGetOnePickupJobResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// RssGetPickupJobs Retrieve RSS pickup jobs based on time range and status filter.
// Path: /rss/pickup-jobs/get
func (s *RedMartServiceOp[T]) RssGetPickupJobs(ctx context.Context) (*RssGetPickupJobsResponse, error) {
	path := "/rss/pickup-jobs/get"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(RssGetPickupJobsResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// RssGetPickupLocations rss get pickupLocations by storeId
// Path: /rss/pickupLocations/get
func (s *RedMartServiceOp[T]) RssGetPickupLocations(ctx context.Context) (*RssGetPickupLocationsResponse, error) {
	path := "/rss/pickupLocations/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(RssGetPickupLocationsResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// RssGetProduct get rss product by storeId and productId
// Path: /rss/product/get
func (s *RedMartServiceOp[T]) RssGetProduct(ctx context.Context) (*RssGetProductResponse, error) {
	path := "/rss/product/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(RssGetProductResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// RssGetProducts rss get products paged by storeId and pickupLocationId
// Path: /rss/products/get
func (s *RedMartServiceOp[T]) RssGetProducts(ctx context.Context) (*RssGetProductsResponse, error) {
	path := "/rss/products/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(RssGetProductsResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// RssGetStockLot rss get stockLot
// Path: /rss/stockLot/get
func (s *RedMartServiceOp[T]) RssGetStockLot(ctx context.Context) (*RssGetStockLotResponse, error) {
	path := "/rss/stockLot/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(RssGetStockLotResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// RssGetStockLots rss get stockLots
// Path: /rss/stockLots/get
func (s *RedMartServiceOp[T]) RssGetStockLots(ctx context.Context) (*RssGetStockLotsResponse, error) {
	path := "/rss/stockLots/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(RssGetStockLotsResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// RssUpdateStockLot rss update stockLot
// Path: /rss/stockLot/update
func (s *RedMartServiceOp[T]) RssUpdateStockLot(ctx context.Context) (*RssUpdateStockLotResponse, error) {
	path := "/rss/stockLot/update"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(RssUpdateStockLotResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}
