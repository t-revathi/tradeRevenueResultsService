package controller

import (
	"fmt"
	"net/http"
	service "tradeRevenueResultsService/service/dataservice"
	mongo "tradeRevenueResultsService/service/mongodb"

	"github.com/go-chi/chi"
	"github.com/google/martian/filter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/go-chi/render"
)

type DataController struct {
	dataService *service.DataService
	dbService   *mongo.DBService
}

func NewDataController(db string) *DataController {
	return &DataController{
		dataService: service.NewDataService(),
		dbService:   mongo.NewDBService(),
	}
}
func Health(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func (d *DataController) WireRoutes(r chi.Router) {
	r.Route("/", func(r chi.Router) {
		r = r.With(Health)
		r.Get("/allRevenue", d.ShowAllRevenue)
		//r.Get("/revenue?user=jon")

		r.Get("/healthCheck", d.healthCheck)
	})
}

func (d *DataController) ShowAllRevenue(w http.ResponseWriter, r *http.Request) {
	/*var dataRevenue service.DataCalculateRevenue

	if err := render.DecodeJSON(r.Body, &dataCalculateRevenue); err != nil {
		return
	}
	
*/
	
	client, ctx, cancel, err := d.dbService.Connectdb("mongodb://0.0.0.0:27017/stockprofitcalculator")

	if err != nil {
		panic(err)
	}
	err = d.dbService.Pingdb(client, ctx)

	defer d.dbService.Closedb(client, ctx, cancel)
	if err != nil {
		fmt.Println("Couldn't connect to Database")
	}


	collection := "plResults"
	filter := bson.M{}
	// Get all records 
	d.dbService.FindAll(client,ctx,"stockprofitcalculator",collection,filter)
	//Return results to client
	render.JSON(w, r,
		result.Items)
}

func (d *DataController) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Healthcheck good"))
}
