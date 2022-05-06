package entities

type SendEmailServiceRequest struct {
	EmailDestination string
	Subject          string
	ContentBody      string
}
