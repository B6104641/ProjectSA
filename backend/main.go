package main

import (
	"context"
	"log"

	"github.com/B6104641/app/controllers"
	"github.com/B6104641/app/ent"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Users struct input data
type Users struct {
	User []User
}

// User struct
type User struct {
	USEREMAIL    string
	USERPASSWORD string
	USERNAME     string
}

// Equipments struct
type Equipments struct {
	Equipment []Equipment
}

// Equipment struct
type Equipment struct {
	EQUIPMENTNAME  string
	EQUIPMENTPRICE int
}

// Companys struct
type Companys struct {
	Company []Company
}

// Company struct
type Company struct {
	COMPANYName string
}

// @title SUT SA Example API Order
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:OrderBye.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewUserController(v1, client)
	controllers.NewEquipmentController(v1, client)
	controllers.NewCompanyController(v1, client)
	controllers.NewOrderController(v1, client)

	// Set Users Data
	users := Users{
		User: []User{
			User{"gamse0505@gmail.com", "Aa123", "Somchai Ngaosri"},
			User{"Panyaporn@gmail.com", "Bb123", "Panyaporn Ngaosri"},
		},
	}

	for _, u := range users.User {
		client.User.
			Create().
			SetUSEREMAIL(u.USEREMAIL).
			SetUSERPASSWORD(u.USERPASSWORD).
			SetUSERNAME(u.USERNAME).
			Save(context.Background())
	}

	// Set Equipments Data
	equipments := Equipments{
		Equipment: []Equipment{
			Equipment{"เครื่องช่วยหายใจ (Ventilator)", 42500},
			Equipment{"ไฟฉายส่องรูม่านตา (Penlight)", 270},
			Equipment{"เปลสนาม (Pitch Crib)", 2499},
			Equipment{"ชุดช่วยหายใจด้วยมือจับ (Resuscitator)", 500},
			Equipment{"เครื่องดูดของเหลว (Suction)", 4050},
		},
	}

	for _, e := range equipments.Equipment {

		client.Equipment.
			Create().
			SetEQUIPMENTNAME(e.EQUIPMENTNAME).
			SetEQUIPMENTPRICE(e.EQUIPMENTPRICE).
			Save(context.Background())
	}

	// Set Companys Data
	companys := Companys{
		Company: []Company{
			Company{"YUWELL"},
			Company{"POLYGREEN"},
		},
	}

	for _, cp := range companys.Company {

		client.Company.
			Create().
			SetCOMPANYName(cp.COMPANYName).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
