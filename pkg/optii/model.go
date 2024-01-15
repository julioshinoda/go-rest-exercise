package optii

type Response struct {
	PageInfo PageInfo `json:"pageInfo"`
	Items    []Items  `json:"items"`
}
type PageInfo struct {
	TotalCount  int  `json:"totalCount"`
	EndCursor   int  `json:"endCursor"`
	HasNextPage bool `json:"hasNextPage"`
}
type Items struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type AuthResponse struct {
	RefreshTokenExpiresIn int      `json:"refresh_token_expires_in"`
	APIProductList        string   `json:"api_product_list"`
	APIProductListJSON    []string `json:"api_product_list_json"`
	OrganizationName      string   `json:"organization_name"`
	DeveloperEmail        string   `json:"developer.email"`
	TokenType             string   `json:"token_type"`
	IssuedAt              string   `json:"issued_at"`
	ClientID              string   `json:"client_id"`
	AccessToken           string   `json:"access_token"`
	ApplicationName       string   `json:"application_name"`
	Scope                 string   `json:"scope"`
	ExpiresIn             int      `json:"expires_in"`
	RefreshCount          string   `json:"refresh_count"`
	Status                string   `json:"status"`
}
