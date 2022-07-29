package resource

type Resource struct {
	Name string `json:"name"`
	Min  int    `json:"min"`
	Max  int    `json:"max"`
}

type ResourceStatus struct {
	Num int `json:"num"`
}

func NewStatus() *ResourceStatus {
	return &ResourceStatus{
		Num: 0,
	}
}
