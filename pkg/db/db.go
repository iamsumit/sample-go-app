package db

type DataBaseClient struct{}

type User struct {
	UID             int
	Name            string
	Email           string
	PipelineVersion string
}

func NewClient() *DataBaseClient {
	return &DataBaseClient{}
}

func (DataBaseClient) GetConfigByUID(uid int) User {
	var pipelineVersion string

	if uid == 1 {
		pipelineVersion = "V1"
	} else if uid == 2 {
		pipelineVersion = "V2"
	} else {
		// Default
		pipelineVersion = "V0"
	}

	user := User{
		UID:             uid,
		Name:            "Test",
		Email:           "test@gmail.com",
		PipelineVersion: pipelineVersion,
	}

	return user
}
