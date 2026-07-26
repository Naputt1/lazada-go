package golazada

import (
	"context"
	"encoding/json"
	"fmt"
)

type CrossBoarderProductService interface {
	// CreateGlobalProduct Use this API to create a single new global product to multiple Lazada sites. (For cross boarder sellers ONLY)
	// Path: /product/global/create
	CreateGlobalProduct(ctx context.Context) (*CreateGlobalProductResponse, error)
	// DeleteMerchantProduct Use this API to delete the product。(CrossBoarderSellersOnly)
	// Path: /product/global/delete
	DeleteMerchantProduct(ctx context.Context) (*DeleteMerchantProductResponse, error)
	// GetGlobalProductExtension Use this API to query the extension info of the specified global product. (CrossBoarderSellersOnly)
	// Path: /product/global/extension
	GetGlobalProductExtension(ctx context.Context) (*GetGlobalProductExtensionResponse, error)
	// GetGlobalProductStatus Use this API to query the status of the specified global product. It takes several minutes for the global product to be created on each site. (CrossBoarderSellersOnly)
	// Path: /product/global/status/get
	GetGlobalProductStatus(ctx context.Context) (*GetGlobalProductStatusResponse, error)
	// GetRecommendPrice get recommend price
	// Path: /product/global/semi/recommend/price/get
	GetRecommendPrice(ctx context.Context) (*GetRecommendPriceResponse, error)
	// GetUnfilledAttribute get the product which have attribute not filled （for cross boarder sellers Only）
	// Path: /product/global/unfilled/attribute/get
	GetUnfilledAttribute(ctx context.Context) (*GetUnfilledAttributeResponse, error)
	// GetUpgradableGlobalPlusProductList get an upgradeable global plus product list
	// Path: /product/global/semi/avaible/get
	GetUpgradableGlobalPlusProductList(ctx context.Context) (*GetUpgradableGlobalPlusProductListResponse, error)
	// SemiProductUpdate SemiProductUpdate
	// Path: /product/global/semi/update
	SemiProductUpdate(ctx context.Context) (*SemiProductUpdateResponse, error)
	// SemiProductUpgrade SemiProductUpgrade
	// Path: /product/global/semi/upgrade
	SemiProductUpgrade(ctx context.Context) (*SemiProductUpgradeResponse, error)
	// UpdateGlobalProductAttribute update global product attribute (For cross boarder sellers only)
	// Path: /product/global/attribute/update
	UpdateGlobalProductAttribute(ctx context.Context) (*UpdateGlobalProductAttributeResponse, error)
	// UpdateProductStatus product up shelf or down shelf，(CrossBoarderSellersOnly)
	// Path: /product/global/update/status
	UpdateProductStatus(ctx context.Context) (*UpdateProductStatusResponse, error)
}

type CrossBoarderProductServiceOp[T any] struct {
	client *Client[T]
}

// CreateGlobalProduct Use this API to create a single new global product to multiple Lazada sites. (For cross boarder sellers ONLY)
// Path: /product/global/create
func (s *CrossBoarderProductServiceOp[T]) CreateGlobalProduct(ctx context.Context) (*CreateGlobalProductResponse, error) {
	path := "/product/global/create"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(CreateGlobalProductResponse)
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

// DeleteMerchantProduct Use this API to delete the product。(CrossBoarderSellersOnly)
// Path: /product/global/delete
func (s *CrossBoarderProductServiceOp[T]) DeleteMerchantProduct(ctx context.Context) (*DeleteMerchantProductResponse, error) {
	path := "/product/global/delete"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(DeleteMerchantProductResponse)
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

// GetGlobalProductExtension Use this API to query the extension info of the specified global product. (CrossBoarderSellersOnly)
// Path: /product/global/extension
func (s *CrossBoarderProductServiceOp[T]) GetGlobalProductExtension(ctx context.Context) (*GetGlobalProductExtensionResponse, error) {
	path := "/product/global/extension"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetGlobalProductExtensionResponse)
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

// GetGlobalProductStatus Use this API to query the status of the specified global product. It takes several minutes for the global product to be created on each site. (CrossBoarderSellersOnly)
// Path: /product/global/status/get
func (s *CrossBoarderProductServiceOp[T]) GetGlobalProductStatus(ctx context.Context) (*GetGlobalProductStatusResponse, error) {
	path := "/product/global/status/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetGlobalProductStatusResponse)
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

// GetRecommendPrice get recommend price
// Path: /product/global/semi/recommend/price/get
func (s *CrossBoarderProductServiceOp[T]) GetRecommendPrice(ctx context.Context) (*GetRecommendPriceResponse, error) {
	path := "/product/global/semi/recommend/price/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetRecommendPriceResponse)
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

// GetUnfilledAttribute get the product which have attribute not filled （for cross boarder sellers Only）
// Path: /product/global/unfilled/attribute/get
func (s *CrossBoarderProductServiceOp[T]) GetUnfilledAttribute(ctx context.Context) (*GetUnfilledAttributeResponse, error) {
	path := "/product/global/unfilled/attribute/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetUnfilledAttributeResponse)
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

// GetUpgradableGlobalPlusProductList get an upgradeable global plus product list
// Path: /product/global/semi/avaible/get
func (s *CrossBoarderProductServiceOp[T]) GetUpgradableGlobalPlusProductList(ctx context.Context) (*GetUpgradableGlobalPlusProductListResponse, error) {
	path := "/product/global/semi/avaible/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetUpgradableGlobalPlusProductListResponse)
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

// SemiProductUpdate SemiProductUpdate
// Path: /product/global/semi/update
func (s *CrossBoarderProductServiceOp[T]) SemiProductUpdate(ctx context.Context) (*SemiProductUpdateResponse, error) {
	path := "/product/global/semi/update"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(SemiProductUpdateResponse)
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

// SemiProductUpgrade SemiProductUpgrade
// Path: /product/global/semi/upgrade
func (s *CrossBoarderProductServiceOp[T]) SemiProductUpgrade(ctx context.Context) (*SemiProductUpgradeResponse, error) {
	path := "/product/global/semi/upgrade"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(SemiProductUpgradeResponse)
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

// UpdateGlobalProductAttribute update global product attribute (For cross boarder sellers only)
// Path: /product/global/attribute/update
func (s *CrossBoarderProductServiceOp[T]) UpdateGlobalProductAttribute(ctx context.Context) (*UpdateGlobalProductAttributeResponse, error) {
	path := "/product/global/attribute/update"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(UpdateGlobalProductAttributeResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// UpdateProductStatus product up shelf or down shelf，(CrossBoarderSellersOnly)
// Path: /product/global/update/status
func (s *CrossBoarderProductServiceOp[T]) UpdateProductStatus(ctx context.Context) (*UpdateProductStatusResponse, error) {
	path := "/product/global/update/status"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(UpdateProductStatusResponse)
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
