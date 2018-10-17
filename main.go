package main

/*

This program is containing the code of aws-sdk for creating , describes instance and get imageId , etc.
Before using aws-go-sdk you must have add it's package by using command
go get -u github.com/aws/aws-sdk-go/...
---------------------------------------------

once you have done you can access the sdk packages methods.
---------------------------------------
-------_________------_______----------
----- / /______/-----/ ____--/---------
-----/ /--___-------/ /---/-/----------
----/ /--/_--/-----/ /---/-/-----------
---/ /____/ /-----/ /___/-/------------
--/________/-----/_______/-------------
---------------------------------------
---------------------------------------
*/

import (
	"./DBConnection"
	"./maps"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

/*type server_model struct {
	Model string `json:"model"`
	Cpu int `json:"cpu"`
	Memory string `json:"memory"`
	Cpu_credit_per_hour int `json:"cpu_credit_per_hour"`
	Storage string `json:"storage"`
}*/


type instance struct {
	Region string `json:"region"`
	Key string	`json:"key"`
	Value string `json:"value"`
	Model string `json:"model"`
	ImageId string`json:"imageId"`
}
type CredAWS struct {
	AccessKeyId string `json:"access_key_id"`
	SecretAccessKey string `json:"secret_access_key"`
	SessionToken string `json:"session_token"`
	Expiration time.Time `json:"expiration"`
}


type Tags struct {
	Key string `json:"key"`
	Value string`json:"value"`
}

func handleHome(w http.ResponseWriter ,r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(maps.CreateMap());
}

func dbCheck(w http.ResponseWriter, r *http.Request)  {
	json.NewEncoder(w).Encode(DBConnection.DbConc())
}
func findAllModels(w http.ResponseWriter, r *http.Request)  {

	w.Header().Set("Content-Type","application/json")
	fmt.Println(DBConnection.FindAllModels());
	json.NewEncoder(w).Encode(DBConnection.FindAllModels())
}

func createInstance(w http.ResponseWriter, r *http.Request)  {
	/*var person User
	_ = json.NewDecoder(r.Body).Decode(&person)*/

	w.Header().Set("Content-Type","application/json")
	/*var instanceData instance
	_ = json.NewDecoder(r.Body).Decode(&instanceData);*/

	instanceData := instance{}

	data , err := ioutil.ReadAll(r.Body);
	if err!=nil{
		panic(err.Error());
	}

	jsonErr := json.Unmarshal(data,&instanceData);
	if jsonErr !=nil{
		panic(err.Error())
	}


}

func addModel(w http.ResponseWriter, r *http.Request)  {

	w.Header().Set("Content-Type","application/json")


	Model_data := DBConnection.Server_model{}
	json.NewDecoder(r.Body).Decode(&Model_data);

	json.NewEncoder(w).Encode(DBConnection.AddModel(Model_data));

}

func findModelByModelName(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r);
	w.Header().Set("Content-Type","application/json")

	json.NewEncoder(w).Encode(DBConnection.FindModel(params["model"]));
}

func deleteModelByModelName(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r);
	json.NewEncoder(w).Encode(DBConnection.DeleteModel(params["model"]));

}
// using wait group for sync the routine with main go routine
var wg	= sync.WaitGroup{}

var i int

var Redis_Client redis.Client;

// using mutex to resolve concurrency problem
var m	= sync.RWMutex{}

func main() {



	/*//GetImageIds();
	for i=0;i<10 ;i++  {

		wg.Add(2)

		m.RLock()
		go PrintHello();

		//log.Printf("before m.lock")
		m.Lock()
		go Increment();
	}

	n:= 3;
	fmt.Println("Implementing channels")
	fmt.Println("number is",n);
	ch := make(chan int)


	go Channels.ImplementingChannels(n,ch)
	fmt.Println("After calculation",<-ch)

	wg.Wait()

	fmt.Printf("threads %v \n",runtime.GOMAXPROCS(-1));


	interfaceImpl.InterfaceImpl()*/

	Redis_Client = *CreateRedisClient();

	router := mux.NewRouter()
	router.HandleFunc("/",checkServer).Methods("GET")
	router.HandleFunc("/findAllModels", findAllModels).Methods("GET")
	router.HandleFunc("/createInstance",createInstance).Methods("POST")
	router.HandleFunc("/addModel",addModel).Methods("POST");
	router.HandleFunc("/findModel/{model}",findModelByModelName).Methods("GET");
	router.HandleFunc("/deleteModel/{model}",deleteModelByModelName).Methods("GET");
	router.HandleFunc("/fetchInstances",fetchAWSInstances).Methods("GET");
	router.HandleFunc("/createAWSCredentials/{token}",createAWSCredentials).Methods("GET")
	http.ListenAndServe(":8080",router)


}

func CreateRedisClient()  *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	if err!=nil{
		panic(err.Error());
	}else {
		fmt.Println("Success making connection with redis",pong)
	}

	return client;
}

func createAWSCredentials(w http.ResponseWriter,r *http.Request)  {
	enableCros(&w);
	params := mux.Vars(r);
	w.Header().Set("Content-Type","application/json")
	response := GetSToken(params["token"]);

	fmt.Println(response);
	Redis_Client.Set("AccessKeyId",*response.AccessKeyId,0).Err();
	Redis_Client.Set("SecretAccessKey",*response.SecretAccessKey,0).Err();
	err := Redis_Client.Set("SessionToken",*response.SessionToken,0).Err();
	if err != nil{
		panic(err)
	}

	json.NewEncoder(w).Encode(response);
}


//fetch Temp credentials from redis server
func FetchCredentialsFromRedis() CredAWS {
	AccessKeyId ,err := Redis_Client.Get("AccessKeyId").Result();
	SecretAccessKey ,err1 := Redis_Client.Get("SecretAccessKey").Result();
	SessionToken ,err2 := Redis_Client.Get("SessionToken").Result();
	if(err!=nil){
		panic(err)
	}
	if(err1!=nil){
		panic(err1)
	}
	if(err2!=nil){
		panic(err2)
	}

	Credentials_Aws := CredAWS{}

	Credentials_Aws.AccessKeyId = AccessKeyId
	Credentials_Aws.SecretAccessKey = SecretAccessKey
	Credentials_Aws.SessionToken = SessionToken

	return Credentials_Aws;

}

//enabling cross origin
func enableCros(w *http.ResponseWriter)  {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}


func checkServer(w http.ResponseWriter, r *http.Request)  {
	enableCros(&w);
	w.Header().Set("Content-Type","application/json")

	dashboardData	:= make(map[string]interface{})
	dashboardData["status"] = true;
	dashboardData["modelList"]	=	DBConnection.FindModelName();
	json.NewEncoder(w).Encode(dashboardData);
}

func fetchAWSInstances(w http.ResponseWriter, r *http.Request)  {

	enableCros(&w);
	w.Header().Set("Content-Type","application/json")

	json.NewEncoder(w).Encode(FetchInstance());
}

var counter int =0;


func PrintHello()  {
	fmt.Printf("Hello %d\n",counter)
	m.RUnlock()
	wg.Done()
}

func Increment()  {
	counter++;
	m.Unlock()
	wg.Done()
}

/*
for fetching sts token for creating session
*/
func GetSToken(token string) *sts.Credentials{

	fmt.Println(token);
	svc := sts.New(session.New())
	input := &sts.GetSessionTokenInput{
		DurationSeconds: aws.Int64(7200),
		SerialNumber:    aws.String("arn:aws:iam::294069028655:mfa/mohammad.afzal"),
		TokenCode:       aws.String(token),
	}

	result, err := svc.GetSessionToken(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case sts.ErrCodeRegionDisabledException:
				fmt.Println(sts.ErrCodeRegionDisabledException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
	}

	return result.Credentials;
}


/*
creating Aws instance having MFA enable
first you have to call get session Token and uses that credentials
*/
func CreateInstanceWithMFA()  {


	creds := GetSToken("aaa");
	log.Println("creating instance......")
	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(*creds.AccessKeyId, *creds.SecretAccessKey, *creds.SessionToken),
	})

	log.Println("created session");


	svc := ec2.New(sess, &aws.Config{
		Region:aws.String("us-east-1"),
	})

	log.Println("created svc simple")




	runResult, err := svc.RunInstances(&ec2.RunInstancesInput{
		// An Amazon Linux AMI ID for t2.micro instances in the us-west-2 region
		ImageId:      aws.String("ami-0b33d91d"),
		InstanceType: aws.String("t2.micro"),
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
				Key:   aws.String("Name"),
				Value: aws.String("MyInstance"),
			},
		},
	})
	if errtag != nil {
		log.Println("Could not create tags for instance", runResult.Instances[0].InstanceId, errtag)
		return
	}

	log.Println("Successfully tagged instance")

}

/*
create instance without MFA
*/
func CreateInstanceWithoutMFA()  {
	svc := ec2.New(session.New(&aws.Config{Region: aws.String("us-east-1")}))
	// Specify the details of the instance that you want to create.
	runResult, err := svc.RunInstances(&ec2.RunInstancesInput{
		// An Amazon Linux AMI ID for t2.micro instances in the us-west-2 region
		ImageId:      aws.String("ami-0b33d91d"),
		InstanceType: aws.String("t2.micro"),
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
				Key:   aws.String("Name"),
				Value: aws.String("test"),
			},
		},
	})
	if errtag != nil {
		log.Println("Could not create tags for instance", runResult.Instances[0].InstanceId, errtag)
		return
	}

	log.Println("Successfully tagged instance")
}


/*
For Session Creating
*/
func sessionCreatetion() *ec2.EC2 {

	creds := FetchCredentialsFromRedis();
	// Load session from shared config
	/*sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config:aws.Config{
			Region:aws.String("us-east-1"),
		},
	}))*/


	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(creds.AccessKeyId, creds.SecretAccessKey, creds.SessionToken),
	})

	if err!=nil {
		panic(err);
	}
	log.Println("created session");

	// Create new EC2 client
	svc := ec2.New(sess)

	return svc;
}


func FetchInstance()  *ec2.DescribeInstancesOutput {

	svc := sessionCreatetion();

	result, err := svc.DescribeInstances(nil)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("Success")
		return result;
	}

	return nil
}

func MonitorInstance(id *string) *ec2.MonitorInstancesOutput  {
	inst := ec2.MonitorInstancesInput{}
	inst.InstanceIds = append(inst.InstanceIds, id);

	svc := sessionCreatetion();
	result, err := svc.MonitorInstances(&inst);
	if(err!=nil){

	}else{
		return result;
	}
	return nil;
}

func GetImageIds()  {

	svc := sessionCreatetion();

	result, err := svc.DescribeImages(nil)

	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("Success", result)
	}
}

