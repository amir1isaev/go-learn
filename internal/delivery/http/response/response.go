package response

import (
	"go-learn/pkg/pagination"
	"math"
	"reflect"
)

type (
	AccessResponse struct {
		Allow bool `json:"allow"`
	}

	PaginationRequest struct {
		CurrentPage uint64 `json:"current_page" default:"1" form:"current_page,default=1" binding:"omitempty"`
		PerPage     uint64 `json:"per_page" default:"15" form:"per_page,default=15" validate:"gte=1" binding:"omitempty"`
	}

	PaginationResponse struct {
		PaginationRequest
		TotalPages   uint `json:"total_pages" validate:"required"`
		TotalEntries int  `json:"total_entries" validate:"required"`
	}

	IDRequest struct {
		ID string `json:"-" form:"id" uri:"id" validate:"uuid4,required"`
	}
)

type BaseResponse struct {
	Meta *PaginationResponse `json:"meta,omitempty" validate:"required"`
	Data interface{}         `json:"data" validate:"required"`
}

func (p PaginationRequest) ToPagination() pagination.Pagination {
	if p.CurrentPage > 0 {
		p.CurrentPage--
	}
	return pagination.Pagination{
		Limit:  p.PerPage,
		Offset: p.CurrentPage * p.PerPage,
	}
}

type IDResponse struct {
	ID string `json:"id"`
}

func NewResponse(data any) BaseResponse {
	return BaseResponse{
		Data: data,
	}
}

func NewListResponse(pagination PaginationRequest, totalRows int, data interface{}) (view BaseResponse) {
	view.Meta = &PaginationResponse{
		PaginationRequest: pagination,
		TotalEntries:      totalRows,
		TotalPages:        uint(math.Ceil(float64(totalRows) / float64(pagination.PerPage))),
	}
	if reflect.ValueOf(data).IsNil() {
		view.Data = make([]int, 0)
	} else {
		view.Data = data
	}

	return view
}
