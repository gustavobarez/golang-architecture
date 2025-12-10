package output

import "github.com/julienschmidt/httprouter"

type userClient struct {

}

func NewUserClient() *userClient {
	client = httprouter.New()
	return &userClient{}
}

var (
	client *httprouter.Router
)