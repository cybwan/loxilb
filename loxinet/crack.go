package loxinet

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/gorilla/mux"

	tk "github.com/loxilb-io/loxilib"
)

type Meta struct {
	PortDpWorkQ      []*PortDpWorkQ
	L2AddrDpWorkQ    []*L2AddrDpWorkQ
	RouteDpWorkQ     []*RouteDpWorkQ
	RouterMacDpWorkQ []*RouterMacDpWorkQ
	NextHopDpWorkQ   []*NextHopDpWorkQ

	MirrDpWorkQ  []*MirrDpWorkQ
	PolDpWorkQ   []*PolDpWorkQ
	NatDpWorkQ   []*NatDpWorkQ
	UlClDpWorkQ  []*UlClDpWorkQ
	StatDpWorkQ  []*StatDpWorkQ
	TableDpWorkQ []*TableDpWorkQ
	FwDpWorkQ    []*FwDpWorkQ
	PeerDpWorkQ  []*PeerDpWorkQ
}

var (
	netlinkQMeta = new(Meta)
)

func DpWorkDemo(dp *DpH, m interface{}) (*DpH, DpRetT) {
	mtype := reflect.TypeOf(m)
	mbytes, _ := json.Marshal(m)
	fmt.Println(mtype, string(mbytes))
	var ret DpRetT
	switch m.(type) {
	//case *MirrDpWorkQ:
	//	netlinkQMeta.MirrDpWorkQ = append(netlinkQMeta.MirrDpWorkQ, m.(*MirrDpWorkQ))
	//case *PolDpWorkQ:
	//	netlinkQMeta.PolDpWorkQ = append(netlinkQMeta.PolDpWorkQ, m.(*PolDpWorkQ))
	// case *PortDpWorkQ:
	// 	netlinkQMeta.PortDpWorkQ = append(netlinkQMeta.PortDpWorkQ, m.(*PortDpWorkQ))
	//case *L2AddrDpWorkQ:
	//	netlinkQMeta.L2AddrDpWorkQ = append(netlinkQMeta.L2AddrDpWorkQ, m.(*L2AddrDpWorkQ))
	//case *RouterMacDpWorkQ:
	//	netlinkQMeta.RouterMacDpWorkQ = append(netlinkQMeta.RouterMacDpWorkQ, m.(*RouterMacDpWorkQ))
	//case *NextHopDpWorkQ:
	//	netlinkQMeta.NextHopDpWorkQ = append(netlinkQMeta.NextHopDpWorkQ, m.(*NextHopDpWorkQ))
	//case *RouteDpWorkQ:
	//	netlinkQMeta.RouteDpWorkQ = append(netlinkQMeta.RouteDpWorkQ, m.(*RouteDpWorkQ))
	//case *NatDpWorkQ:
	//	netlinkQMeta.NatDpWorkQ = append(netlinkQMeta.NatDpWorkQ, m.(*NatDpWorkQ))
	//case *UlClDpWorkQ:
	//	netlinkQMeta.UlClDpWorkQ = append(netlinkQMeta.UlClDpWorkQ, m.(*UlClDpWorkQ))
	//case *StatDpWorkQ:
	//	netlinkQMeta.StatDpWorkQ = append(netlinkQMeta.StatDpWorkQ, m.(*StatDpWorkQ))
	//case *TableDpWorkQ:
	//	netlinkQMeta.TableDpWorkQ = append(netlinkQMeta.TableDpWorkQ, m.(*TableDpWorkQ))
	//case *FwDpWorkQ:
	//	netlinkQMeta.FwDpWorkQ = append(netlinkQMeta.FwDpWorkQ, m.(*FwDpWorkQ))
	//case *PeerDpWorkQ:
	//	netlinkQMeta.PeerDpWorkQ = append(netlinkQMeta.PeerDpWorkQ, m.(*PeerDpWorkQ))
	default:
		ret = DpWqUnkErr
		return dp, ret
	}
	return nil, ret
}

func NetMetaHttpServer() {
	r := mux.NewRouter()
	handle(r, "")
	handle(r, "PortDpWorkQ")
	handle(r, "L2AddrDpWorkQ")
	handle(r, "RouteDpWorkQ")
	handle(r, "RouterMacDpWorkQ")
	handle(r, "NextHopDpWorkQ")
	handle(r, "MirrDpWorkQ")
	handle(r, "PolDpWorkQ")
	handle(r, "NatDpWorkQ")
	handle(r, "UlClDpWorkQ")
	handle(r, "StatDpWorkQ")
	handle(r, "TableDpWorkQ")
	handle(r, "FwDpWorkQ")
	handle(r, "PeerDpWorkQ")

	srv := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:80",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		tk.LogIt(tk.LogError, "unexpected ListenAndServe %v\n", err)
	}
}

func handle(r *mux.Router, path string) *mux.Route {
	if len(path) == 0 {
		return r.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
			response.Header().Set("Content-Type", "application/json")
			encoder := json.NewEncoder(response)
			encoder.SetIndent("", " ")
			if err := encoder.Encode(netlinkQMeta); err != nil {
				tk.LogIt(tk.LogError, "unexpected Encode %v\n", err)
			}
		}).Methods("GET")
	}
	return r.HandleFunc(fmt.Sprintf("/%s", path), func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Content-Type", "application/json")
		encoder := json.NewEncoder(response)
		encoder.SetIndent("", " ")
		s := reflect.ValueOf(netlinkQMeta).Elem()
		f := s.FieldByName(path)
		v := f.Interface()
		if err := encoder.Encode(v); err != nil {
			tk.LogIt(tk.LogError, "unexpected Encode %v\n", err)
		}
	}).Methods("GET")
}
