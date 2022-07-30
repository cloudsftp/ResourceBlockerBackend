package resource

type Resource struct {
	Name string `json:"name"`
	Min  int    `json:"min"`
	Max  int    `json:"max"`
}

type ResourceStatus struct {
	Name string `json:"name"`
	Num  int    `json:"num"`
}

func NewStatus(res *Resource) *ResourceStatus {
	return &ResourceStatus{
		Name: res.Name,
		Num:  0,
	}
}
