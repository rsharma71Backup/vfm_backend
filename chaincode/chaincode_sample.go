package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	//"time"
	//"strings"
	//"reflect"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

var userIndexStr = "_userindex" //name for the key/value that will store all users

type User struct {
	Id        int    `json:"id"`
	UserType  string `json:"usertype"`
	FisrtName string `json:"firstname"` //the fieldtags of user are needed to store in the ledger
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	//ReTypePassword string `json:"retypepassword"`
	Operationalemail         string `json:"operationalemail"`
	Phone                    int    `json:"phone"`
	RelationshipManagerEmail string `json:"relationshipmanageremail"`
	CustomersLimit           int    `json:"customerslimit"`
	FeePercentage            int    `json:"feepercentage"`
	InterestEarning          int    `json:"interestearning"`
	AccountNo                int    `json:"accountno"`
	IfscCode                 string `json:"ifsccode"`
	Pan                      string `json:"pan"`
	Address                  string `json:"address"`
}

type AllUsers struct {
	Userlist []User `json:"userlist"` // contains array of users
}

type SessionAunthentication struct {
	Token string `json:"token"` //the fieldtags of user seesion are needed to store in the ledger
	Email string `json:"email"`
}
type Session struct {
	StoreSession []SessionAunthentication `json:"session"` //contains array of users session
}

type SimpleChaincode struct {
}

// Main Function

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init Function - reset all the things
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	//_, args := stub.GetFunctionAndParameters()
	var Aval int
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	// Initialize the chaincode
	Aval, err = strconv.Atoi(args[0])
	if err != nil {
		return nil, errors.New("Expecting integer value for asset holding")
	}

	// Write the state to the ledger
	err = stub.PutState("abc", []byte(strconv.Itoa(Aval))) //making a test var "abc" to read/write into ledger to test the network
	if err != nil {
		return nil, err
	}

	var empty []string
	jsonAsBytes, _ := json.Marshal(empty) //marshal an emtpy array of strings to clear the index
	err = stub.PutState(userIndexStr, jsonAsBytes)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" { //initialize the chaincode state, used as reset
		return t.Init(stub, "init", args)
	} else if function == "write" {
		return t.write(stub, args) //writes a value to the chaincode state

	} else if function == "registerUser" { //writes user registraion values to ledger
		return t.registerUser(stub, args)

	} else if function == "Delete" { //deletes an entity from its state
		return t.Delete(stub, args)

	} else if function == "SaveSession" { //writes user session details to chaincode state
		return t.SaveSession(stub, args)

	}

	fmt.Println("invoke did not find func: " + function)

	return nil, errors.New("Received unknown function invocation: " + function)
}

// write - invoke function to write key/value pair
func (t *SimpleChaincode) write(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var key, value string
	var err error
	fmt.Println("running write()")

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the key and value to set")
	}

	key = args[0]
	value = args[1]
	err = stub.PutState(key, []byte(value)) //write the variable into the chaincode state
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	if function == "readuser" { //read values for particular keys
		return t.readuser(stub, args)
	} else if function == "login" { //authenticate user
		return t.login(stub, args)

	} else if function == "auntheticatetoken" { //authenticate user's token for session
		return t.SetUserForSession(stub, args)

	}
	fmt.Println("query did not find func: " + function)

	return nil, errors.New("Received unknown function query: " + function)
}

// read - query function to read key/value pair

func (t *SimpleChaincode) readuser(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var name, jsonResp string
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the var to query")
	}

	name = args[0]
	valAsbytes, err := stub.GetState(name) //get the key value from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + name + "\"}"
		return nil, errors.New(jsonResp)
	}

	return valAsbytes, nil //send it onward
}

//HANDLES user registration and writes data in chaincode state to ledger
func (t *SimpleChaincode) registerUser(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var err error

	if len(args) != 16 {
		return nil, errors.New("Incorrect number of arguments. Expecting 8")
	}

	//input sanitation
	fmt.Println("- start registration")
	if len(args[0]) <= 0 {
		return nil, errors.New("0th argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return nil, errors.New("1st argument must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return nil, errors.New("2nd argument must be a non-empty string")
	}
	if len(args[3]) <= 0 {
		return nil, errors.New("3rd argument must be a non-empty string")
	}
	if len(args[4]) <= 0 {
		return nil, errors.New("4th argument must be a non-empty string")
	}
	if len(args[5]) <= 0 {
		return nil, errors.New("5th argument must be a non-empty string")
	}

	if len(args[8]) <= 0 {
		return nil, errors.New("8th argument must be a non-empty string")
	}
	if len(args[13]) <= 0 {
		return nil, errors.New("13th argument must be a non-empty string")
	}
	if len(args[14]) <= 0 {
		return nil, errors.New("14th argument must be a non-empty string")
	}
	if len(args[15]) <= 0 {
		return nil, errors.New("15th argument must be a non-empty string")
	}
	user := User{}
	user.Id, err = strconv.Atoi(args[0])
	if err != nil {
		return nil, errors.New("Failed to get id as cannot convert it to int")
	}
	user.UserType = args[1]
	user.FisrtName = args[2]
	user.LastName = args[3]
	user.Email = args[4]
	user.Password = args[5]
	//user.ReTypePassword=args[6]
	user.Operationalemail = args[6]
	user.Phone, err = strconv.Atoi(args[7])
	if err != nil {
		return nil, errors.New("Failed to get phone as cannot convert it to int")
	}
	user.RelationshipManagerEmail = args[8]
	user.CustomersLimit, err = strconv.Atoi(args[9])
	if err != nil {
		return nil, errors.New("Failed to get CustomersLimit as cannot convert it to int")
	}
	user.FeePercentage, err = strconv.Atoi(args[10])
	if err != nil {
		return nil, errors.New("Failed to get FeePercentage as cannot convert it to int")
	}
	user.InterestEarning, err = strconv.Atoi(args[11])
	if err != nil {
		return nil, errors.New("Failed to get InterestEarning as cannot convert it to int")
	}
	user.AccountNo, err = strconv.Atoi(args[12])
	if err != nil {
		return nil, errors.New("Failed to get AccountNo as cannot convert it to int")
	}
	user.IfscCode = args[13]
	user.Pan = args[14]

	user.Address = args[15]

	fmt.Println("user", user)
	// get users data from chaincode
	UserAsBytes, err := stub.GetState("getvfmuser")
	if err != nil {
		return nil, errors.New("Failed to get users")
	}
	var allusers AllUsers
	json.Unmarshal(UserAsBytes, &allusers) //un stringify it aka JSON.parse()

	allusers.Userlist = append(allusers.Userlist, user)
	fmt.Println("allusers", allusers.Userlist) //append usersdetails to allusers[]
	fmt.Println("! appended user to allusers")
	jsonAsBytes, _ := json.Marshal(allusers)
	fmt.Println("json", jsonAsBytes)
	err = stub.PutState("getvfmuser", jsonAsBytes) //rewrite allusers[]
	if err != nil {
		return nil, err
	}
	fmt.Println("- end user_register")
	return nil, nil
}

//login user
func (t *SimpleChaincode) login(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var err error

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2")
	}

	//input sanitation
	fmt.Println("- login")
	if len(args[0]) <= 0 {
		return nil, errors.New("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return nil, errors.New("2nd argument must be a non-empty string")
	}

	emailid := args[0]

	password := args[1]

	UserAsBytes, err := stub.GetState("getvfmuser")
	if err != nil {
		return nil, errors.New("Failed to get users")
	}
	var allusers AllUsers
	json.Unmarshal(UserAsBytes, &allusers) //un stringify it aka JSON.parse()

	for i := 0; i < len(allusers.Userlist); i++ {

		//comparing emailid and password for login
		if allusers.Userlist[i].Email == emailid && allusers.Userlist[i].Password == password {

			return []byte(allusers.Userlist[i].Email), nil
		}
	}
	return nil, nil
}

// Delete - remove a key/value pair from state
func (t *SimpleChaincode) Delete(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	name := args[0]
	err := stub.DelState(name) //remove the key from chaincode state
	if err != nil {
		return nil, errors.New("Failed to delete state")
	}

	//get the user index
	userAsBytes, err := stub.GetState(userIndexStr)
	if err != nil {
		return nil, errors.New("Failed to get array index")
	}
	var userIndex []string
	json.Unmarshal(userAsBytes, &userIndex) //un stringify it aka JSON.parse()

	//remove user from index
	for i, val := range userIndex {
		fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for " + name)
		if val == name { //find the correct index

			userIndex = append(userIndex[:i], userIndex[i+1:]...) //remove it
			for x := range userIndex {                            //debug prints...
				fmt.Println(string(x) + " - " + userIndex[x])
			}
			break
		}
	}
	jsonAsBytes, _ := json.Marshal(userIndex) //save new index
	err = stub.PutState(userIndexStr, jsonAsBytes)
	return nil, nil
}

// save user's token for session
func (t *SimpleChaincode) SaveSession(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var err error
	fmt.Println("running savesession")

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2.")
	}
	if len(args[0]) <= 0 {
		return nil, errors.New("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return nil, errors.New("2nd argument must be a non-empty string")
	}
	authsession := SessionAunthentication{}
	authsession.Token = args[0]
	authsession.Email = args[1]
	//get session empty[]
	UserAsBytes, err := stub.GetState("savesessionvfm")
	if err != nil {
		return nil, errors.New("Failed to get users session")
	}
	var session Session
	json.Unmarshal(UserAsBytes, &session) //un stringify it aka JSON.parse()

	session.StoreSession = append(session.StoreSession, authsession)
	fmt.Println("allsessions", session.StoreSession) //append each users session to allsession[]
	fmt.Println("! appended user to allsessions")
	jsonAsBytes, _ := json.Marshal(session)
	fmt.Println("json", jsonAsBytes)
	err = stub.PutState("savesessionvfm", jsonAsBytes) //rewrite allsession[]
	if err != nil {
		return nil, err
	}
	fmt.Println("- end save session")
	return nil, nil
}

//Authenticate token and set user for session
func (t *SimpleChaincode) SetUserForSession(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var token string
	var err error
	fmt.Println("running SetUserForSession()")

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2.")
	}
	token = args[0]

	UserAsBytes, err := stub.GetState("savesessionvfm")
	if err != nil {
		return nil, errors.New("failed to get sessions")
	}
	var session Session
	json.Unmarshal(UserAsBytes, &session)
	for i := 0; i < len(session.StoreSession); i++ {
		if session.StoreSession[i].Token == token {

			return []byte(session.StoreSession[i].Email), nil
		}
	}
	return nil, nil
}
