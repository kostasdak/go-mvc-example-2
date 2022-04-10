# Example 02 - gomvc package

Example 02</br>
gomvc - MVC (Model View Controller) implementation with Golang using MySQL database

## Overview
Web app with 5 pages :</br>
    * Home (static)</br>
    * Products -> View products / product (related table colors for each product)</br>
    * Products Edit, Create, Delete product</br>
    * About (static)</br>
    * Contact (static)</br>

Database :</br>
`/database/example_2.sql`</br>

Steps :</br>
* Setup MySQL database `example_2.sql` and MySQL server
* Edit config file `configs/config.yml`
* Load config file `configs/config.yaml`</br>
* Connect to MySQL database</br>
* Write code to initialize your Models and Controllers</br>
* Write your standard text/Template files (Views)</br>
* Build and enjoy</br>


### Edit configuration file

```
#UseCache true/false 
#Read files for every request, use this option for debug and development, set to true on production server
UseCache: false

#EnableInfoLog true/false
#Enable information log in console window, set to false in production server
EnableInfoLog: true

#InfoFile "path.to.filename"
#Set info filename, direct info log to file instead of console window
InfoFile: ""

#ShowStackOnError true/false
#Set to true to see the stack error trace in web page error report, set to false in production server
ShowStackOnError: true

#ErrorFile "path.to.filename"
#Set error filename, direct error log to file instead of web page, set this file name in production server
ErrorFile: ""

#Server Settings
server:
  #Listening port
  port: 8080

  #Session timeout in hours 
  sessionTimeout: 24

  #Use secure session, set to tru in production server
  sessionSecure: true

#Database settings
database:
  #Database name
  dbname: "golang"

  #Database server/ip address
  server: "localhost"

  #Database user
  dbuser: "root"

  #Database password
  dbpass: ""
```

### Load config file, Connect database, Start http server

```

var c gomvc.Controller

func main() {

	// Load Configuration file
	cfg := gomvc.LoadConfig("./config/config.yml")

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
```

### Write code with gomvc package
### AppHandler

```
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

	// home page : can have two urls "/" and "/home"
	c.RegisterAction("/", "", gomvc.ActionView, nil)
	c.RegisterAction("/home", "", gomvc.ActionView, nil)

	// create model for [products] database table
	// use the same model for all action in this example
	pViewModel := gomvc.Model{DB: db, PKField: "id", TableName: "products", OrderString: "ORDER BY id DESC"}
	pViewModel.AddRelation(db, "colors", "id", "product_id", gomvc.ModelJoinLeft, gomvc.ResultStyleSubresult)

	// optional assign labels for each table field. Can be used in template view code
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
	c.RegisterAction("/products", "", gomvc.ActionView, &pViewModel)
	c.RegisterAction("/products/view/*", "", gomvc.ActionView, &pViewModel)

	// create edit model for table [products] -> used for create, update, delete actions only
	pEditModel := gomvc.Model{DB: db, PKField: "id", TableName: "products"}

	// build create product action ... this url has two actions
	// #1 View page -> empty product form no redirect url (no next url required)
	// #2 Post form data to create a new record in table [products] -> then redirect to [next] url -> products page
	c.RegisterAction("/products/create", "", gomvc.ActionView, &pEditModel)
	c.RegisterAction("/products/create", "/products", gomvc.ActionCreate, &pEditModel)

	// build edit product actions ... this url has two actions
	// #1 View page with the product form -> edit form (no next url required)
	// #2 Post form data to update record in table [products] -> then redirect to [next] url -> products page
	c.RegisterAction("/products/edit/*", "", gomvc.ActionView, &pEditModel)
	c.RegisterAction("/products/edit/*", "/products", gomvc.ActionUpdate, &pEditModel)

	// build delete product actions ... this url has two actions
	// #1 View page with the product form -> edit form [locked] to confirm detetion (no next url required)
	// #2 Post form data to delete record in table [products] -> then redirect to [next] url -> products page
	c.RegisterAction("/products/delete/*", "", gomvc.ActionView, &pEditModel)
	c.RegisterAction("/products/delete/*", "/products", gomvc.ActionDelete, &pEditModel)

	// build about page ... static page, no table/model, no [next] url
	c.RegisterAction("/about", "", gomvc.ActionView, nil)

	// build contact page ... static page, no table/model, no [next] url
	c.RegisterAction("/contact", "", gomvc.ActionView, nil)

	// build contact page POST action ... static page, no table/model, no [next] url
	// Demostrating how to register a custom func to handle the http request/response using your oun code
	// and handle POST data and have access to database through the controller and model object
	c.RegisterCustomAction("/contact", "", gomvc.HttpPOST, nil, ContactPostForm)
	return c.Router
}
```

### Custom handler

```
// Custom handler for specific page and action, 
// this function handles the POST action from "Contact Us" page 
func ContactPostForm(w http.ResponseWriter, r *http.Request) {

	//test ... I have access to products model !!!
	fmt.Println("**********")
	fmt.Println("Table Fields : ", c.Models["/products"].Fields)

	//read data from table products even this is a POST action for contact page
	rows, _ := c.Models["/products"].GetAllRecords(100)
	fmt.Println("**********")
	fmt.Println("Select Rows : ", rows)

	//test form Posted fields
	fmt.Println("**********")
	fmt.Println("Form Fields : ", r.Form)

	//test session message
	c.GetSession().Put(r.Context(), "error", "Hello From Session")

	//redirect to homepage
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
```