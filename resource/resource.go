package resource

type Resource struct {
	Name string `json:"name"`
	Min  int    `json:"min"`
	Max  int    `json:"max"`
}

type ResourceStatus struct {
	Num      int      `json:"num"`
	Resource Resource `json:"resource"`
}

func NewStatus(res Resource) *ResourceStatus {
	return &ResourceStatus{
		Num:      0,
		Resource: res,
	}
}
