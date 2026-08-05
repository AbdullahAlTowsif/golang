package product

import (
	"ecommerce/util"
	"net/http"
	"strconv"
	"sync"
	// "time"
)

// when we use go routine, we need to declare cnt as global variable
var cnt int64

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

	// cnt, err := h.svc.Count()
	// if err != nil {
	// 	util.SendError(w, http.StatusInternalServerError, "Internal Server Error")
	// 	return
	// }

	var wg sync.WaitGroup
	var mu sync.Mutex

	// if there are multiple go routine add --> wg.Add(1) --> before each go routine function
	wg.Add(1)

	// if use go routine
	go func() {
		// defer wg.Add(-1) // way 2
		defer wg.Done() // way 3

		mu.Lock()
		defer mu.Unlock()

		cnt1, err := h.svc.Count()
		if err != nil {
			util.SendError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		cnt = cnt1
		// wg.Add(-1) // way 1
	}()
	// time.Sleep(2 * time.Second)
	wg.Wait()

	util.SendPage(w, productList, page, limit, cnt)
}

/*
	** To prevent concurrency issues, we use locking.
	** Go routines are lightweight threads that allow concurrent execution of functions. However, when multiple goroutines access shared data, it can lead to race conditions and unpredictable behavior. To prevent this, we can use synchronization mechanisms like mutexes or channels.
	- Why go channels matter?
	- To prevent race conditions, we can use channels
	- To share data between goroutines, we can use channels
*/
