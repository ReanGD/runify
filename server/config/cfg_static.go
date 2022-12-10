package config

type RpcCfg struct {
	ChannelLen   uint32
	SendMsgChLen uint32
}

func newRpcCfg() *RpcCfg {
	return &RpcCfg{
		ChannelLen:   100,
		SendMsgChLen: 500,
	}
}

type ProviderCfg struct {
	ChannelLen          uint32
	SubModuleChannelLen uint32
}

func newProviderCfg() *ProviderCfg {
	return &ProviderCfg{
		ChannelLen:          100,
		SubModuleChannelLen: 100,
	}
}

type DsX11Cfg struct {
	ModuleChLen   uint32
	X11EventChLen uint32
}

func newDsX11Cfg() *DsX11Cfg {
	return &DsX11Cfg{
		ModuleChLen:   100,
		X11EventChLen: 100,
	}
}

type DesktopCfg struct {
	ModuleChLen                uint32
	PrimarySubscriptionChLen   uint32
	ClipboardSubscriptionChLen uint32
	HotkeySubscriptionChLen    uint32
}

func newDesktopCfg() *DesktopCfg {
	return &DesktopCfg{
		ModuleChLen:                100,
		PrimarySubscriptionChLen:   100,
		ClipboardSubscriptionChLen: 100,
		HotkeySubscriptionChLen:    100,
	}
}

// Static configuration
type CfgStatic struct {
	Rpc      *RpcCfg
	Provider *ProviderCfg
	DsX11    *DsX11Cfg
	Desktop  *DesktopCfg
}

func newCfgStatic() *CfgStatic {
	return &CfgStatic{
		Rpc:      newRpcCfg(),
		Provider: newProviderCfg(),
		DsX11:    newDsX11Cfg(),
		Desktop:  newDesktopCfg(),
	}
}
