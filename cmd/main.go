package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/onuryilmaz/k8s-scheduler-extender/pkg/filter"
	"github.com/onuryilmaz/k8s-scheduler-extender/pkg/prioritize"
	"github.com/sirupsen/logrus"

	extenderv1 "k8s.io/kube-scheduler/extender/v1"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/filter", filterHandler)
	r.HandleFunc("/prioritize", prioritizeHandler)

	log.Fatal(http.ListenAndServe(":8888", r))
}

func filterHandler(w http.ResponseWriter, r *http.Request) {

	args := extenderv1.ExtenderArgs{}
	response := extenderv1.ExtenderFilterResult{}

	if err := json.NewDecoder(r.Body).Decode(&args); err != nil {
		response.Error = err.Error()
	} else {
		response = filter.Filter(args)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		logrus.Error(err)
		return
	}

}

func prioritizeHandler(w http.ResponseWriter, r *http.Request) {

	args := extenderv1.ExtenderArgs{}
	response := make(extenderv1.HostPriorityList, 0)

	if err := json.NewDecoder(r.Body).Decode(&args); err == nil {
		response = prioritize.Prioritize(args)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		logrus.Error(err)
		return
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	var text string = `
              __             __      __                          __                 __         
   __________/ /_  ___  ____/ /_  __/ /__  _____      ___  _  __/ /____  ____  ____/ /__  _____
  / ___/ ___/ __ \/ _ \/ __  / / / / / _ \/ ___/_____/ _ \| |/_/ __/ _ \/ __ \/ __  / _ \/ ___/
 (__  ) /__/ / / /  __/ /_/ / /_/ / /  __/ /  /_____/  __/>  </ /_/  __/ / / / /_/ /  __/ /    
/____/\___/_/ /_/\___/\__,_/\__,_/_/\___/_/         \___/_/|_|\__/\___/_/ /_/\__,_/\___/_/                                                                                                 
`

	w.Write([]byte(text))
}
