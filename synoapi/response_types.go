package synoapi

type SynoBaseResponse interface {
	ErrorCode() int
	Successful() bool
}

type synoBaseResponse struct {
	Error struct {
		Code int
	}
	Success bool
}

func (s *synoBaseResponse) ErrorCode() int {
	return s.Error.Code
}

func (s *synoBaseResponse) Successful() bool {
	return s.Success
}

type SynoResponse struct {
	synoBaseResponse
	Data map[string]interface{}
}
