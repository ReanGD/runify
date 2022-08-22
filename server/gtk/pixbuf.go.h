// Same copyright and license as the rest of the files in this project

#include <stdlib.h>

static GdkPixbuf *toGdkPixbuf(void *p) { return (GDK_PIXBUF(p)); }

gboolean _gdk_pixbuf_save_png(GdkPixbuf *pixbuf, const char *filename, GError **err, const char *compression) {
  return gdk_pixbuf_save(pixbuf, filename, "png", err, "compression", compression, NULL);
}
