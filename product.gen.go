package golazada

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
)

type ProductService interface {
	// AdjustSellableQuantity Use this API to increase or decrease sellable quantity of one or more existing products. The maximum number of products that can be updated is 50, but 20 is recommended.
	// Path: /product/stock/sellable/adjust
	AdjustSellableQuantity(ctx context.Context) (*AdjustSellableQuantityResponse, error)
	// BatchUpdateSizeChart 批量更新尺码表
	// Path: /size/chart/batch/update
	BatchUpdateSizeChart(ctx context.Context) (*BatchUpdateSizeChartResponse, error)
	// CreateProduct Use this API to create a single new product.
	//
	// Find more details below: https://open.lazada.com/apps/doc/doc?nodeId=30720&docId=120949
	// Path: /product/create
	CreateProduct(ctx context.Context, req CreateProductRequest) (*CreateProductResponse, error)
	// DeactivateProduct Use this API to deactivate Product or SKUs corresponding to the product
	// Path: /product/deactivate
	DeactivateProduct(ctx context.Context) (*DeactivateProductResponse, error)
	// GetBrandByPages Use this API to retrieve all product brands by page index in the system.
	// Path: /category/brands/query
	GetBrandByPages(ctx context.Context) (*GetBrandByPagesResponse, error)
	// GetCategoryAttributes Use this API to get a list of attributes for a specified product category.
	// Path: /category/attributes/get
	GetCategoryAttributes(ctx context.Context) (*GetCategoryAttributesResponse, error)
	// GetCategorySuggestion Get product's category suggestion by product title
	// Path: /product/category/suggestion/get
	GetCategorySuggestion(ctx context.Context) (*GetCategorySuggestionResponse, error)
	// GetCategoryTree Use this API to retrieve the list of all product categories in the system.
	// Path: /category/tree/get
	GetCategoryTree(ctx context.Context) (*GetCategoryTreeResponse, error)
	// GetNextCascadeProp Use this API to query next cascade prop.
	// Path: /category/cascade/getNextCascadeProp
	GetNextCascadeProp(ctx context.Context) (*GetNextCascadePropResponse, error)
	// GetPreQcRules query pre qc rules
	// Path: /product/seller/item/getPreQcRules
	GetPreQcRules(ctx context.Context) (*GetPreQcRulesResponse, error)
	// GetProductContentScore get product content score
	// Path: /product/content/score/get
	GetProductContentScore(ctx context.Context) (*GetProductContentScoreResponse, error)
	// GetProductItem Get single product by ItemId or SellerSku.
	// Path: /product/item/get
	GetProductItem(ctx context.Context, opt GetProductItemRequest) (*GetProductItemResponse, error)
	// GetProducts Use this API to get detailed information of the specified products.
	// Path: /products/get
	GetProducts(ctx context.Context, opt GetProductsRequest) (*GetProductsResponse, error)
	// GetQCAlertProducts Getting seller's products that have been alerted by quality control.
	// Path: /product/qc/alert/list
	GetQCAlertProducts(ctx context.Context) (*GetQCAlertProductsResponse, error)
	// GetResponse Use this API to get the returned information from the system for the MigrateImages API.
	// Path: /image/response/get
	GetResponse(ctx context.Context, filename string, reader io.Reader) (*GetResponseResponse, error)
	// GetSellerItemLimit The platform will provide the product quantity limit information by this interface. The qps will be limited by seller, 10 qps per seller.
	// Path: /product/seller/item/limit
	GetSellerItemLimit(ctx context.Context) (*GetSellerItemLimitResponse, error)
	// GetSizeChartTemplate 获取尺码模板列表
	// Path: /size/chart/template/get
	GetSizeChartTemplate(ctx context.Context) (*GetSizeChartTemplateResponse, error)
	// GetUnfilledAttributeItem Get products without key attributes. (For cross boarder sellers Only)
	// Path: /product/unfilled/attribute/get
	GetUnfilledAttributeItem(ctx context.Context) (*GetUnfilledAttributeItemResponse, error)
	// MigrateImage Use this API to migrate a single image from an external site to Lazada site. Allowed image formats are JPG and PNG. The maximum size of an image file is 1MB.
	// Path: /image/migrate
	MigrateImage(ctx context.Context, filename string, reader io.Reader) (*MigrateImageResponse, error)
	// MigrateImages Use this API to migrate multiple images from an external site to Lazada site. Allowed image formats are JPG and PNG. The maximum size of an image file is 1MB. A single call can migrate 8 images at most.
	// Path: /images/migrate
	MigrateImages(ctx context.Context, filename string, reader io.Reader) (*MigrateImagesResponse, error)
	// ProductCheck Use this API to check CB seller quantity limit of adding product .
	// Path: /product/pre/check
	ProductCheck(ctx context.Context) (*ProductCheckResponse, error)
	// RemoveProduct Use this API to remove an existing product, some SKUs in one product, or all SKUs in one product. System supports a maximum number of 50 SellerSkus in one request.
	// Path: /product/remove
	RemoveProduct(ctx context.Context) (*RemoveProductResponse, error)
	// RemoveSku Use this API to delete SKUs and sales attributes of corresponding products.
	// Path: /product/sku/remove
	RemoveSku(ctx context.Context) (*RemoveSkuResponse, error)
	// SetImages Use this API to set the images for an existing product by associating one or more image URLs with it.
	// Path: /images/set
	SetImages(ctx context.Context, filename string, reader io.Reader) (*SetImagesResponse, error)
	// UpdatePriceQuantity Use this API to update the price and quantity of one or more existing products. The maximum number of products that can be updated is 50, but 20 is recommended.
	// Path: /product/price_quantity/update
	UpdatePriceQuantity(ctx context.Context) (*UpdatePriceQuantityResponse, error)
	// UpdateProduct Use this API to update attributes or SKUs of an existing product. if need update inventory, offline, price, not recommended to use this API.
	// The iteration 25/6/2020 Updated for DBS changes. Refer to Input Parameters Payload
	// Path: /product/update
	UpdateProduct(ctx context.Context, req UpdateProductRequest) (*UpdateProductResponse, error)
	// UpdateSellableQuantity Use this API to update sellable quantity of one or more existing products. The maximum number of products that can be updated is 50, but 20 is recommended.
	// Path: /product/stock/sellable/update
	UpdateSellableQuantity(ctx context.Context) (*UpdateSellableQuantityResponse, error)
	// UploadImage Use this API to upload a single image file to Lazada site. Allowed image formats are JPG and PNG. The maximum size of an image file is 1MB.
	// Path: /image/upload
	UploadImage(ctx context.Context, filename string, reader io.Reader) (*UploadImageResponse, error)
}

type ProductServiceOp[T any] struct {
	client *Client[T]
}

// AdjustSellableQuantity Use this API to increase or decrease sellable quantity of one or more existing products. The maximum number of products that can be updated is 50, but 20 is recommended.
// Path: /product/stock/sellable/adjust
func (s *ProductServiceOp[T]) AdjustSellableQuantity(ctx context.Context) (*AdjustSellableQuantityResponse, error) {
	path := "/product/stock/sellable/adjust"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(AdjustSellableQuantityResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// BatchUpdateSizeChart 批量更新尺码表
// Path: /size/chart/batch/update
func (s *ProductServiceOp[T]) BatchUpdateSizeChart(ctx context.Context) (*BatchUpdateSizeChartResponse, error) {
	path := "/size/chart/batch/update"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(BatchUpdateSizeChartResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// CreateProduct Use this API to create a single new product.
//
// Find more details below: https://open.lazada.com/apps/doc/doc?nodeId=30720&docId=120949
// Path: /product/create
func (s *ProductServiceOp[T]) CreateProduct(ctx context.Context, req CreateProductRequest) (*CreateProductResponse, error) {
	path := "/product/create"
	params := paramsFromStruct(req)
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(CreateProductResponse)
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

// DeactivateProduct Use this API to deactivate Product or SKUs corresponding to the product
// Path: /product/deactivate
func (s *ProductServiceOp[T]) DeactivateProduct(ctx context.Context) (*DeactivateProductResponse, error) {
	path := "/product/deactivate"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(DeactivateProductResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetBrandByPages Use this API to retrieve all product brands by page index in the system.
// Path: /category/brands/query
func (s *ProductServiceOp[T]) GetBrandByPages(ctx context.Context) (*GetBrandByPagesResponse, error) {
	path := "/category/brands/query"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetBrandByPagesResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetCategoryAttributes Use this API to get a list of attributes for a specified product category.
// Path: /category/attributes/get
func (s *ProductServiceOp[T]) GetCategoryAttributes(ctx context.Context) (*GetCategoryAttributesResponse, error) {
	path := "/category/attributes/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetCategoryAttributesResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetCategorySuggestion Get product's category suggestion by product title
// Path: /product/category/suggestion/get
func (s *ProductServiceOp[T]) GetCategorySuggestion(ctx context.Context) (*GetCategorySuggestionResponse, error) {
	path := "/product/category/suggestion/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetCategorySuggestionResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetCategoryTree Use this API to retrieve the list of all product categories in the system.
// Path: /category/tree/get
func (s *ProductServiceOp[T]) GetCategoryTree(ctx context.Context) (*GetCategoryTreeResponse, error) {
	path := "/category/tree/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetCategoryTreeResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetNextCascadeProp Use this API to query next cascade prop.
// Path: /category/cascade/getNextCascadeProp
func (s *ProductServiceOp[T]) GetNextCascadeProp(ctx context.Context) (*GetNextCascadePropResponse, error) {
	path := "/category/cascade/getNextCascadeProp"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetNextCascadePropResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetPreQcRules query pre qc rules
// Path: /product/seller/item/getPreQcRules
func (s *ProductServiceOp[T]) GetPreQcRules(ctx context.Context) (*GetPreQcRulesResponse, error) {
	path := "/product/seller/item/getPreQcRules"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetPreQcRulesResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetProductContentScore get product content score
// Path: /product/content/score/get
func (s *ProductServiceOp[T]) GetProductContentScore(ctx context.Context) (*GetProductContentScoreResponse, error) {
	path := "/product/content/score/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetProductContentScoreResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetProductItem Get single product by ItemId or SellerSku.
// Path: /product/item/get
func (s *ProductServiceOp[T]) GetProductItem(ctx context.Context, opt GetProductItemRequest) (*GetProductItemResponse, error) {
	path := "/product/item/get"
	params := paramsFromStruct(opt)
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetProductItemResponse)
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

// GetProducts Use this API to get detailed information of the specified products.
// Path: /products/get
func (s *ProductServiceOp[T]) GetProducts(ctx context.Context, opt GetProductsRequest) (*GetProductsResponse, error) {
	path := "/products/get"
	params := paramsFromStruct(opt)
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetProductsResponse)
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

// GetQCAlertProducts Getting seller's products that have been alerted by quality control.
// Path: /product/qc/alert/list
func (s *ProductServiceOp[T]) GetQCAlertProducts(ctx context.Context) (*GetQCAlertProductsResponse, error) {
	path := "/product/qc/alert/list"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetQCAlertProductsResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetResponse Use this API to get the returned information from the system for the MigrateImages API.
// Path: /image/response/get
func (s *ProductServiceOp[T]) GetResponse(ctx context.Context, filename string, reader io.Reader) (*GetResponseResponse, error) {
	path := "/image/response/get"
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", filename)
	if err != nil {
		return nil, err
	}
	if _, err := io.Copy(part, reader); err != nil {
		return nil, err
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}

	wrapper, err := s.client.execute(ctx, "POST", path, nil, map[string][]byte{"image": {}})
	if err != nil {
		return nil, err
	}
	resp := new(GetResponseResponse)
	if err := json.Unmarshal(wrapper.Data, resp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return resp, nil
}

// GetSellerItemLimit The platform will provide the product quantity limit information by this interface. The qps will be limited by seller, 10 qps per seller.
// Path: /product/seller/item/limit
func (s *ProductServiceOp[T]) GetSellerItemLimit(ctx context.Context) (*GetSellerItemLimitResponse, error) {
	path := "/product/seller/item/limit"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetSellerItemLimitResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetSizeChartTemplate 获取尺码模板列表
// Path: /size/chart/template/get
func (s *ProductServiceOp[T]) GetSizeChartTemplate(ctx context.Context) (*GetSizeChartTemplateResponse, error) {
	path := "/size/chart/template/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetSizeChartTemplateResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetUnfilledAttributeItem Get products without key attributes. (For cross boarder sellers Only)
// Path: /product/unfilled/attribute/get
func (s *ProductServiceOp[T]) GetUnfilledAttributeItem(ctx context.Context) (*GetUnfilledAttributeItemResponse, error) {
	path := "/product/unfilled/attribute/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetUnfilledAttributeItemResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// MigrateImage Use this API to migrate a single image from an external site to Lazada site. Allowed image formats are JPG and PNG. The maximum size of an image file is 1MB.
// Path: /image/migrate
func (s *ProductServiceOp[T]) MigrateImage(ctx context.Context, filename string, reader io.Reader) (*MigrateImageResponse, error) {
	path := "/image/migrate"
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", filename)
	if err != nil {
		return nil, err
	}
	if _, err := io.Copy(part, reader); err != nil {
		return nil, err
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}

	wrapper, err := s.client.execute(ctx, "POST", path, nil, map[string][]byte{"image": {}})
	if err != nil {
		return nil, err
	}
	resp := new(MigrateImageResponse)
	if err := json.Unmarshal(wrapper.Data, resp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return resp, nil
}

// MigrateImages Use this API to migrate multiple images from an external site to Lazada site. Allowed image formats are JPG and PNG. The maximum size of an image file is 1MB. A single call can migrate 8 images at most.
// Path: /images/migrate
func (s *ProductServiceOp[T]) MigrateImages(ctx context.Context, filename string, reader io.Reader) (*MigrateImagesResponse, error) {
	path := "/images/migrate"
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", filename)
	if err != nil {
		return nil, err
	}
	if _, err := io.Copy(part, reader); err != nil {
		return nil, err
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}

	wrapper, err := s.client.execute(ctx, "POST", path, nil, map[string][]byte{"image": {}})
	if err != nil {
		return nil, err
	}
	resp := new(MigrateImagesResponse)
	if err := json.Unmarshal(wrapper.Data, resp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return resp, nil
}

// ProductCheck Use this API to check CB seller quantity limit of adding product .
// Path: /product/pre/check
func (s *ProductServiceOp[T]) ProductCheck(ctx context.Context) (*ProductCheckResponse, error) {
	path := "/product/pre/check"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(ProductCheckResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// RemoveProduct Use this API to remove an existing product, some SKUs in one product, or all SKUs in one product. System supports a maximum number of 50 SellerSkus in one request.
// Path: /product/remove
func (s *ProductServiceOp[T]) RemoveProduct(ctx context.Context) (*RemoveProductResponse, error) {
	path := "/product/remove"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(RemoveProductResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// RemoveSku Use this API to delete SKUs and sales attributes of corresponding products.
// Path: /product/sku/remove
func (s *ProductServiceOp[T]) RemoveSku(ctx context.Context) (*RemoveSkuResponse, error) {
	path := "/product/sku/remove"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(RemoveSkuResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// SetImages Use this API to set the images for an existing product by associating one or more image URLs with it.
// Path: /images/set
func (s *ProductServiceOp[T]) SetImages(ctx context.Context, filename string, reader io.Reader) (*SetImagesResponse, error) {
	path := "/images/set"
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", filename)
	if err != nil {
		return nil, err
	}
	if _, err := io.Copy(part, reader); err != nil {
		return nil, err
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}

	wrapper, err := s.client.execute(ctx, "POST", path, nil, map[string][]byte{"image": {}})
	if err != nil {
		return nil, err
	}
	resp := new(SetImagesResponse)
	if err := json.Unmarshal(wrapper.Data, resp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return resp, nil
}

// UpdatePriceQuantity Use this API to update the price and quantity of one or more existing products. The maximum number of products that can be updated is 50, but 20 is recommended.
// Path: /product/price_quantity/update
func (s *ProductServiceOp[T]) UpdatePriceQuantity(ctx context.Context) (*UpdatePriceQuantityResponse, error) {
	path := "/product/price_quantity/update"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(UpdatePriceQuantityResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// UpdateProduct Use this API to update attributes or SKUs of an existing product. if need update inventory, offline, price, not recommended to use this API.
// The iteration 25/6/2020 Updated for DBS changes. Refer to Input Parameters Payload
// Path: /product/update
func (s *ProductServiceOp[T]) UpdateProduct(ctx context.Context, req UpdateProductRequest) (*UpdateProductResponse, error) {
	path := "/product/update"
	params := paramsFromStruct(req)
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(UpdateProductResponse)
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

// UpdateSellableQuantity Use this API to update sellable quantity of one or more existing products. The maximum number of products that can be updated is 50, but 20 is recommended.
// Path: /product/stock/sellable/update
func (s *ProductServiceOp[T]) UpdateSellableQuantity(ctx context.Context) (*UpdateSellableQuantityResponse, error) {
	path := "/product/stock/sellable/update"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(UpdateSellableQuantityResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// UploadImage Use this API to upload a single image file to Lazada site. Allowed image formats are JPG and PNG. The maximum size of an image file is 1MB.
// Path: /image/upload
func (s *ProductServiceOp[T]) UploadImage(ctx context.Context, filename string, reader io.Reader) (*UploadImageResponse, error) {
	path := "/image/upload"
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", filename)
	if err != nil {
		return nil, err
	}
	if _, err := io.Copy(part, reader); err != nil {
		return nil, err
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}

	wrapper, err := s.client.execute(ctx, "POST", path, nil, map[string][]byte{"image": {}})
	if err != nil {
		return nil, err
	}
	resp := new(UploadImageResponse)
	if err := json.Unmarshal(wrapper.Data, resp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return resp, nil
}
