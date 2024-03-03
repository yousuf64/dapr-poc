package main

import (
	"encoding/json"
	"github.com/dapr/go-sdk/client"
	"github.com/yousuf64/go-event-sourcing/command-service/api/requests"
	"github.com/yousuf64/go-event-sourcing/command-service/internal/commands"
	"github.com/yousuf64/go-event-sourcing/command-service/internal/handlers"
	"github.com/yousuf64/go-event-sourcing/command-service/pkg/idprovider"
	"github.com/yousuf64/go-event-sourcing/command-service/pkg/store"
	"github.com/yousuf64/shift"
	"log"
	"net/http"
)

type Containers struct {
	ProductHandlers *handlers.ProductHandlers
}

func main() {
	router := shift.New()
	router.Use(func(next shift.HandlerFunc) shift.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request, route shift.Route) error {
			err := next(w, r, route)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte(err.Error()))
			}
			return nil
		}
	})

	idp, err := idprovider.New()
	if err != nil {
		log.Fatalln("error initializing id provider", err)
	}
	daprClient, err := client.NewClient()
	if err != nil {
		log.Println("bombolocat", err)
	}
	str := store.New(daprClient)
	containers := &Containers{ProductHandlers: handlers.NewProductHandlers(idp, str)}

	router.POST("/v1/buckets/:bucketId/products", createProduct(containers))
	router.POST("/v1/buckets/:bucketId/products/:productId", updateProduct(containers))
	router.DELETE("/v1/buckets/:bucketId/products/:productId", deleteProduct(containers))

	log.Println("running on :8888")
	err = http.ListenAndServe(":8888", router.Serve())
	if err != nil {
		log.Fatalln(err)
	}
}

func createProduct(containers *Containers) shift.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, route shift.Route) error {
		var data requests.CreateProduct
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			return err
		}

		out, err := containers.ProductHandlers.CreateProduct(r.Context(), commands.CreateProduct{
			BucketId:    route.Params.Get("bucketId"),
			Name:        data.Name,
			Description: data.Description,
			Price:       data.Price,
			Quantity:    data.Quantity,
		})
		if err != nil {
			return err
		}

		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(out)
		if err != nil {
			return err
		}
		return nil
	}
}

func updateProduct(containers *Containers) shift.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, route shift.Route) error {
		var data requests.UpdateProduct
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			return err
		}

		return containers.ProductHandlers.UpdateProduct(r.Context(), commands.UpdateProduct{
			BucketId:    route.Params.Get("bucketId"),
			ProductId:   route.Params.Get("productId"),
			Name:        data.Name,
			Description: data.Description,
			Price:       data.Price,
		})
	}
}

func deleteProduct(containers *Containers) shift.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, route shift.Route) error {
		return containers.ProductHandlers.DeleteProduct(r.Context(), commands.DeleteProduct{
			BucketId:  route.Params.Get("bucketId"),
			ProductId: route.Params.Get("productId"),
		})
	}
}
