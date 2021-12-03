package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-batch-master-record-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
)

func ConvertToBatch(raw []byte, l *logger.Logger) (*Batch, error) {
	pm := &responses.Batch{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Batch. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 1 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &Batch{
		Material:                  data.Material,
		BatchIdentifyingPlant:     data.BatchIdentifyingPlant,
		Batch:                     data.Batch,
		Supplier:                  data.Supplier,
		BatchBySupplier:           data.BatchBySupplier,
		CountryOfOrigin:           data.CountryOfOrigin,
		RegionOfOrigin:            data.RegionOfOrigin,
		MatlBatchAvailabilityDate: data.MatlBatchAvailabilityDate,
		ShelfLifeExpirationDate:   data.ShelfLifeExpirationDate,
		ManufactureDate:           data.ManufactureDate,
		CreationDateTime:          data.CreationDateTime,
		LastChangeDateTime:        data.LastChangeDateTime,
		BatchIsMarkedForDeletion:  data.BatchIsMarkedForDeletion,
	}, nil
}