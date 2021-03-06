package api

import (
	"github.com/nvellon/hal"
	"strings"
)

type APIResource interface {
	LinkSelf() string
	Resource() *hal.Resource
	GetMap() hal.Entry
	Serialize() ([]byte, error)
}

type APIResourceList struct {
	Resources []APIResource
	SelfLink  string
}

func (al APIResourceList) Resource() *hal.Resource {
	rl := hal.NewResource(struct{}{}, al.LinkSelf())
	for _, apiResource := range al.Resources {
		r := apiResource.Resource()
		rl.Embed("records", r)
	}
	rl.AddLink("prev", hal.NewLink(al.LinkSelf())) //TODO: set prev/next url
	rl.AddLink("next", hal.NewLink(al.LinkSelf()))

	return rl
}

func (al APIResourceList) Serialize() (encoded []byte, err error) {
	return al.Resource().MarshalJSON()
}

func (al APIResourceList) LinkSelf() string {
	return al.SelfLink
}
func (al APIResourceList) GetMap() hal.Entry {
	return hal.Entry{}
}

const (
	UrlAccounts     = "/accounts/{id}"
	UrlTransactions = "/transactions/{id}"
	UrlOperations   = "/operations/{id}"
)

type APIResourceAccount struct {
	accountId  string
	sequenceID uint64
	balance    string
}

func (aa APIResourceAccount) GetMap() hal.Entry {
	return hal.Entry{
		"id":          aa.accountId,
		"account_id":  aa.accountId,
		"sequence_id": aa.sequenceID,
		"balance":     aa.balance,
	}
}

func (aa APIResourceAccount) Resource() *hal.Resource {
	r := hal.NewResource(aa, aa.LinkSelf())
	r.AddLink("transactions", hal.NewLink(strings.Replace(UrlAccounts, "{id}", aa.accountId, -1)+"/transactions{?cursor,limit,order}", hal.LinkAttr{"templated": true}))
	r.AddLink("operations", hal.NewLink(strings.Replace(UrlAccounts, "{id}", aa.accountId, -1)+"/operations{?cursor,limit,order}", hal.LinkAttr{"templated": true}))
	return r
}

func (aa APIResourceAccount) LinkSelf() string {
	return strings.Replace(UrlAccounts, "{id}", aa.accountId, -1)
}

func (aa APIResourceAccount) Serialize() (encoded []byte, err error) {
	return aa.Resource().MarshalJSON()
}

type APIResourceTransaction struct {
	hash       string
	sequenceID uint64
	signature  string
	source     string
	fee        string
	amount     string
	created    string
	operations []string
}

func (at APIResourceTransaction) GetMap() hal.Entry {
	return hal.Entry{
		"id":              at.hash,
		"hash":            at.hash,
		"account":         at.source,
		"fee_paid":        at.fee,
		"sequence_id":     at.sequenceID,
		"created_at":      at.created,
		"operation_count": len(at.operations),
	}
}
func (at APIResourceTransaction) Resource() *hal.Resource {

	r := hal.NewResource(at, at.LinkSelf())
	r.AddLink("accounts", hal.NewLink(strings.Replace(UrlAccounts, "{id}", at.source, -1)))
	r.AddLink("operations", hal.NewLink(strings.Replace(UrlTransactions, "{id}", at.hash, -1)+"/operations{?cursor,limit,order}", hal.LinkAttr{"templated": true}))
	return r
}

func (at APIResourceTransaction) LinkSelf() string {
	return strings.Replace(UrlTransactions, "{id}", at.hash, -1)
}

func (at APIResourceTransaction) Serialize() (encoded []byte, err error) {
	return at.Resource().MarshalJSON()
}

type APIResourceOperation struct {
	hash    string
	txHash  string
	funder  string //Source Account
	account string //Target Account
	otype   string
	amount  string
}

func (ao APIResourceOperation) GetMap() hal.Entry {
	return hal.Entry{
		"id":      ao.hash,
		"hash":    ao.hash,
		"funder":  ao.funder,
		"account": ao.account,
		"type":    ao.otype,
		"amount":  ao.amount,
	}
}

func (ao APIResourceOperation) Resource() *hal.Resource {

	r := hal.NewResource(ao, ao.LinkSelf())
	r.AddNewLink("transactions", strings.Replace(UrlTransactions, "{id}", ao.txHash, -1))
	return r
}

func (ao APIResourceOperation) LinkSelf() string {
	return strings.Replace(UrlOperations, "{id}", ao.hash, -1)
}

func (ao APIResourceOperation) Serialize() (encoded []byte, err error) {
	return ao.Resource().MarshalJSON()
}
