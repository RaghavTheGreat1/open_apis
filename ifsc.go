package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func ifscToBankInfo() {
	var ifscCode string
	fmt.Print("Enter IFSC Code: ")
	fmt.Scanf("%s", &ifscCode)
	bankDetails, _ := fetchBankDetails(ifscCode)
	if bankDetails != nil {
		fmt.Printf("%+v", *bankDetails)
	}
}

const url = "https://ifsc.razorpay.com/"

func fetchBankDetails(ifscCode string) (*BankDetails, error) {
	bankDetails := &BankDetails{}

	apiResponse, err := http.Get(url + strings.ToUpper(ifscCode))

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer apiResponse.Body.Close()

	body := apiResponse.Body

	bytesData, _ := io.ReadAll(body)

	parseErr := json.Unmarshal(bytesData, bankDetails)

	if parseErr != nil {
		fmt.Println(parseErr)
		return nil, parseErr
	}

	return bankDetails, nil
}

type BankDetails struct {
	Branch   string
	Centre   string
	District string
	State    string
	Address  string
	Contact  string
	IMPS     bool
	City     string
	UPI      bool
	MICR     string
	RTGS     bool
	NEFT     bool
	Swift    string
	ISO3166  string
	Bank     string
	BankCode string
	IFSC     string
}

/*

"BRANCH": "Delhi Nagrik Sehkari Bank IMPS",
"CENTRE": "DELHI",
"DISTRICT": "DELHI",
"STATE": "MAHARASHTRA",
"ADDRESS": "720, NEAR GHANTAGHAR, SUBZI MANDI, DELHI - 110007",
"CONTACT": "+919560344685",
"IMPS": true,
"CITY": "MUMBAI",
"UPI": true,
"MICR": "110196002",
"RTGS": true,
"NEFT": true,
"SWIFT": "",
"ISO3166": "IN-MH",
"BANK": "Delhi Nagrik Sehkari Bank",
"BANKCODE": "DENS",
"IFSC": "YESB0DNB002"

*/
