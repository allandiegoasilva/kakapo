package reply

type Empty struct{}

const (
	OK                   = 200
	ValidationError      = 400
	Unauthorized         = 401
	Forbidden            = 403
	NotFound             = 404
	InternalServerError  = 500
)

type Reply[T Empty] struct {
	Code       int         `json:"code"`
	Message    string      `json:"message,omitempty"`
	Data       *T          `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}
