package controllers

import (
	"net/http"

	"github.com/mendezdev/bookstore_items-api/domain/items"
	"github.com/mendezdev/bookstore_items-api/services"
	"github.com/mendezdev/bookstore_oauth-go/oauth"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
}

type itemsController struct {
}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		// TODO: return error json to the user.
		return
	}

	item := items.Item{
		Seller: oauth.GetCallerId(r),
	}

	_, err := services.ItemsService.Create(item)
	if err != nil {
		//TODO return error json to the user
	}

	//TODO return created item as JSON with HTTP status 201 - Created
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
