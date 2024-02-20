package repository

import (
	"auth-services/model"
	"fmt"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
)

type CustomerRepository interface {
	Login(model.Login) (model.LoginResponse, *resty.Response, error)
}
type customerRepository struct {
	client *resty.Client
}

// Login implements CustomerRepository.
func (r customerRepository) Login(req model.Login) (model.LoginResponse, *resty.Response, error) {
	var result model.LoginResponse
	resp, err := r.client.R().
		SetBody(req).
		SetResult(&result).
		Post(fmt.Sprintf("%s%s", os.Getenv("SERVICE_HOST_CUSTOMER"), "/api/login"))

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

func NewCustomerRepo() CustomerRepository {
	return customerRepository{
		client: resty.New().SetTimeout(30 * time.Second).SetRetryCount(3),
	}
}
