#include "runify_native_plugin.h"
#include "runify_native_window.h"


#define RUNIFY_NATIVE_PLUGIN(obj) (G_TYPE_CHECK_INSTANCE_CAST((obj), runify_native_plugin_get_type(), RunifyNativePlugin))

struct _RunifyNativePlugin {
  GObject parent_instance;
  RNWindow* rnw = nullptr;
};

G_DEFINE_TYPE(RunifyNativePlugin, runify_native_plugin, g_object_get_type())

static void runify_native_plugin_dispose(GObject* object) {
  RunifyNativePlugin* plugin = RUNIFY_NATIVE_PLUGIN(object);
  if (plugin->rnw != nullptr) {
    delete plugin->rnw;
    plugin->rnw = nullptr;
  }

  G_OBJECT_CLASS(runify_native_plugin_parent_class)->dispose(object);
}

static void runify_native_plugin_class_init(RunifyNativePluginClass* klass) {
  G_OBJECT_CLASS(klass)->dispose = runify_native_plugin_dispose;
}

static void runify_native_plugin_init(RunifyNativePlugin* self) {}

static void method_call_cb(FlMethodChannel* channel, FlMethodCall* method_call, gpointer user_data) {
  RunifyNativePlugin* plugin = RUNIFY_NATIVE_PLUGIN(user_data);
  plugin->rnw->HandleMethodCall(method_call);
}

void runify_native_plugin_pre_init(GtkWindow* window) {
  gtk_window_set_modal(window, TRUE);
  gtk_window_set_keep_above(window, TRUE); // window stays on top
  gtk_window_set_decorated(window, FALSE);
  gtk_window_set_skip_taskbar_hint(window, TRUE);
  gtk_window_resize(window, 1, 1);
}

void runify_native_plugin_register(FlPluginRegistry* registry) {
  g_autoptr(FlPluginRegistrar) registrar = fl_plugin_registry_get_registrar_for_plugin(registry, "RunifyNativePlugin");
  RunifyNativePlugin* plugin = RUNIFY_NATIVE_PLUGIN(g_object_new(runify_native_plugin_get_type(), nullptr));
  g_autoptr(FlStandardMethodCodec) codec = fl_standard_method_codec_new();
  g_autoptr(FlMethodChannel) channel = fl_method_channel_new(
    fl_plugin_registrar_get_messenger(registrar), "runify_native", FL_METHOD_CODEC(codec));

  FlView* view = fl_plugin_registrar_get_view(registrar);
  GtkWindow* gtk_window = GTK_WINDOW(gtk_widget_get_toplevel(GTK_WIDGET(view)));
  plugin->rnw = new RNWindow(gtk_window, channel);

  fl_method_channel_set_method_call_handler(channel, method_call_cb, g_object_ref(plugin), g_object_unref);
  g_object_unref(plugin);
}
