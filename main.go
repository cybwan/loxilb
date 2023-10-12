package main

import (
	"encoding/json"
	"fmt"

	tk "github.com/loxilb-io/loxilib"

	"github.com/cybwan/loxilb/pkg/cmn"
	"github.com/cybwan/loxilb/pkg/nlp"
)

type NetHook struct {
}

func (n NetHook) NetMirrorGet() ([]cmn.MirrGetMod, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetHook) NetMirrorAdd(mod *cmn.MirrMod) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetHook) NetMirrorDel(mod *cmn.MirrMod) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetHook) NetPortGet() ([]cmn.PortDump, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetHook) NetPortAdd(mod *cmn.PortMod) (int, error) {
	bytes, _ := json.Marshal(mod)
	fmt.Println("NetPortAdd: ", string(bytes))
	return 0, nil
}

func (n NetHook) NetPortDel(mod *cmn.PortMod) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetHook) NetVlanGet() ([]cmn.VlanGet, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetHook) NetVlanAdd(mod *cmn.VlanMod) (int, error) {
	bytes, _ := json.Marshal(mod)
	fmt.Println("NetVlanAdd: ", string(bytes))
	return 0, nil
}

func (n NetHook) NetVlanDel(mod *cmn.VlanMod) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetHook) NetVlanPortAdd(mod *cmn.VlanPortMod) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetHook) NetVlanPortDel(mod *cmn.VlanPortMod) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetHook) NetFdbAdd(mod *cmn.FdbMod) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetHook) NetFdbDel(mod *cmn.FdbMod) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetHook) NetAddrGet() ([]cmn.IPAddrGet, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetHook) NetAddrAdd(mod *cmn.IPAddrMod) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetHook) NetAddrDel(mod *cmn.IPAddrMod) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetHook) NetNeighGet() ([]cmn.NeighMod, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetHook) NetNeighAdd(mod *cmn.NeighMod) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetHook) NetNeighDel(mod *cmn.NeighMod) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetHook) NetRouteGet() ([]cmn.RouteGet, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetHook) NetRouteAdd(mod *cmn.RouteMod) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetHook) NetRouteDel(mod *cmn.RouteMod) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetHook) NetLbRuleAdd(mod *cmn.LbRuleMod) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetHook) NetLbRuleDel(mod *cmn.LbRuleMod) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetHook) NetLbRuleGet() ([]cmn.LbRuleMod, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetHook) NetCtInfoGet() ([]cmn.CtInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetHook) NetSessionGet() ([]cmn.SessionMod, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetHook) NetSessionUlClGet() ([]cmn.SessionUlClMod, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetHook) NetSessionAdd(mod *cmn.SessionMod) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetHook) NetSessionDel(mod *cmn.SessionMod) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetHook) NetSessionUlClAdd(mod *cmn.SessionUlClMod) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetHook) NetSessionUlClDel(mod *cmn.SessionUlClMod) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetHook) NetPolicerGet() ([]cmn.PolMod, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetHook) NetPolicerAdd(mod *cmn.PolMod) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetHook) NetPolicerDel(mod *cmn.PolMod) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetHook) NetCIStateMod(mod *cmn.HASMod) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetHook) NetCIStateGet() ([]cmn.HASMod, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetHook) NetFwRuleAdd(mod *cmn.FwRuleMod) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetHook) NetFwRuleDel(mod *cmn.FwRuleMod) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetHook) NetFwRuleGet() ([]cmn.FwRuleMod, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetHook) NetEpHostAdd(fm *cmn.EndPointMod) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetHook) NetEpHostDel(fm *cmn.EndPointMod) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetHook) NetEpHostGet() ([]cmn.EndPointMod, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetHook) NetParamSet(param cmn.ParamMod) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetHook) NetParamGet(param *cmn.ParamMod) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetHook) NetGoBGPNeighAdd(nm *cmn.GoBGPNeighMod) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetHook) NetGoBGPNeighDel(nm *cmn.GoBGPNeighMod) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetHook) NetGoBGPGCAdd(gc *cmn.GoBGPGlobalConfig) (int, error) {
	//TODO implement me
	panic("implement me")
}

func main() {
	tk.LogItInit("/tmp/netlink.log", tk.LogDebug, true)
	hook := new(NetHook)
	nlp.NlpRegister(hook)
	nlp.NlpInit()
	wait := make(chan int)
	<-wait
}
