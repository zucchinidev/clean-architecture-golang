package infrastructure

import (
	"github.com/zucchinidev/clean-architecture-golang/usescases"
	"net/http"
	"strconv"
	"io"
	"fmt"
)

type WebServiceHandler struct {
	OrderInteractor usescases.OrderInteractorService
}

func (handler WebServiceHandler) ShowOrder(res http.ResponseWriter, req *http.Request) {
	userId, _ := strconv.Atoi(req.FormValue("userId"))
	orderId, _ := strconv.Atoi(req.FormValue("orderId"))
	items, _ := handler.OrderInteractor.Items(userId, orderId)
	for _, item := range items {
		io.WriteString(res, fmt.Sprintf("item id: %d\n", item.Id))
		io.WriteString(res, fmt.Sprintf("item name: %v\n", item.Name))
		io.WriteString(res, fmt.Sprintf("item value: %f\n", item.Value))
	}
}
