package atom_utils

type Meta struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Data struct {
	Data interface{} `json:"data"`
}
