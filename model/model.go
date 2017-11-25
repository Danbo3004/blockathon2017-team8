package model

type Company struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	NumberOfToken int64  `json:"number_of_token"`
	TokenName     string `json:"token_name"`
	Owner         string `json:"owner"`
	IssuerAddress string `json:"issuer_address"`
}

type User struct {
	ID      string `json:"id"`
	Address string `json:"address"`
}

type UserCompany struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	CompanyID string `json:"company_id"`
	TokenName string `json:"token_name"`
}
