package structs

type PaginatedResponse struct {
    Success bool        `json:"success"`
    Message string      `json:"message"`
    Data    interface{} `json:"data"`
    Page    int         `json:"page"`
    Size    int         `json:"size"`
    Total   int64       `json:"total"`
}