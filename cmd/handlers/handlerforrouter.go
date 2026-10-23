package handlers

import (
	"encoding/json/v2"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

// GetTypeIncidentDocument получает тип документа описывающего инцидент
func GetTypeIncidentDocument(data []byte) (datamodels.IrpIncidentType, error) {
	caseTypeElement := datamodels.IrpIncidentType{}

	if err := json.Unmarshal(data, &caseTypeElement); err != nil {
		return caseTypeElement, supportingfunctions.CustomError(err)
	}

	return caseTypeElement, nil
}
