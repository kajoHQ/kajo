package main

import (
	"github.com/adhaamehab/kajo/pkg/server"
)

func main() {
	server.NewKajoServer().Start(":8080")
}

// func jobsHandler(w http.ResponseWriter, r *http.Request) {
// 	config, err := rest.InClusterConfig()
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	clientset, err := kubernetes.NewForConfig(config)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	jobs, err := clientset.BatchV1().Jobs("workers").List(r.Context(), metav1.ListOptions{})
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	json.NewEncoder(w).Encode(jobs)
// }
