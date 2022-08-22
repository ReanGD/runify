package gtk

// #cgo pkg-config: gtk+-3.0
// #include <gtk/gtk.h>
// #include <glib-object.h>
// #include "pixbuf.go.h"
import "C"
import (
	"errors"
	"runtime"
	"strconv"
	"unsafe"
)

func Init() {
	C.gtk_init(nil, nil)
}

type object struct {
	GObject *C.GObject
}

// newObject creates a new Object from a GObject pointer.
func newObject(p *C.GObject) *object {
	if p == nil {
		return nil
	}
	return &object{GObject: p}
}

// ToGObject type converts an unsafe.Pointer as a native C GObject.
// This function is exported for visibility in other gotk3 packages and
// is not meant to be used by applications.
func toGObject(p unsafe.Pointer) *C.GObject {
	return (*C.GObject)(p)
}

func takeObject(ptr unsafe.Pointer) *object {
	obj := newObject(toGObject(ptr))
	if obj == nil {
		return nil
	}

	obj.refSink()
	runtime.SetFinalizer(obj, (*object).unref)
	return obj
}

func (o *object) refSink() {
	C.g_object_ref_sink(C.gpointer(o.GObject))
}

// Unref is a wrapper around g_object_unref().
func (v *object) unref() {
	C.g_object_unref(C.gpointer(v.GObject))
}

type Pixbuf struct {
	*object
}

// native returns a pointer to the underlying GdkPixbuf.
func (v *Pixbuf) native() *C.GdkPixbuf {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGdkPixbuf(p)
}

func (v *Pixbuf) SavePNG(path string, compression int) error {
	cpath := C.CString(path)
	ccompression := C.CString(strconv.Itoa(compression))
	defer C.free(unsafe.Pointer(cpath))
	defer C.free(unsafe.Pointer(ccompression))

	var err *C.GError
	c := C._gdk_pixbuf_save_png(v.native(), cpath, &err, ccompression)
	if c == C.FALSE {
		defer C.g_error_free(err)
		return errors.New(C.GoString((*C.char)(err.message)))
	}
	return nil
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

func (v *IconTheme) LoadIcon(iconName string, size int, flags IconLookupFlags) (*Pixbuf, error) {
	cstr := C.CString(iconName)
	defer C.free(unsafe.Pointer(cstr))
	var err *C.GError = nil
	c := C.gtk_icon_theme_load_icon(v.Theme, (*C.gchar)(cstr), C.gint(size), C.GtkIconLookupFlags(flags), &err)
	if c == nil {
		defer C.g_error_free(err)
		return nil, errors.New(goString(err.message))
	}
	return &Pixbuf{takeObject(unsafe.Pointer(c))}, nil
}
