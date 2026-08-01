package product

import (
	"ecommerce/util"
	"net/http"
	"strconv"
)

// when we use go routine, we need to declare cnt as global variable
// var cnt int64

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	reqQuery := r.URL.Query()

	pageAsStr := reqQuery.Get("page")
	limitAsStr := reqQuery.Get("limit")

	page, _ := strconv.ParseInt(pageAsStr, 10, 32)
	limit, _ := strconv.ParseInt(limitAsStr, 10, 32)

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	productList, err := h.svc.List(page, limit)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	cnt, err := h.svc.Count()
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// if use go routine
	// go func() {
	// 	cnt1, err := h.svc.Count()
	// 	if err != nil {
	// 		util.SendError(w, http.StatusInternalServerError, "Internal Server Error")
	// 		return
	// 	}
	// 	cnt = cnt1
	// }()
	// time.Sleep(1 * time.Second)

	util.SendPage(w, productList, page, limit, cnt)
}

/*
	** To prevent concurrency issues, we use locking.
	** Go routines are lightweight threads that allow concurrent execution of functions. However, when multiple goroutines access shared data, it can lead to race conditions and unpredictable behavior. To prevent this, we can use synchronization mechanisms like mutexes or channels.
	- Why go channels matter?
	- To prevent race conditions, we can use channels
	- To share data between goroutines, we can use channels
*/
