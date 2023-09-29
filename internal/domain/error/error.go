package error

import "net/http"

const (
	DefaultErrorCaseCode = "00"
)

const (
	BadRequest                  = "Bad Request"
	InternalServerError         = "internal server error"
	InvalidFormat               = "Invalid Format"
	InvalidMandatory            = "Invalid Mandatory"
	InvalidCardAccountCustomer  = "Invalid Card/Account/Customer"
	DuplicatePartnerReferenceNo = "Duplicate partnerReferenceNo"
	PartnerNotFound             = "Partner Not Found"
	CouponNotFound              = "Coupon Not Found"
)

var (
	ErrorMapCaseCode = map[string]string{
		BadRequest:                  "00",
		InternalServerError:         "00",
		InvalidFormat:               "01",
		InvalidMandatory:            "02",
		InvalidCardAccountCustomer:  "11",
		DuplicatePartnerReferenceNo: "01",
		PartnerNotFound:             "16",
		CouponNotFound:              "17",
	}

	ErrorMapMessage = map[string]string{
		BadRequest:                  "Bad Request",
		InternalServerError:         "internal server error",
		InvalidFormat:               "Invalid Field Format %v",
		InvalidMandatory:            "Invalid Mandatory Field %v",
		InvalidCardAccountCustomer:  "Invalid Card/Account/Customer",
		DuplicatePartnerReferenceNo: "Duplicate partnerReferenceNo",
		PartnerNotFound:             "Partner Not Found",
		CouponNotFound:              "Coupon Type Not Found",
	}

	ErrorMapHttpCode = map[string]int{
		BadRequest:                  http.StatusBadRequest,
		InternalServerError:         http.StatusInternalServerError,
		InvalidFormat:               http.StatusBadRequest,
		InvalidMandatory:            http.StatusBadRequest,
		InvalidCardAccountCustomer:  http.StatusNotFound,
		DuplicatePartnerReferenceNo: http.StatusConflict,
		PartnerNotFound:             http.StatusNotFound,
		CouponNotFound:              http.StatusNotFound,
	}
)
