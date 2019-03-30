package NtmResource

import (
	"errors"
	"github.com/PharbersDeveloper/NtmPods/NtmDataStorage"
	"github.com/PharbersDeveloper/NtmPods/NtmModel"
	"reflect"
	"net/http"

	"github.com/alfredyang1986/BmServiceDef/BmDataStorage"
	"github.com/manyminds/api2go"
)

type NtmHospitalSalesReportResource struct {
	NtmHospitalSalesReportStorage       *NtmDataStorage.NtmHospitalSalesReportStorage
	NtmDestConfigStorage 				*NtmDataStorage.NtmDestConfigStorage
	NtmGoodsConfigStorage 		        *NtmDataStorage.NtmGoodsConfigStorage
	NtmSalesReportStorage               *NtmDataStorage.NtmSalesReportStorage
}

func (c NtmHospitalSalesReportResource) NewHospitalSalesReportResource(args []BmDataStorage.BmStorage) *NtmHospitalSalesReportResource {
	var hsr  *NtmDataStorage.NtmHospitalSalesReportStorage
	var dc *NtmDataStorage.NtmDestConfigStorage
	var gc *NtmDataStorage.NtmGoodsConfigStorage
	var sr *NtmDataStorage.NtmSalesReportStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "NtmHospitalSalesReportStorage" {
			hsr = arg.(*NtmDataStorage.NtmHospitalSalesReportStorage)
		} else if tp.Name() == "NtmDestConfigStorage" {
			dc = arg.(*NtmDataStorage.NtmDestConfigStorage)
		} else if tp.Name() == "NtmGoodsConfigStorage" {
			gc = arg.(*NtmDataStorage.NtmGoodsConfigStorage)
		} else if tp.Name() == "NtmSalesReportStorage" {
			sr = arg.(*NtmDataStorage.NtmSalesReportStorage)
		}
	}
	return &NtmHospitalSalesReportResource{
		NtmHospitalSalesReportStorage: hsr,
		NtmDestConfigStorage: dc,
		NtmGoodsConfigStorage: gc,
		NtmSalesReportStorage: sr,
	}
}

// FindAll SalesConfigs
func (c NtmHospitalSalesReportResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	salesreportID, dcok := r.QueryParams["salesreportsID"]

	if dcok {
		modelRootID := salesreportID[0]
		modelRoot, err := c.NtmSalesReportStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}
		r.QueryParams["ids"] = modelRoot.HospitalSalesReportIDs

		model := c.NtmHospitalSalesReportStorage.GetAll(r, -1,-1)


		if err != nil {
			return &Response{}, nil
		}
		return &Response{Res: model}, nil
	}

	var result []NtmModel.Hospitalsalesreport

	models := c.NtmHospitalSalesReportStorage.GetAll(r, -1, -1)

	for _, model := range models {
		result = append(result, *model)
	}

	return &Response{Res: result}, nil
}

// FindOne choc
func (c NtmHospitalSalesReportResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.NtmHospitalSalesReportStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c NtmHospitalSalesReportResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(NtmModel.Hospitalsalesreport)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := c.NtmHospitalSalesReportStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c NtmHospitalSalesReportResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.NtmHospitalSalesReportStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c NtmHospitalSalesReportResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(NtmModel.Hospitalsalesreport)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := c.NtmHospitalSalesReportStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}