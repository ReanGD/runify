package config

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
type SConfiguration struct {
	DsX11   *DsX11Cfg
	Desktop *DesktopCfg
}

func newSConfiguration() *SConfiguration {
	return &SConfiguration{
		DsX11:   newDsX11Cfg(),
		Desktop: newDesktopCfg(),
	}
}
