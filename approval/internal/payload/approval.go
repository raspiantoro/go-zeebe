package payload

type ApprovalRequest struct {
	PurchaseID string `json:"purchase_id"`
	Action     string `json:"action"`
	Username   string `json:"username"`
}
