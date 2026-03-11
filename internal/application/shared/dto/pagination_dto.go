package dto

type PaginationDTO struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

type PaginatedResultDTO struct {
	Data       interface{}   `json:"data"`
	Pagination PaginationDTO `json:"pagination"`
}

func NewPaginatedResult(data interface{}, limit, offset, total int) PaginatedResultDTO {
	return PaginatedResultDTO{
		Data: data,
		Pagination: PaginationDTO{
			Limit:  limit,
			Offset: offset,
			Total:  total,
		},
	}
}
