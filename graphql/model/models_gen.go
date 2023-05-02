// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Commitment interface {
	IsCommitment()
	GetID() string
	GetName() string
	GetDeadline() string
	GetStatus() string
}

type Bill struct {
	ID       string  `_id,omitempty`
	Name     string  `json:"name"`
	Deadline string  `json:"deadline"`
	Status   string  `json:"status"`
	Amount   float64 `json:"amount"`
}

func (Bill) IsCommitment()            {}
func (this Bill) GetID() string       { return this.ID }
func (this Bill) GetName() string     { return this.Name }
func (this Bill) GetDeadline() string { return this.Deadline }
func (this Bill) GetStatus() string   { return this.Status }

type CreateBillInput struct {
	Name     string  `json:"name"`
	Deadline string  `json:"deadline"`
	Status   string  `json:"status"`
	Amount   float64 `json:"amount"`
}

type DeleteBillResponse struct {
	DeletedBillID string `json:"deletedBillId"`
}

type UpdateBillInput struct {
	Name     *string  `json:"name"`
	Deadline *string  `json:"deadline"`
	Status   *string  `json:"status"`
	Amount   *float64 `json:"amount"`
}
