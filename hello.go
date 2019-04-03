package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Name struct {
	LastName   string
	FirstName  string
	MiddleName string
}
type Address struct {
	Address1 string
	Address2 string
	Address3 string
}
type Report struct {
	HeaderRecord struct {
		HeaderRecordIndicator string
		SupervisingAgency     string
		InstitutionCode       string
		RecordDate            string
		ReportType            string
		FormatCode            string
		SubmissionType        string
	}
	DetailRecord struct {
		DetailRecordIndicator          string
		TransactionDate                int
		TransactionCode                string
		TransactionReferenceNumber     string
		AccountNumber                  string
		OldAccountNumber               string
		TransactionAmount              float64
		TransactionAmountFX            float64
		FxCurrencyCode                 string
		PurposeOfTransaction           string
		EffectivityDate                int
		ExpiryDate                     int
		AmountOfClaim                  float64
		NumberOfShares                 int
		NetAssetValue                  int
		Address                        Address
		CountryCodeOfCorrespondentBank int
	}
}

// SubjectData struct {
// 	AccountHolder struct {
// 		CustomerReferenceNumber string
// 		NameFlag                string
// 		Name                    Name
// 		Address                 Address
// 		RegistrationDate        int
// 		PlaceOfBirth            string
// 		Nationality             string
// 		IdType                  string
// 		IdentificationNumber    string
// 		TelephoneNumber         string
// 		NatureOfBusiness        string
// 	}
// 	Benefeciary struct {
// 		CustomerReferenceNumber string
// 		NameFlag                string
// 		Name                    Name
// 		Address                 Address
// 		AccountNumber           string
// 		RegistrationDate        int
// 		PlaceOfBirth            string
// 		Nationality             string
// 		IdType                  string
// 		IdentificationNumber    string
// 		TelephoneNumber         string
// 		NatureOfBusiness        string
// 	}
// 	CounterParty struct {
// 		CustomerReferenceNumber string
// 		NameFlag                string
// 		Name                    Name
// 		Address                 Address
// 		AccountNumber           string
// 	}
// 	OtherParticipant struct {
// 		CustomerReferenceNumber string
// 		NameFlag                string
// 		Name                    Name
// 		Address                 Address
// 		AccountNumber           string
// 	}
// 	Issuer struct {
// 		CustomerReferenceNumber string
// 		NameFlag                string
// 		Name                    Name
// 		Address                 Address
// 		AccountNumber           string
// 	}
// 	Transactor struct {
// 		CustomerReferenceNumber string
// 		NameFlag                string
// 		Name                    Name
// 		Address                 Address
// 		AccountNumber           string
// 	}
// 	SubjectOfSuspicion struct {
// 		CustomerReferenceNumber string
// 		NameFlag                string
// 		Name                    Name
// 		Address                 Address
// 		AccountNumber           string
// 		RegistrationDate        int
// 		PlaceOfBirth            string
// 		Nationality             string
// 		idType                  string
// 		IdentificationNumber    string
// 		TelephoneNumber         string
// 		NatureOfBusiness        string
// 	}
// }
// DetailOfSuspicion struct {
// 	Reason    string
// 	Narrative string
// }
// TrailerRecord struct {
// 	PhpAmoungTotal             float64
// 	RecordTotalOfBatchToBeSent int
// }

func main() {
	csvFile, err := os.Open("./data.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var rep Report
	var reports []Report

	for _, each := range csvData {
		var header = each[0]

		if header == "H" {
			rep.HeaderRecord.HeaderRecordIndicator = each[0]
			rep.HeaderRecord.SupervisingAgency = each[1]
			rep.HeaderRecord.InstitutionCode = each[2]
			rep.HeaderRecord.RecordDate = each[3]
			rep.HeaderRecord.ReportType = each[4]
			rep.HeaderRecord.FormatCode = each[5]
			rep.HeaderRecord.SubmissionType = each[6]
		}
		if header == "D" {
			rep.DetailRecord.DetailRecordIndicator = each[0]
			rep.DetailRecord.TransactionDate, _ = strconv.Atoi(each[1])
			rep.DetailRecord.TransactionCode = each[2]
			rep.DetailRecord.TransactionReferenceNumber = each[3]
			rep.DetailRecord.AccountNumber = each[4]
			rep.DetailRecord.OldAccountNumber = each[5]
			rep.DetailRecord.TransactionAmount, _ = strconv.ParseFloat(each[6], 64)
			rep.DetailRecord.TransactionAmountFX, _ = strconv.ParseFloat(each[7], 64)
			rep.DetailRecord.FxCurrencyCode = each[8]
			rep.DetailRecord.PurposeOfTransaction = each[9]
			rep.DetailRecord.EffectivityDate, _ = strconv.Atoi(each[10])
			rep.DetailRecord.ExpiryDate, _ = strconv.Atoi(each[11])
			rep.DetailRecord.AmountOfClaim, _ = strconv.ParseFloat(each[12], 64)
			rep.DetailRecord.NumberOfShares, _ = strconv.Atoi(each[13])
			rep.DetailRecord.NetAssetValue, _ = strconv.Atoi(each[14])
			rep.DetailRecord.Address.Address1 = each[15]
			rep.DetailRecord.Address.Address2 = each[16]
			rep.DetailRecord.Address.Address3 = each[17]
			rep.DetailRecord.CountryCodeOfCorrespondentBank, _ = strconv.Atoi(each[18])
		}

		reports = append(reports, rep)
	}

	// Convert to JSON
	jsonData, err := json.Marshal(reports)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(jsonData))

	jsonFile, err := os.Create("./data.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	jsonFile.Write(jsonData)
	jsonFile.Close()
}
