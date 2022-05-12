package entities

// SendEmailServiceRequest request for sending an email
type SendEmailServiceRequest struct {
	EmailDestination string
	Subject          string
	ContentBody      string
}
