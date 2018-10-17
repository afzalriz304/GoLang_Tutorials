package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)


var(

	//configure OAuth
	googleOauthConfig = oauth2.Config{
		RedirectURL:"http://localhost:8080/callback",						// it is callback url
		ClientID:os.Getenv("CLIENT_ID"),								// getting the client id from environment variable
		ClientSecret:os.Getenv("CLIENT_SECRET"),						// getting secret key from environment variable
		Scopes:[]string{"https://www.googleapis.com/auth/userinfo.email"},	// fetching data from google api
		Endpoint:google.Endpoint,
	}

	randomState = "intial";
)

// create a UserInfo struct for Api Response
type UserInfo struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Link          string `json:"link"`
	Picture       string `json:"picture"`
} 

func loginHandler(w http.ResponseWriter,r *http.Request)  {

	//create a temporary url
	url:= googleOauthConfig.AuthCodeURL(randomState)
	//redirect user to this url
	http.Redirect(w,r,url,http.StatusTemporaryRedirect)
}

// create a callback handler for redirect url
func callbackHandler(w http.ResponseWriter,r *http.Request) {
	// passing the formValues
	content, err := CallBack(r.FormValue("state"),r.FormValue("code"))
	if err !=nil{
		//redirect user to home page if any error occurs
		http.Redirect(w,r,"/",http.StatusTemporaryRedirect);
		return
	}

	// create a structure for parsing the data as json
	var userInfo UserInfo
	// Unmarshal the byte array and get the data in struct of type UserInfo
	err = json.Unmarshal(content, &userInfo)

	/*
	encode the data in the form of Json and send it to the
	client Application in the form of JSON ast.Object
	*/
	json.NewEncoder(w).Encode(userInfo);

}

func CallBack(state string, code string) ([]byte, error) {
	// checking whether the state is same or not
	if(state!=randomState){
		return nil, fmt.Errorf("Invalid State")
	}

	// fetching the token from google using .Exhange() method
	token, err := googleOauthConfig.Exchange(oauth2.NoContext,code)
	if err != nil{
		return nil, fmt.Errorf("error in fetching token")
	}

	/*
	fetching the userInfo using userInfo api of Google along
	access token to verify it on Google Authorisation server
	*/
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token="+token.AccessToken);

	if err != nil{
		return nil, fmt.Errorf("unable to fetch the user info")
	}

	defer response.Body.Close();

	// reading the response from the response body
	content, err := ioutil.ReadAll(response.Body);

	if err != nil{
		return nil, fmt.Errorf("unable to parse the content")
	}

	// return the response as byte array
	return content,nil
}



func homeHandler(w http.ResponseWriter,r *http.Request)  {

	company:="hashedIn"
	var html	=	`<html>
						<body>
							<a href="/login">Google Auth by `+company+`</a>
						</body>
					</html>`

	fmt.Fprint(w,html);
}



func main()  {
	router:= mux.NewRouter();
	log.Printf("Server starts at port 8080")
	router.HandleFunc("/",homeHandler)
	router.HandleFunc("/callback",callbackHandler)
	router.HandleFunc("/login",loginHandler)
	http.ListenAndServe(":8080",router);
}
