package models

type ComputeImages struct {
	Images []ComputeImagesValue `json:"images"`
}

type ComputeImagesValue struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Status   string `json:"status"`
	MinDisk  int    `json:"minDisk"`
	MinRam   int    `json:"minRam"`
	Progress int    `json:"progress"`
	//      Created  time.Time `json:"created"`
	//      Updated  time.Time `json:"updated"`
}
