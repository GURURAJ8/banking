package app
import (
	"github.com/GURURAJ8/banking/Service"
	"github.com/GURURAJ8/banking/dto"
	"net/http"
	"encoding/json"
	"strconv"
	"github.com/gorilla/mux"
)


type AccountHandler struct{
	Service Service.AccountService
}

func (h AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	cucstomerId := vars["customer_id"]
	var req dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		WriteResponse(w, nil,err.Error(),http.StatusBadRequest)
	}else{
	req.CustomerId,_=strconv.Atoi(cucstomerId)
	response, appErr := h.Service.NewAccountService(req)
	if appErr != nil {
		WriteResponse(w, nil,appErr.Message,appErr.Code)
		return
	}else{
	WriteResponse(w,response,"All Good",http.StatusCreated)
	}
	}
}