package loxinet

import "C"
import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"
	"unsafe"

	"github.com/gorilla/mux"
	nlp "github.com/vishvananda/netlink"

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
	case *PortDpWorkQ:
		w := m.(*PortDpWorkQ)
		if w.Work == DpCreate {
			if w.LoadEbpf != "" && w.LoadEbpf != "lo" && w.LoadEbpf != "llb0" && w.LoadEbpf != "flb0" {
				loadEbpfPgm(w.LoadEbpf)
			}
		}
		netlinkQMeta.PortDpWorkQ = append(netlinkQMeta.PortDpWorkQ, m.(*PortDpWorkQ))
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

// loadEbpfPgm - load eBPF program to an interface
func loadEbpfPgm(name string) int {
	ifStr := C.CString(name)
	xSection := C.CString(string(C.XDP_LL_SEC_DEFAULT))
	link, err := nlp.LinkByName(name)
	if err != nil {
		tk.LogIt(tk.LogWarning, "[DP] Port %s not found\n", name)
		return -1
	}
	//if e.RssEn {
	//	C.llb_dp_link_attach(ifStr, xSection, C.LL_BPF_MOUNT_XDP, 0)
	//}
	section := C.CString(string(C.TC_LL_SEC_DEFAULT))
	ret := C.llb_dp_link_attach(ifStr, section, C.LL_BPF_MOUNT_TC, 0)

	filters, err := nlp.FilterList(link, nlp.HANDLE_MIN_INGRESS)
	if err != nil {
		tk.LogIt(tk.LogWarning, "[DP] Filter on %s not found\n", name)
		return -1
	}
	ret = -1
	for _, f := range filters {
		if t, ok := f.(*nlp.BpfFilter); ok {
			if strings.Contains(t.Name, C.TC_LL_SEC_DEFAULT) {
				ret = 0
				break
			}
		}
	}
	C.free(unsafe.Pointer(ifStr))
	C.free(unsafe.Pointer(xSection))
	C.free(unsafe.Pointer(section))
	return int(ret)
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
