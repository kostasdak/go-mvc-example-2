# Example 02 - gomvc package

Example 02 for MVC (Model View Controller) implementation with Golang using MySql database

## Overview
Web app with 5 pages :</br>
    - Home (static)</br>
    - Products -> View, Edit, Create, Delete product</br>
    - Related table colors for each product</br>
    - About (static)</br>
    - Contact (static)</br>

Database :</br>
`/database/example_1.sql`</br>

Steps :</br>
* Edit config file `configs/config.yml`
* Setup MySql database `example_1.sql`
* Load config file `configs/config.yaml`</br>
* Connect to MySql database</br>
* Start your server</br>
* Write code to initialize your Models and Controllers</br>
* Write your standard text/Template files (Views)</br>
* Enjoy</br>


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

### Load config file, Connect database, Start server

```
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

### Write code to use gomvc package
### AppHandler

```
func AppHandler(db *sql.DB, cfg *gomvc.AppConfig) http.Handler {

	// initialize
	c.Initialize(db, cfg)

	// load template files ... path : /web/templates
	c.CreateTemplateCache("home.view.tmpl", "base.layout.html")

	// *** Start registering url, actions and models ***
	// home page
	c.RegisterAction("/", "", gomvc.ActionView, nil)
	c.RegisterAction("/home", "", gomvc.ActionView, nil)

	// create view model for [products] table, use this model only for view data
	pViewModel := gomvc.Model{DB: db, IdField: "id", TableName: "products"}
	pViewModel.AddRelation(db, "colors", "id", "product_id", gomvc.ModelJoinLeft)

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

	// view products
	c.RegisterAction("/products", "", gomvc.ActionView, &pViewModel)
	c.RegisterAction("/products/view/*", "", gomvc.ActionView, &pViewModel)

	// create edit model for table [products] -> used for create, update, delete only
	pEditModel := gomvc.Model{DB: db, IdField: "id", TableName: "products"}

	// create product
	c.RegisterAction("/products/create", "", gomvc.ActionView, &pEditModel)
	c.RegisterAction("/products/create", "/products", gomvc.ActionCreate, &pEditModel)

	// edit product
	c.RegisterAction("/products/edit/*", "", gomvc.ActionView, &pEditModel)
	c.RegisterAction("/products/edit/*", "/products", gomvc.ActionUpdate, &pEditModel)

	// delete product
	c.RegisterAction("/products/delete/*", "", gomvc.ActionView, &pEditModel)
	c.RegisterAction("/products/delete/*", "/products", gomvc.ActionDelete, &pEditModel)

	// about page
	c.RegisterAction("/about", "", gomvc.ActionView, nil)

	// contact page
	c.RegisterAction("/contact", "", gomvc.ActionView, nil)

	// custom action
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