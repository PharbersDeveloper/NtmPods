package NtmResource

import (
	"errors"
	"github.com/PharbersDeveloper/NtmPods/NtmDataStorage"
	"github.com/PharbersDeveloper/NtmPods/NtmModel"
	"github.com/alfredyang1986/BmServiceDef/BmDataStorage"
	"github.com/manyminds/api2go"
	"net/http"
	"reflect"
	"strconv"
)

type NtmRegionConfigResource struct {
	NtmRegionConfigStorage		*NtmDataStorage.NtmRegionConfigStorage
	NtmRegionStorage			*NtmDataStorage.NtmRegionStorage
	NtmDestConfigStorage 		*NtmDataStorage.NtmDestConfigStorage
}

func (s NtmRegionConfigResource) NewRegionConfigResource(args []BmDataStorage.BmStorage) *NtmRegionConfigResource {
	var rcs *NtmDataStorage.NtmRegionConfigStorage
	var rs *NtmDataStorage.NtmRegionStorage
	var dcs *NtmDataStorage.NtmDestConfigStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "NtmRegionConfigStorage" {
			rcs = arg.(*NtmDataStorage.NtmRegionConfigStorage)
		} else if tp.Name() == "NtmRegionStorage" {
			rs = arg.(*NtmDataStorage.NtmRegionStorage)
		} else if tp.Name() == "NtmDestConfigStorage" {
			dcs = arg.(*NtmDataStorage.NtmDestConfigStorage)
		}
	}
	return &NtmRegionConfigResource{
		NtmRegionStorage: rs,
		NtmRegionConfigStorage: rcs,
		NtmDestConfigStorage: dcs,
	}
}

func (s NtmRegionConfigResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	destConfigsID, dcok := r.QueryParams["destConfigsID"]
	var result []NtmModel.RegionConfig

	if dcok {
		modelRootID := destConfigsID[0]
		modelRoot, err := s.NtmDestConfigStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}
		model, err := s.NtmRegionConfigStorage.GetOne(modelRoot.DestID)
		if err != nil {
			return &Response{}, nil
		}
		result = append(result, model)
		return &Response{Res: result}, nil
	}

	models := s.NtmRegionConfigStorage.GetAll(r, -1, -1)
	for _, model := range models {
		result = append(result, *model)
	}

	return &Response{Res: result}, nil
}

// PaginatedFindAll can be used to load models in chunks
func (s NtmRegionConfigResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	var (
		result                      []NtmModel.RegionConfig
		number, size, offset, limit string
	)

	numberQuery, ok := r.QueryParams["page[number]"]
	if ok {
		number = numberQuery[0]
	}
	sizeQuery, ok := r.QueryParams["page[size]"]
	if ok {
		size = sizeQuery[0]
	}
	offsetQuery, ok := r.QueryParams["page[offset]"]
	if ok {
		offset = offsetQuery[0]
	}
	limitQuery, ok := r.QueryParams["page[limit]"]
	if ok {
		limit = limitQuery[0]
	}

	if size != "" {
		sizeI, err := strconv.ParseInt(size, 10, 64)
		if err != nil {
			return 0, &Response{}, err
		}

		numberI, err := strconv.ParseInt(number, 10, 64)
		if err != nil {
			return 0, &Response{}, err
		}

		start := sizeI * (numberI - 1)
		for _, iter := range s.NtmRegionConfigStorage.GetAll(r, int(start), int(sizeI)) {
			result = append(result, *iter)
		}

	} else {
		limitI, err := strconv.ParseUint(limit, 10, 64)
		if err != nil {
			return 0, &Response{}, err
		}

		offsetI, err := strconv.ParseUint(offset, 10, 64)
		if err != nil {
			return 0, &Response{}, err
		}

		for _, iter := range s.NtmRegionConfigStorage.GetAll(r, int(offsetI), int(limitI)) {
			result = append(result, *iter)
		}
	}

	in := NtmModel.RegionConfig{}
	count := s.NtmRegionConfigStorage.Count(r, in)

	return uint(count), &Response{Res: result}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the model with the given ID, otherwise an error
func (s NtmRegionConfigResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	model, err := s.NtmRegionConfigStorage.GetOne(ID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}

	if model.RegionID != "" {
		regionModel, err := s.NtmRegionStorage.GetOne(model.RegionID)
		if err != nil {
			return &Response{}, err
		}
		model.Region = &regionModel
	}

	return &Response{Res: model}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (s NtmRegionConfigResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(NtmModel.RegionConfig)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := s.NtmRegionConfigStorage.Insert(model)
	model.ID = id

	return &Response{Res: model, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (s NtmRegionConfigResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := s.NtmRegionConfigStorage.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

//Update stores all changes on the model
func (s NtmRegionConfigResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(NtmModel.RegionConfig)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := s.NtmRegionConfigStorage.Update(model)
	return &Response{Res: model, Code: http.StatusNoContent}, err
}
