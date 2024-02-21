package repository

import (
	"auth-services/model"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

type CategoryRepo interface {
	AddCategory(model.RequestAdd) (model.ResponseCategory, *resty.Response, error)
	UpdateCategory(model.RequestUpdateCategory) (model.ResponseSuccess, *resty.Response, error)
	ListCategory() ([]model.ResponseCategory, *resty.Response, error)
	DetailCategory(model.RequesListCategoryById) (model.ResponseCategory, *resty.Response, error)
}

type categoryRepo struct {
	client *resty.Client
}

// DetailCategory implements CategoryRepo.
func (r categoryRepo) DetailCategory(req model.RequesListCategoryById) (model.ResponseCategory, *resty.Response, error) {
	var result model.ResponseCategory
	resp, err := r.client.R().
		SetQueryParams(map[string]string{
			"category_id": req.Id,
		}).
		SetResult(&result).
		Get("http://localhost:8085/api/detail")
		//Get(fmt.Sprintf("%s%s", os.Getenv("SERVICE_HOST_RATE"), "/api/detail-lps"))

	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()

	return result, resp, err
}

// ListCategory implements CategoryRepo.
func (r categoryRepo) ListCategory() ([]model.ResponseCategory, *resty.Response, error) {
	var result []model.ResponseCategory
	resp, err := r.client.R().
		SetResult(&result).
		Get("http://localhost:8085/api/list")
		//Get(fmt.Sprintf("%s%s", os.Getenv("SERVICE_HOST_RATE"), "/api/list-libor"))

	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()

	return result, resp, err
}

// UpdateCategory implements CategoryRepo.
func (r categoryRepo) UpdateCategory(req model.RequestUpdateCategory) (model.ResponseSuccess, *resty.Response, error) {
	var result model.ResponseSuccess
	resp, err := r.client.R().
		SetBody(req).
		SetResult(&result).
		Post("http://localhost:8085/api/update")
		//Post(fmt.Sprintf("%s%s", os.Getenv("SERVICE_HOST_CUSTOMER"), "/api/add"))

	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()

	return result, resp, err
}

// AddCategory implements CategoryRepo.
func (r categoryRepo) AddCategory(req model.RequestAdd) (model.ResponseCategory, *resty.Response, error) {
	var result model.ResponseCategory
	resp, err := r.client.R().
		SetBody(req).
		SetResult(&result).
		Post("http://localhost:8085/api/add")
		//Post(fmt.Sprintf("%s%s", os.Getenv("SERVICE_HOST_CUSTOMER"), "/api/add"))

	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()

	return result, resp, err
}

func NewCategoryRepo() CategoryRepo {
	return categoryRepo{
		client: resty.New().SetTimeout(30 * time.Second).SetRetryCount(3),
	}
}
