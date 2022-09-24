#ifndef FLUTTER_PLUGIN_RUNIFY_PLUGIN_H_
#define FLUTTER_PLUGIN_RUNIFY_PLUGIN_H_

#include <flutter_linux/flutter_linux.h>

G_BEGIN_DECLS

#ifdef FLUTTER_PLUGIN_IMPL
#define FLUTTER_PLUGIN_EXPORT __attribute__((visibility("default")))
#else
#define FLUTTER_PLUGIN_EXPORT
#endif

typedef struct _RunifyNativePlugin RunifyNativePlugin;
typedef struct {
  GObjectClass parent_class;
} RunifyNativePluginClass;

FLUTTER_PLUGIN_EXPORT GType runify_native_plugin_get_type();

FLUTTER_PLUGIN_EXPORT void runify_native_plugin_pre_init(GtkWindow* window);
FLUTTER_PLUGIN_EXPORT void runify_native_plugin_register(FlPluginRegistry* registry);

G_END_DECLS

#endif  // FLUTTER_PLUGIN_RUNIFY_PLUGIN_H_
