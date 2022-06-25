package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/kostasdak/gomvc"
)

var c gomvc.Controller

func main() {

	// Load Configuration file
	cfg := gomvc.ReadConfig("./config/config.yml")

	// Connect to database
	db, err := gomvc.ConnectDatabase(cfg.Database.Dbuser, cfg.Database.Dbpass, cfg.Database.Dbname)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	//Start Server
	srv := &http.Server{
		Addr:    ":" + strconv.FormatInt(int64(cfg.Server.Port), 10),
		Handler: AppHandler(db, cfg),
	}

	fmt.Println("Web app starting at port : ", cfg.Server.Port)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

// App handler ... Builds the structure of the app !!!
func AppHandler(db *sql.DB, cfg *gomvc.AppConfig) http.Handler {

	// initialize controller
	c.Initialize(db, cfg)

	// load template files ... path : /web/templates
	// required : homepagefile & template file
	// see [template names] for details
	c.CreateTemplateCache("home.view.tmpl", "base.layout.html")

	// *** Start registering urls, actions and models ***

	// RegisterAction(url, next, action, model)
	// url = url routing path
	// next = redirect after action complete, use in POST actions if necessary
	// model = database model object for CRUD operations

	// create model for Home page content - Set default query to read specific only content from table pages
	pHomeModel := gomvc.Model{DB: db, PKField: "id", TableName: "pages", OrderString: "ORDER BY id DESC"}
	pHomeModel.DefaultQuery = "SELECT content FROM pages WHERE name='home'"

	// home page : can have two urls "/" and "/home" -> home.view.tmpl must exist
	// use pHomeModel for dynamic page content
	c.RegisterAction(gomvc.ActionRouting{URL: "/"}, gomvc.ActionView, &pHomeModel)
	c.RegisterAction(gomvc.ActionRouting{URL: "/home"}, gomvc.ActionView, &pHomeModel)

	// create model for [products] database table
	// use the same model for all action in this example
	pViewModel := gomvc.Model{DB: db, PKField: "id", TableName: "products", OrderString: "ORDER BY id DESC"}
	pViewModel.AddRelation(db, "colors", "id", gomvc.SQLKeyPair{LocalKey: "id", ForeignKey: "product_id"}, gomvc.ModelJoinLeft, gomvc.ResultStyleSubresult)

	// optionally assign labels for each table field. Can be used in template view code
	pViewModel.AssignLabels(map[string]string{
		"id":                "Id",
		"code":              "Code",
		"type":              "Veh. Type",
		"name":              "Product",
		"description":       "Description",
		"price":             "Price",
		"images":            "Photos",
		"status":            "Availability",
		"colors.id":         "Id",    //Related field in table [colors]
		"colors.product_id": "PId",   //Related field in table [colors]
		"colors.color":      "Color", //Related field in table [colors]
	})

	// view products ... / show all products || /products/view/{id} for one product
	c.RegisterAction(gomvc.ActionRouting{URL: "/products"}, gomvc.ActionView, &pViewModel)
	c.RegisterAction(gomvc.ActionRouting{URL: "/products/view/*"}, gomvc.ActionView, &pViewModel)

	// create edit model for table [products] -> used for create, update, delete actions only
	pEditModel := gomvc.Model{DB: db, PKField: "id", TableName: "products"}
	pEditModel.AddRelation(db, "colors", "id", gomvc.SQLKeyPair{LocalKey: "id", ForeignKey: "product_id"}, gomvc.ModelJoinLeft, gomvc.ResultStyleSubresult)

	// prepare create product action ... this url has two actions
	// #1 View page -> empty product form no redirect url (no next url required)
	// #2 Post form data to create a new record in table [products] -> then redirect to [next] url -> products page
	c.RegisterAction(gomvc.ActionRouting{URL: "/products/create"}, gomvc.ActionView, &pEditModel)
	c.RegisterAction(gomvc.ActionRouting{URL: "/products/create", NextURL: "/products"}, gomvc.ActionCreate, &pEditModel)

	// prepare edit product actions ... this url has two actions
	// #1 View page with the product form -> edit form (no next url required)
	// #2 Post form data to update record in table [products] -> then redirect to [next] url -> products page
	c.RegisterAction(gomvc.ActionRouting{URL: "/products/edit/*"}, gomvc.ActionView, &pEditModel)
	c.RegisterAction(gomvc.ActionRouting{URL: "/products/edit/*", NextURL: "products"}, gomvc.ActionUpdate, &pEditModel)

	// prepare delete product actions ... this url has two actions
	// #1 View page with the product form -> edit form [locked] to confirm detetion (no next url required)
	// #2 Post form data to delete record in table [products] -> then redirect to [next] url -> products page
	c.RegisterAction(gomvc.ActionRouting{URL: "/products/delete/*"}, gomvc.ActionView, &pEditModel)
	c.RegisterAction(gomvc.ActionRouting{URL: "/products/delete/*", NextURL: "/products"}, gomvc.ActionDelete, &pEditModel)

	// create product color model
	cModel := gomvc.Model{DB: db, PKField: "id", TableName: "colors"}
	cModel.AddRelation(db, "products", "id", gomvc.SQLKeyPair{LocalKey: "product_id", ForeignKey: "id"}, gomvc.ModelJoinInner, gomvc.ResultStyleFullresult)

	// prepare add product color page
	c.RegisterAction(gomvc.ActionRouting{URL: "/productcolor/add/*"}, gomvc.ActionView, &cModel)
	c.RegisterAction(gomvc.ActionRouting{URL: "/productscolor/add/*", NextURL: "/products"}, gomvc.ActionCreate, &cModel)

	// prepare edit product color page
	c.RegisterAction(gomvc.ActionRouting{URL: "/productcolor/edit/*"}, gomvc.ActionView, &cModel)
	c.RegisterAction(gomvc.ActionRouting{URL: "/productcolor/edit/*", NextURL: "/products"}, gomvc.ActionUpdate, &cModel)

	//prepare delete action / ONLY for post, no view file required
	//redirect will take action after delete action finish
	//see productcolor.edit.tmpl -> file has a form for post to -> productcolor/delete/{id}
	c.RegisterAction(gomvc.ActionRouting{URL: "/productcolor/delete/*", NextURL: "/products"}, gomvc.ActionDelete, &cModel)

	// prepare about page ... static page, no table/model, no [next] url
	c.RegisterAction(gomvc.ActionRouting{URL: "/about"}, gomvc.ActionView, nil)

	// prepare contact page ... static page, no table/model, no [next] url
	c.RegisterAction(gomvc.ActionRouting{URL: "/contact"}, gomvc.ActionView, nil)

	// prepare contact page POST action ... static page, no table/model, no [next] url
	// Demostrating how to register a custom func to handle the http request/response using your oun code
	// and handle POST data and have access to database through the controller and model object
	c.RegisterCustomAction(gomvc.ActionRouting{URL: "/contact"}, gomvc.HttpPOST, nil, ContactPostForm)
	return c.Router
}

// Custom handler for specific page and action
func ContactPostForm(w http.ResponseWriter, r *http.Request) {

	//test : I have access to products model !!!
	fmt.Print("\n\n")
	fmt.Println("********** ContactPostForm **********")
	fmt.Println("Table Fields : ", c.Models["/products"].Fields)

	//read data from table products (Model->products) even if this is a POST action for contact page
	fmt.Print("\n\n")
	rows, _ := c.Models["/products"].GetRecords([]gomvc.Filter{}, 100)
	fmt.Println("Select Rows Example 1 : ", rows)

	//read data from table products (Model->products) even if this is a POST action for contact page
	fmt.Print("\n\n")
	id, _ := c.Models["/products"].GetLastId()
	fmt.Println("Select Rows Example 1 : ", id)

	//read data from table products (Model->products) with filter (id=1)
	fmt.Print("\n\n")
	var f = make([]gomvc.Filter, 0)
	f = append(f, gomvc.Filter{Field: "id", Operator: "=", Value: 1})
	rows, err := c.Models["/products"].GetRecords(f, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Select Rows Example 2 : ", rows)

	//test : Print Posted Form fields
	fmt.Print("\n\n")
	fmt.Println("Form fields : ", r.Form)

	//test : Set session message
	c.GetSession().Put(r.Context(), "error", "Hello From Session")

	//redirect to homepage
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
