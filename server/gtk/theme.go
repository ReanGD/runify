package gtk

// #cgo pkg-config: gtk+-3.0
// #include <gtk/gtk.h>
import "C"
import (
	"errors"
	"unsafe"
)

func Init() {
	C.gtk_init(nil, nil)
}

type IconLookupFlags int

const (
	ICON_LOOKUP_NO_SVG           IconLookupFlags = C.GTK_ICON_LOOKUP_NO_SVG
	ICON_LOOKUP_FORCE_SVG        IconLookupFlags = C.GTK_ICON_LOOKUP_FORCE_SVG
	ICON_LOOKUP_USE_BUILTIN      IconLookupFlags = C.GTK_ICON_LOOKUP_USE_BUILTIN
	ICON_LOOKUP_GENERIC_FALLBACK IconLookupFlags = C.GTK_ICON_LOOKUP_GENERIC_FALLBACK
	ICON_LOOKUP_FORCE_SIZE       IconLookupFlags = C.GTK_ICON_LOOKUP_FORCE_SIZE
)

type IconTheme struct {
	Theme *C.GtkIconTheme
}

var nilPtrErr = errors.New("cgo returned unexpected nil pointer")

func IconThemeGetDefault() (*IconTheme, error) {
	c := C.gtk_icon_theme_get_default()
	if c == nil {
		return nil, nilPtrErr
	}
	return &IconTheme{c}, nil
}

func goString(cstr *C.gchar) string {
	return C.GoString((*C.char)(cstr))
}

func (v *IconTheme) LookupIcon(iconName string, size int, flags IconLookupFlags) string {
	cstr := C.CString(iconName)
	defer C.free(unsafe.Pointer(cstr))
	iconInfo := C.gtk_icon_theme_lookup_icon(v.Theme, (*C.gchar)(cstr), C.gint(size), C.GtkIconLookupFlags(flags))
	if iconInfo == nil {
		return ""
	}
	defer C.g_object_unref(C.gpointer(iconInfo))

	cFilename := C.gtk_icon_info_get_filename(iconInfo)
	if cFilename == nil {
		return ""
	}

	return goString(cFilename)
}
