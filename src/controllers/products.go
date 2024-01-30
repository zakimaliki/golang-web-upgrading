package controllers

import (
	"encoding/json"
	"fmt"
	"golang-web/src/helper"
	"golang-web/src/middleware"
	"golang-web/src/models"
	"net/http"
)

func ProductsController(w http.ResponseWriter, r *http.Request) {
	helper.EnableCors(w)
	middleware.GetCleanedInput(r)
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		res := models.SelectALL()
		result, _ := json.Marshal(res.Value)
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	} else if r.Method == "POST" {
		var input models.Product
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "invalid request body")
			return
		}

		newProduct := models.Product{
			Name:  input.Name,
			Price: input.Price,
			Stock: input.Stock,
		}
		res := models.Create(&newProduct)
		var _, _ = json.Marshal(res)
		fmt.Fprintln(w, "Product created")
		w.WriteHeader(http.StatusCreated)
		return
	} else {
		http.Error(w, "", http.StatusBadRequest)
	}
}

func ProductController(w http.ResponseWriter, r *http.Request) {
	helper.EnableCors(w)
	middleware.GetCleanedInput(r)
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Path[len("/product/"):]

	if r.Method == "GET" {
		res := models.Select(id)
		result, _ := json.Marshal(res.Value)
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	} else if r.Method == "PUT" {
		var input models.Product
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "invalid request body")
			return
		}
		updateProduct := models.Product{
			Name:  input.Name,
			Price: input.Price,
			Stock: input.Stock,
		}
		res := models.Updates(id, &updateProduct)
		var _, _ = json.Marshal(res)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Product updated")
		return
	} else if r.Method == "DELETE" {
		res := models.Deletes(id)
		var _, _ = json.Marshal(res)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Product Deleted")
	} else {
		http.Error(w, "", http.StatusBadRequest)
	}
}

func HandleUpload(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("File")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()
	helper.ValidationUpload(w, handler)
	helper.Upload(w, file, handler)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Upload file Successful")
}

// func ProductsController(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "GET" {
// 		w.Header().Set("Content-Type", "application/json")
// 		result, _ := json.Marshal(models.Products)
// 		w.WriteHeader(http.StatusOK)
// 		w.Write(result)
// 	} else if r.Method == "POST" {
// 		var product models.Product
// 		err := json.NewDecoder(r.Body).Decode(&product)
// 		if err != nil {
// 			w.WriteHeader(http.StatusBadRequest)
// 			fmt.Fprintln(w, "invalid request body")
// 			return
// 		}
// 		models.Products = append(models.Products, product)
// 		w.WriteHeader(http.StatusCreated)
// 		fmt.Fprintln(w, "Product created")
// 		return
// 	}
// 	http.Error(w, "", http.StatusBadRequest)
// }

// func ProductsController(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "GET" {
// 		w.Header().Set("Content-Type", "application/json")
// 		result, _ := json.Marshal(models.Products)
// 		w.WriteHeader(http.StatusOK)
// 		w.Write(result)
// 	} else if r.Method == "POST" {
// 		var product models.Product
// 		err := json.NewDecoder(r.Body).Decode(&product)
// 		if err != nil {
// 			w.WriteHeader(http.StatusBadRequest)
// 			fmt.Fprintln(w, "invalid request body")
// 			return
// 		}
// 		models.Products = append(models.Products, product)
// 		w.WriteHeader(http.StatusCreated)
// 		fmt.Fprintln(w, "Product created")
// 		return
// 	}
// 	http.Error(w, "", http.StatusBadRequest)
// }

// func ProductController(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	idParam := r.URL.Path[len("/product/"):]
// 	id, _ := strconv.Atoi(idParam)

// 	var foundIndex = -1
// 	for i, p := range models.Products {
// 		if p.Id == id {
// 			foundIndex = i
// 			break
// 		}
// 	}

// 	if foundIndex == -1 {
// 		http.Error(w, "Product Not Found", http.StatusNotFound)
// 		return
// 	}

// 	if r.Method == "GET" {
// 		result, _ := json.Marshal(models.Products[foundIndex])
// 		w.WriteHeader(http.StatusOK)
// 		w.Write(result)
// 	} else if r.Method == "PUT" {
// 		var updateProduct models.Product
// 		err := json.NewDecoder(r.Body).Decode(&updateProduct)
// 		if err != nil {
// 			w.WriteHeader(http.StatusBadRequest)
// 			fmt.Fprintln(w, "invalid request body")
// 			return
// 		}
// 		models.Products[foundIndex] = updateProduct
// 		w.WriteHeader(http.StatusOK)
// 		fmt.Fprintln(w, "Product updated")
// 		return

// 	} else if r.Method == "DELETE" {
// 		_ = append(models.Products[:foundIndex], models.Products[foundIndex+1:]...)
// 		w.WriteHeader(http.StatusOK)
// 		fmt.Fprintln(w, "Product Deleted")
// 	}
// 	http.Error(w, "", http.StatusBadRequest)
// }
