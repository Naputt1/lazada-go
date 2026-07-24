package golazada

import (
	"context"
)

type MembershipService interface {
	// GetLinkMember Query the linkmember relationship between buyers and sellers.
	// Path: /membership/linkmember/get
	GetLinkMember(ctx context.Context) (*GetLinkMemberResponse, error)
	// GetLinkMember1 Query the linkmember relationship between buyers and sellers.
	// Path: /partner/get
	GetLinkMember1(ctx context.Context) (*GetLinkMember1Response, error)
	// GetLinkMemberList Query all linkmembers of the seller
	// Path: /membership/linkmember/list
	GetLinkMemberList(ctx context.Context) (*GetLinkMemberListResponse, error)
	// GetLinkMemberList1 Query all linkmembers of the seller
	// Path: /partner/list
	GetLinkMemberList1(ctx context.Context) (*GetLinkMemberList1Response, error)
	// LinkMembership Used to push a new membership to Lazada for proactively linking memberships.
	// Path: /membership/link
	LinkMembership(ctx context.Context) (*LinkMembershipResponse, error)
	// PartnerLink Used to push a new membership to Lazada for proactively linking memberships.
	// Path: /partner/link
	PartnerLink(ctx context.Context) (*PartnerLinkResponse, error)
	// PartnerTransaction Using this interface, you can obtain the seller's transaction order based on the conditions, and also contain the membership information
	// Path: /partner/transaction
	PartnerTransaction(ctx context.Context) (*PartnerTransactionResponse, error)
	// PartnerUnlink Used to remove a linked membership from Lazada. Please note that the link will not physically be removed, but deactivated.
	// Path: /partner/unlink
	PartnerUnlink(ctx context.Context) (*PartnerUnlinkResponse, error)
	// PartnerUpdate Used to push membership bulk status updates to Lazada. Please note that this is not an incremental update, thus information left out that haven been in our system before, will be removed on our end.
	// Path: /partner/update
	PartnerUpdate(ctx context.Context) (*PartnerUpdateResponse, error)
	// UpdatePartnerUserId Used to update the partner user id to new partner user id
	// Path: /partner/updatePartnerUserId
	UpdatePartnerUserId(ctx context.Context) (*UpdatePartnerUserIdResponse, error)
}

type MembershipServiceOp[T any] struct {
	client *Client[T]
}

// GetLinkMember Query the linkmember relationship between buyers and sellers.
// Path: /membership/linkmember/get
func (s *MembershipServiceOp[T]) GetLinkMember(ctx context.Context) (*GetLinkMemberResponse, error) {
	path := "/membership/linkmember/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetLinkMemberResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetLinkMember1 Query the linkmember relationship between buyers and sellers.
// Path: /partner/get
func (s *MembershipServiceOp[T]) GetLinkMember1(ctx context.Context) (*GetLinkMember1Response, error) {
	path := "/partner/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetLinkMember1Response)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetLinkMemberList Query all linkmembers of the seller
// Path: /membership/linkmember/list
func (s *MembershipServiceOp[T]) GetLinkMemberList(ctx context.Context) (*GetLinkMemberListResponse, error) {
	path := "/membership/linkmember/list"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetLinkMemberListResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetLinkMemberList1 Query all linkmembers of the seller
// Path: /partner/list
func (s *MembershipServiceOp[T]) GetLinkMemberList1(ctx context.Context) (*GetLinkMemberList1Response, error) {
	path := "/partner/list"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetLinkMemberList1Response)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// LinkMembership Used to push a new membership to Lazada for proactively linking memberships.
// Path: /membership/link
func (s *MembershipServiceOp[T]) LinkMembership(ctx context.Context) (*LinkMembershipResponse, error) {
	path := "/membership/link"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(LinkMembershipResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// PartnerLink Used to push a new membership to Lazada for proactively linking memberships.
// Path: /partner/link
func (s *MembershipServiceOp[T]) PartnerLink(ctx context.Context) (*PartnerLinkResponse, error) {
	path := "/partner/link"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(PartnerLinkResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// PartnerTransaction Using this interface, you can obtain the seller's transaction order based on the conditions, and also contain the membership information
// Path: /partner/transaction
func (s *MembershipServiceOp[T]) PartnerTransaction(ctx context.Context) (*PartnerTransactionResponse, error) {
	path := "/partner/transaction"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(PartnerTransactionResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// PartnerUnlink Used to remove a linked membership from Lazada. Please note that the link will not physically be removed, but deactivated.
// Path: /partner/unlink
func (s *MembershipServiceOp[T]) PartnerUnlink(ctx context.Context) (*PartnerUnlinkResponse, error) {
	path := "/partner/unlink"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(PartnerUnlinkResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// PartnerUpdate Used to push membership bulk status updates to Lazada. Please note that this is not an incremental update, thus information left out that haven been in our system before, will be removed on our end.
// Path: /partner/update
func (s *MembershipServiceOp[T]) PartnerUpdate(ctx context.Context) (*PartnerUpdateResponse, error) {
	path := "/partner/update"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(PartnerUpdateResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// UpdatePartnerUserId Used to update the partner user id to new partner user id
// Path: /partner/updatePartnerUserId
func (s *MembershipServiceOp[T]) UpdatePartnerUserId(ctx context.Context) (*UpdatePartnerUserIdResponse, error) {
	path := "/partner/updatePartnerUserId"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(UpdatePartnerUserIdResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}
