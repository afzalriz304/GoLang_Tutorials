package DBConnection

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris/core/errors"
	"log"
)



//creating connection to Database
func DbConc() *sql.DB  {
	db, err := sql.Open("mysql",
		"root:root@tcp(127.0.0.1:3306)/server_monitoring")
	if err != nil {
		errors.New("failed to connect")
		log.Fatal(err)
	}
	fmt.Println("Successfully makes the connection")
	return db;
}

type server_model struct {
	Model string `json:"model"`
	Cpu int		`json:"cpu"`
	Memory string	`json:"memory"`
	Cpu_credit_per_hour int`json:"cpu_credit_per_hour"`
	Storage string`json:"storage"`
}

func InsertDemo()  {
	db := DbConc();
	insert, err := db.Query("INSERT into new_table VALUES (122)")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	defer insert.Close()

	fmt.Println("Successfully entered")
}

type user struct {
	userid int
}


func FindUser()  {
	db := DbConc();

	var (
		userid int
	)
	result , err := db.Query("select * from new_table")
	if err != nil{
		panic(err.Error())
	}
	defer db.Close()
	//fmt.Println(result)

	for result.Next(){
		result.Scan(&userid)
		fmt.Println(userid)
	}
	defer result.Close()
}



func FindAllModels() []server_model {
	var (
		model string
		cpu int
		memory string
		cpu_credit_per_hour int
		storage string
	)

	S_model := server_model{}
	models	:=	[]server_model{}

	db := DbConc();
	result, err := db.Query("SELECT * FROM server_model");
	if err != nil {
		errors.New("failed to get data")
		panic(err.Error())
	}

	defer db.Close()

	for result.Next(){
		err := result.Scan(&model,&cpu,&memory,&cpu_credit_per_hour,&storage);
		if err != nil {
			errors.New("failed to scan data")
			panic(err.Error())
		}
		S_model.Model	=	model
		S_model.Cpu		=	cpu
		S_model.Storage	=	storage
		S_model.Memory	=	memory
		S_model.Cpu_credit_per_hour	=	cpu_credit_per_hour

		models	=	append(models,S_model);

		fmt.Println(models)
	}

	defer result.Close()

	/*formattedData := make(map[string][]server_model);
	formattedData["data"]	=	models;*/
	//response :=formatDbData(models)

	return models;
}

type instance struct {
	Region string `json:"region"`
	Key string	`json:"key"`
	Value string `json:"value"`
	Model string `json:"model"`
	ImageId string`json:"imageId"`
}

type Tags struct {
	Key string `json:"key"`
	Value string`json:"value"`
}

/*func CreateInstance(instanceDetails *instance)  {


	/*svc := ec2.New(session.New(&aws.Config{Region: aws.String(instanceDetails.region)}))
	// Specify the details of the instance that you want to create.
	runResult, err := svc.RunInstances(&ec2.RunInstancesInput{
		// An Amazon Linux AMI ID for t2.micro instances in the us-west-2 region
		ImageId:      aws.String(instanceDetails.imageId),
		InstanceType: aws.String(instanceDetails.model),
		MinCount:     aws.Int64(1),
		MaxCount:     aws.Int64(1),
	})

	if err != nil {
		log.Println("Could not create instance", err)
		return
	}

	log.Println("Created instance", *runResult.Instances[0].InstanceId)

	// Add tags to the created instance
	_ , errtag := svc.CreateTags(&ec2.CreateTagsInput{
		Resources: []*string{runResult.Instances[0].InstanceId},
		Tags: []*ec2.Tag{
			{
				Key:   aws.String(instanceDetails.tags.key),
				Value: aws.String(instanceDetails.tags.value),
			},
		},
	})
	if errtag != nil {
		log.Println("Could not create tags for instance", runResult.Instances[0].InstanceId, errtag)
		return
	}

	log.Println("Successfully tagged instance")
}
*/

