package main

import (
	sap_api_caller "sap-api-integrations-batch-master-record-reads/SAP_API_Caller"
	sap_api_input_reader "sap-api-integrations-batch-master-record-reads/SAP_API_Input_Reader"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

func main() {
	l := logger.NewLogger()
	fr := sap_api_input_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs/SDC_Batch_Master_Record_sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/", l,
	)

	accepter := inoutSDC.Accepter
	if len(accepter) == 0 || accepter[0] == "All" {
		accepter = []string{
			"Batch",
		}
	}

	caller.AsyncGetBatchMasterRecord(
		inoutSDC.Batch.Material,
		inoutSDC.Batch.BatchIdentifyingPlant,
		inoutSDC.Batch.Batch,
		accepter,
	)
}
