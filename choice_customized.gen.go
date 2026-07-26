package golazada

import (
	"context"
	"encoding/json"
	"fmt"
)

type ChoiceCustomizedService interface {
	// BatchDeliverJitPurchaseOrder Batch Pickup Deliver Jit Purchase Order.
	// Path: /jit/purchase_order/batch_pickup_deliver
	BatchDeliverJitPurchaseOrder(ctx context.Context) (*BatchDeliverJitPurchaseOrderResponse, error)
	// EditChoiceSkuStock batch update choice jit product stock by skuId
	// Path: /choice/stock/edit
	EditChoiceSkuStock(ctx context.Context) (*EditChoiceSkuStockResponse, error)
	// GetChoiceProductItem Get single product by ItemId or SellerSku.
	// Path: /choice/product/item/get
	GetChoiceProductItem(ctx context.Context) (*GetChoiceProductItemResponse, error)
	// GetChoiceProducts Use this API to get detailed information of the specified products.
	// Path: /choice/products/get
	GetChoiceProducts(ctx context.Context) (*GetChoiceProductsResponse, error)
	// GetChoiceSeller Get choice seller information by seller ID and site
	// Path: /choice/seller/get
	GetChoiceSeller(ctx context.Context) (*GetChoiceSellerResponse, error)
	// GetChoiceSkuItemRelationBySku get the relation between platformSku and item by sku
	// Path: /choice/sku_item_relation/get_by_sku
	GetChoiceSkuItemRelationBySku(ctx context.Context) (*GetChoiceSkuItemRelationBySkuResponse, error)
	// PackageJitPurchaseOrder Package Jit Purchase Order.
	// Path: /jit/purchase_order/package
	PackageJitPurchaseOrder(ctx context.Context) (*PackageJitPurchaseOrderResponse, error)
	// PrintJitPurchaseOrderAndItem Print Jit Purchase Order And Item.
	// Path: /jit/purchase_order/print
	PrintJitPurchaseOrderAndItem(ctx context.Context) (*PrintJitPurchaseOrderAndItemResponse, error)
	// PrintPickuoOrder Print Pickuo Order.
	// Path: /pickup_order/print
	PrintPickuoOrder(ctx context.Context) (*PrintPickuoOrderResponse, error)
	// QueryListJitPurchaseOrder Query List Jit Purchase Order.
	// Path: /jit/purchase_order/query_list
	QueryListJitPurchaseOrder(ctx context.Context) (*QueryListJitPurchaseOrderResponse, error)
	// QueryListPurchaseItem Query List Purchase Item.
	// Path: /jit/purchase_order/query_list_purchase_item
	QueryListPurchaseItem(ctx context.Context) (*QueryListPurchaseItemResponse, error)
	// QueryPickupOrder Query Pickup Order.
	// Path: /pickup_order/query
	QueryPickupOrder(ctx context.Context) (*QueryPickupOrderResponse, error)
}

type ChoiceCustomizedServiceOp[T any] struct {
	client *Client[T]
}

// BatchDeliverJitPurchaseOrder Batch Pickup Deliver Jit Purchase Order.
// Path: /jit/purchase_order/batch_pickup_deliver
func (s *ChoiceCustomizedServiceOp[T]) BatchDeliverJitPurchaseOrder(ctx context.Context) (*BatchDeliverJitPurchaseOrderResponse, error) {
	path := "/jit/purchase_order/batch_pickup_deliver"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(BatchDeliverJitPurchaseOrderResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// EditChoiceSkuStock batch update choice jit product stock by skuId
// Path: /choice/stock/edit
func (s *ChoiceCustomizedServiceOp[T]) EditChoiceSkuStock(ctx context.Context) (*EditChoiceSkuStockResponse, error) {
	path := "/choice/stock/edit"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(EditChoiceSkuStockResponse)
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

// GetChoiceProductItem Get single product by ItemId or SellerSku.
// Path: /choice/product/item/get
func (s *ChoiceCustomizedServiceOp[T]) GetChoiceProductItem(ctx context.Context) (*GetChoiceProductItemResponse, error) {
	path := "/choice/product/item/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetChoiceProductItemResponse)
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

// GetChoiceProducts Use this API to get detailed information of the specified products.
// Path: /choice/products/get
func (s *ChoiceCustomizedServiceOp[T]) GetChoiceProducts(ctx context.Context) (*GetChoiceProductsResponse, error) {
	path := "/choice/products/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetChoiceProductsResponse)
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

// GetChoiceSeller Get choice seller information by seller ID and site
// Path: /choice/seller/get
func (s *ChoiceCustomizedServiceOp[T]) GetChoiceSeller(ctx context.Context) (*GetChoiceSellerResponse, error) {
	path := "/choice/seller/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetChoiceSellerResponse)
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

// GetChoiceSkuItemRelationBySku get the relation between platformSku and item by sku
// Path: /choice/sku_item_relation/get_by_sku
func (s *ChoiceCustomizedServiceOp[T]) GetChoiceSkuItemRelationBySku(ctx context.Context) (*GetChoiceSkuItemRelationBySkuResponse, error) {
	path := "/choice/sku_item_relation/get_by_sku"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetChoiceSkuItemRelationBySkuResponse)
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

// PackageJitPurchaseOrder Package Jit Purchase Order.
// Path: /jit/purchase_order/package
func (s *ChoiceCustomizedServiceOp[T]) PackageJitPurchaseOrder(ctx context.Context) (*PackageJitPurchaseOrderResponse, error) {
	path := "/jit/purchase_order/package"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(PackageJitPurchaseOrderResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// PrintJitPurchaseOrderAndItem Print Jit Purchase Order And Item.
// Path: /jit/purchase_order/print
func (s *ChoiceCustomizedServiceOp[T]) PrintJitPurchaseOrderAndItem(ctx context.Context) (*PrintJitPurchaseOrderAndItemResponse, error) {
	path := "/jit/purchase_order/print"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(PrintJitPurchaseOrderAndItemResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// PrintPickuoOrder Print Pickuo Order.
// Path: /pickup_order/print
func (s *ChoiceCustomizedServiceOp[T]) PrintPickuoOrder(ctx context.Context) (*PrintPickuoOrderResponse, error) {
	path := "/pickup_order/print"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(PrintPickuoOrderResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// QueryListJitPurchaseOrder Query List Jit Purchase Order.
// Path: /jit/purchase_order/query_list
func (s *ChoiceCustomizedServiceOp[T]) QueryListJitPurchaseOrder(ctx context.Context) (*QueryListJitPurchaseOrderResponse, error) {
	path := "/jit/purchase_order/query_list"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(QueryListJitPurchaseOrderResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// QueryListPurchaseItem Query List Purchase Item.
// Path: /jit/purchase_order/query_list_purchase_item
func (s *ChoiceCustomizedServiceOp[T]) QueryListPurchaseItem(ctx context.Context) (*QueryListPurchaseItemResponse, error) {
	path := "/jit/purchase_order/query_list_purchase_item"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(QueryListPurchaseItemResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// QueryPickupOrder Query Pickup Order.
// Path: /pickup_order/query
func (s *ChoiceCustomizedServiceOp[T]) QueryPickupOrder(ctx context.Context) (*QueryPickupOrderResponse, error) {
	path := "/pickup_order/query"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(QueryPickupOrderResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}
