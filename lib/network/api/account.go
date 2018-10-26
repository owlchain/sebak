package api

import (
	"fmt"
	"net/http"
	"strings"

	"boscoin.io/sebak/lib/block"
	"boscoin.io/sebak/lib/common"
	"boscoin.io/sebak/lib/common/observer"
	"boscoin.io/sebak/lib/error"
	"boscoin.io/sebak/lib/network/api/resource"
	"boscoin.io/sebak/lib/network/httputils"
	"boscoin.io/sebak/lib/storage"
	"boscoin.io/sebak/lib/transaction/operation"
	"github.com/gorilla/mux"
)

func (api NetworkHandlerAPI) GetAccountHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := vars["id"]

	readFunc := func() (payload interface{}, err error) {
		found, err := block.ExistsBlockAccount(api.storage, address)
		if err != nil {
			return nil, err
		}
		if !found {
			return nil, errors.ErrorBlockAccountDoesNotExists
		}
		ba, err := block.GetBlockAccount(api.storage, address)
		if err != nil {
			return nil, err
		}
		payload = resource.NewAccount(ba)
		return payload, nil
	}

	if httputils.IsEventStream(r) {
		event := fmt.Sprintf("address-%s", address)
		es := NewEventStream(w, r, renderEventStream, DefaultContentType)
		payload, err := readFunc()
		if err == nil {
			es.Render(payload)
		}
		es.Run(observer.BlockAccountObserver, event)
		return
	}

	payload, err := readFunc()
	if err != nil {
		httputils.WriteJSONError(w, err)
		return
	}

	if err := httputils.WriteJSON(w, 200, payload); err != nil {
		httputils.WriteJSONError(w, err)
	}
}

func (api NetworkHandlerAPI) GetFrozenAccountsByAccountHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := vars["id"]
	options, err := storage.NewDefaultListOptionsFromQuery(r.URL.Query())
	if err != nil {
		http.Error(w, errors.ErrorInvalidQueryString.Error(), http.StatusBadRequest)
		return
	}
	//var s string
	var cursor []byte
	readFunc := func() []resource.Resource {
		var txs []resource.Resource
		iterFunc, closeFunc := block.GetBlockAccountsByLinked(api.storage, address, options)
		for {
			var (
				create_block_height     uint64
				create_opHash           string
				sequenceid              uint64
				amount                  common.Amount
				state                   resource.FrozenAccountState
				unfreezing_block_height uint64
				unfreezing_opHash       string
				payment_opHash          string
			)

			t, hasNext, c := iterFunc()

			if !hasNext {
				break
			}
			cursor = make([]byte, len(c))
			copy(cursor, c)
			var bo block.BlockOperation
			if bo, err = block.GetBlockOperationCreateFrozen(api.storage, t.Address); err != nil {
				break
			} else {
				var (
					casted operation.CreateAccount
					ok     bool
					body   operation.Body
				)
				if body, err = operation.UnmarshalBodyJSON(bo.Type, bo.Body); err != nil {
					break
				}
				if casted, ok = body.(operation.CreateAccount); !ok {
					break
				}
				create_block_height = bo.BlockHeight
				create_opHash = bo.OpHash
				var tx block.BlockTransaction
				if tx, err = block.GetBlockTransaction(api.storage, bo.TxHash); err != nil {
					break
				}
				sequenceid = tx.SequenceID
				amount = casted.Amount
			}

			opIterFunc, opCloseFunc := block.GetBlockOperationsBySource(api.storage, t.Address, nil)
			state = resource.FrozenState
			for {
				bo, hasNext, _ := opIterFunc()
				switch bo.Type {
				case operation.TypePayment:
					state = resource.RefundState
					payment_opHash = bo.OpHash
				case operation.TypeUnfreezingRequest:
					lastblock := block.GetLatestBlock(api.storage)
					if lastblock.Height-bo.BlockHeight >= common.UnfreezingPeriod {
						state = resource.UnfreezedState
					} else {
						state = resource.MeltingState
					}
					unfreezing_opHash = bo.OpHash
					unfreezing_block_height = bo.BlockHeight
				}
				if !hasNext {
					break
				}
			}
			opCloseFunc()
			frozenAccountResource := resource.NewFrozenAccount(
				t,
				create_block_height,
				create_opHash,
				sequenceid,
				amount,
				state,
				unfreezing_block_height,
				unfreezing_opHash,
				payment_opHash,
			)
			txs = append(txs, frozenAccountResource)
		}
		closeFunc()
		return txs
	}
	// TODO
	if httputils.IsEventStream(r) {
		event := fmt.Sprintf("source-%s", address)
		es := NewEventStream(w, r, renderEventStream, DefaultContentType)
		txs := readFunc()
		for _, tx := range txs {
			es.Render(tx)
		}
		es.Run(observer.BlockOperationObserver, event)
		return
	}

	txs := readFunc() //TODO paging support
	self := r.URL.String()
	next := strings.Replace(resource.URLAccountFrozenAccounts, "{id}", address, -1) + "?" + options.SetCursor(cursor).SetReverse(false).Encode()
	prev := strings.Replace(resource.URLAccountFrozenAccounts, "{id}", address, -1) + "?" + options.SetReverse(true).Encode()
	list := resource.NewResourceList(txs, self, next, prev)

	if err := httputils.WriteJSON(w, 200, list); err != nil {
		httputils.WriteJSONError(w, err)
		return
	}
}

func (api NetworkHandlerAPI) GetFrozenAccountsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := vars["id"]
	options, err := storage.NewDefaultListOptionsFromQuery(r.URL.Query())
	if err != nil {
		http.Error(w, errors.ErrorInvalidQueryString.Error(), http.StatusBadRequest)
		return
	}
	//var s string
	var cursor []byte
	readFunc := func() []resource.Resource {
		var txs []resource.Resource
		iterFunc, closeFunc := block.GetBlockAccountsByFrozen(api.storage, options)
		for {
			var (
				create_block_height     uint64
				create_opHash           string
				sequenceid              uint64
				amount                  common.Amount
				state                   resource.FrozenAccountState
				unfreezing_block_height uint64
				unfreezing_opHash       string
				payment_opHash          string
			)

			t, hasNext, c := iterFunc()

			if !hasNext {
				break
			}
			cursor = make([]byte, len(c))
			copy(cursor, c)
			var bo block.BlockOperation
			if bo, err = block.GetBlockOperationCreateFrozen(api.storage, t.Address); err != nil {
				break
			} else {
				var (
					casted operation.CreateAccount
					ok     bool
					body   operation.Body
				)
				if body, err = operation.UnmarshalBodyJSON(bo.Type, bo.Body); err != nil {
					break
				}
				if casted, ok = body.(operation.CreateAccount); !ok {
					break
				}
				create_block_height = bo.BlockHeight
				create_opHash = bo.OpHash
				var tx block.BlockTransaction
				if tx, err = block.GetBlockTransaction(api.storage, bo.TxHash); err != nil {
					break
				}
				sequenceid = tx.SequenceID
				amount = casted.Amount
			}

			opIterFunc, opCloseFunc := block.GetBlockOperationsBySource(api.storage, t.Address, nil)
			state = resource.FrozenState
			for {
				bo, hasNext, _ := opIterFunc()
				switch bo.Type {
				case operation.TypePayment:
					state = resource.RefundState
					payment_opHash = bo.OpHash
				case operation.TypeUnfreezingRequest:
					lastblock := block.GetLatestBlock(api.storage)
					if lastblock.Height-bo.BlockHeight >= common.UnfreezingPeriod {
						state = resource.UnfreezedState
					} else {
						state = resource.MeltingState
					}
					unfreezing_opHash = bo.OpHash
					unfreezing_block_height = bo.BlockHeight
				}
				if !hasNext {
					break
				}
			}
			opCloseFunc()
			frozenAccountResource := resource.NewFrozenAccount(
				t,
				create_block_height,
				create_opHash,
				sequenceid,
				amount,
				state,
				unfreezing_block_height,
				unfreezing_opHash,
				payment_opHash,
			)
			txs = append(txs, frozenAccountResource)
		}
		closeFunc()
		return txs
	}
	// TODO
	if httputils.IsEventStream(r) {
		event := fmt.Sprintf("source-%s", address)
		es := NewEventStream(w, r, renderEventStream, DefaultContentType)
		txs := readFunc()
		for _, tx := range txs {
			es.Render(tx)
		}
		es.Run(observer.BlockOperationObserver, event)
		return
	}

	txs := readFunc() //TODO paging support
	self := r.URL.String()
	next := strings.Replace(resource.URLAccountFrozenAccounts, "{id}", address, -1) + "?" + options.SetCursor(cursor).SetReverse(false).Encode()
	prev := strings.Replace(resource.URLAccountFrozenAccounts, "{id}", address, -1) + "?" + options.SetReverse(true).Encode()
	list := resource.NewResourceList(txs, self, next, prev)

	if err := httputils.WriteJSON(w, 200, list); err != nil {
		httputils.WriteJSONError(w, err)
		return
	}
}
