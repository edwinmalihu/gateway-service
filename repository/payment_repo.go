package repository

import (
	"auth-services/model"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

type PaymentRepo interface {
	AddPayment(model.RequestAddPayment) (model.ResponseSuccessPayment, *resty.Response, error)
	DeletePayment(model.RequestByIdPayment) (model.Response, *resty.Response, error)
}

type paymentRepo struct {
	client *resty.Client
}

// AddPayment implements PaymentRepo.
func (r paymentRepo) AddPayment(req model.RequestAddPayment) (model.ResponseSuccessPayment, *resty.Response, error) {
	var result model.ResponseSuccessPayment
	resp, err := r.client.R().
		SetBody(req).
		SetResult(&result).
		Post("http://localhost:8088/api/add")
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

// DeletePayment implements PaymentRepo.
func (r paymentRepo) DeletePayment(req model.RequestByIdPayment) (model.Response, *resty.Response, error) {
	var result model.Response
	resp, err := r.client.R().
		SetQueryParams(map[string]string{
			"payment_id": req.PaymentID,
		}).
		SetResult(&result).
		Delete("http://localhost:8088/api/delete")
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

func NewPaymentRepo() PaymentRepo {
	return paymentRepo{
		resty.New().SetTimeout(30 * time.Second).SetRetryCount(3),
	}
}
