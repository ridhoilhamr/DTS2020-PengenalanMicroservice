package handler

import (
	"net/http"

	"github.com/ridhoilhamr/DTS2020-PengenalanMicroservice/menu-service/utils"
)

func AddMenu(w http.ResponseWriter, r *http.Request) {
	utils.WrapAPISuccess(w, r, "success", http.StatusOK)
}
