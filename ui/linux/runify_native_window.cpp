#include "runify_native_window.h"


struct Geometry {
  float x;
  float y;
  float width;
  float height;

  Geometry() = default;
  Geometry(FlValue* args);
  FlMethodResponse* ToFlResponse();
};

Geometry::Geometry(FlValue* args)
  : x(fl_value_get_float(fl_value_lookup_string(args, "x")))
  , y(fl_value_get_float(fl_value_lookup_string(args, "y")))
  , width(fl_value_get_float(fl_value_lookup_string(args, "width")))
  , height(fl_value_get_float(fl_value_lookup_string(args, "height"))) {

}

FlMethodResponse* Geometry::ToFlResponse() {
  g_autoptr(FlValue) result = fl_value_new_map();
  fl_value_set_string_take(result, "x", fl_value_new_float(x));
  fl_value_set_string_take(result, "y", fl_value_new_float(y));
  fl_value_set_string_take(result, "width", fl_value_new_float(width));
  fl_value_set_string_take(result, "height", fl_value_new_float(height));

  return FL_METHOD_RESPONSE(fl_method_success_response_new(result));
}

static FlMethodResponse* flBool(bool value) {
  g_autoptr(FlValue) result = fl_value_new_bool(value);
  return FL_METHOD_RESPONSE(fl_method_success_response_new(result));
}

static FlMethodResponse* flNotImplemented() {
  return FL_METHOD_RESPONSE(fl_method_not_implemented_response_new());
}

static gboolean gSignalCallback(GtkWidget* widget, GdkEvent* event, gpointer data) {
  return RNWindow::HandleGtkSignal(event);
}

RNWindow* RNWindow::instance = nullptr;

RNWindow::RNWindow(GtkWindow* gtk_window, FlMethodChannel* channel)
  : m_gtk_window(gtk_window)
  , m_gdk_window(gtk_widget_get_window(GTK_WIDGET(gtk_window)))
  , m_channel(channel) {

  RNWindow::instance = this;

  GdkDisplay* display = gdk_display_get_default();
  GdkMonitor* monitor = gdk_display_get_monitor_at_window(display, m_gdk_window);
  GdkRectangle frame;
  gdk_monitor_get_geometry(monitor, &frame);

  m_monitor_hwidth_px = frame.width / 2;
  m_monitor_hheight_px = frame.height / 2;
  m_window_width_ppm = static_cast<float>(frame.width) / static_cast<float>(gdk_monitor_get_width_mm(monitor));
  m_window_height_ppm = static_cast<float>(frame.height) / static_cast<float>(gdk_monitor_get_height_mm(monitor));
}

RNWindow::~RNWindow() {
  RNWindow::instance = nullptr;
}

void RNWindow::InitPlugin() {
  Hide();

  auto callback = G_CALLBACK(gSignalCallback);
  m_delete_handler = g_signal_connect(m_gtk_window, "delete_event", callback, nullptr);
  m_focus_in_handler = g_signal_connect(m_gtk_window, "focus-in-event", callback, nullptr);
  m_focus_out_handler = g_signal_connect(m_gtk_window, "focus-out-event", callback, nullptr);
  m_configure_handler = g_signal_connect(m_gtk_window, "configure-event", callback, nullptr);
}

void RNWindow::ClosePlugin() {
  auto disconnect = [this](gulong id) {
    if ((id != 0) && g_signal_handler_is_connected(m_gtk_window, id)) {
      g_signal_handler_disconnect(m_gtk_window, id);
    }
  };

  disconnect(m_delete_handler);
  disconnect(m_focus_in_handler);
  disconnect(m_focus_out_handler);
  disconnect(m_configure_handler);

  m_delete_handler = 0;
  m_focus_in_handler = 0;
  m_focus_out_handler = 0;
  m_configure_handler = 0;

  RNWindow::instance = nullptr;
}

bool RNWindow::IsVisible() const {
  return gtk_widget_is_visible(GTK_WIDGET(m_gtk_window));
}

void RNWindow::Show() const {
  gtk_widget_show(GTK_WIDGET(m_gtk_window));
  gtk_window_present(m_gtk_window);
}

void RNWindow::Hide() const {
  gtk_widget_hide(GTK_WIDGET(m_gtk_window));
}

bool RNWindow::IsFocused() const {
  return gtk_window_is_active(m_gtk_window);
}

void RNWindow::Focus() const {
  gtk_window_present(m_gtk_window);
}

void RNWindow::SetOpacity(double opacity) const {
  gtk_widget_set_opacity(GTK_WIDGET(m_gtk_window), opacity);
}

void RNWindow::GetGeometry(Geometry& g) const {
  gint x_px, y_px, width_px, height_px;
  gtk_window_get_position(m_gtk_window, &x_px, &y_px);
  gtk_window_get_size(m_gtk_window, &width_px, &height_px);

  // Result is in millimeter
  g.x = static_cast<float>(x_px + width_px / 2 - m_monitor_hwidth_px) / m_window_width_ppm;
  g.y = static_cast<float>(y_px + height_px / 2 - m_monitor_hheight_px) / m_window_height_ppm;
  g.width = static_cast<float>(width_px) / m_window_width_ppm;
  g.height = static_cast<float>(height_px) / m_window_height_ppm;
}

void RNWindow::SetGeometry(const Geometry& g) const {
  gint width_px = static_cast<gint>(g.width * m_window_width_ppm + 0.5);
  gint height_px = static_cast<gint>(g.height * m_window_height_ppm + 0.5);
  gint x_px = static_cast<gint>((g.x - g.width * 0.5) * m_window_width_ppm + 0.5) + m_monitor_hwidth_px;
  gint y_px = static_cast<gint>((g.y - g.height * 0.5) * m_window_height_ppm + 0.5) + m_monitor_hheight_px;

  gdk_window_move_resize(m_gdk_window, x_px, y_px, width_px, height_px);
}

void RNWindow::SetGeometryHint(int min_width, int min_height) const {
  GdkGeometry geometry;
  geometry.min_width = min_width;
  geometry.min_height = min_height;
  geometry.max_width = G_MAXINT;
  geometry.max_height = G_MAXINT;

  gdk_window_set_geometry_hints(m_gdk_window, &geometry, static_cast<GdkWindowHints>(GDK_HINT_MIN_SIZE | GDK_HINT_MAX_SIZE));
}

bool RNWindow::OnDelete() const {
  fl_method_channel_invoke_method(m_channel, "onTryClose", nullptr, nullptr, nullptr, nullptr);
  return true;
}

bool RNWindow::OnFocusChange(GdkEventFocus* event) const {
  g_autoptr(FlValue) result_data = fl_value_new_map();
  fl_value_set_string_take(result_data, "hasFocus", fl_value_new_bool(event->in == TRUE));
  fl_method_channel_invoke_method(m_channel, "onFocusChange", result_data, nullptr, nullptr, nullptr);

  return false;
}

bool RNWindow::OnConfigure(GdkEventConfigure* event) const {
  g_autoptr(FlValue) result_data = fl_value_new_map();
  fl_value_set_string_take(result_data, "x", fl_value_new_int(event->x));
  fl_value_set_string_take(result_data, "y", fl_value_new_int(event->y));
  fl_value_set_string_take(result_data, "width", fl_value_new_int(event->width));
  fl_value_set_string_take(result_data, "height", fl_value_new_int(event->height));
  fl_method_channel_invoke_method(m_channel, "onConfigure", result_data, nullptr, nullptr, nullptr);

  return false;
}

void RNWindow::HandleMethodCall(FlMethodCall* method_call) {
  g_autoptr(FlMethodResponse) response = nullptr;

  const gchar* method = fl_method_call_get_name(method_call);
  FlValue* args = fl_method_call_get_args(method_call);

  if (strcmp(method, "initPlugin") == 0) {
    InitPlugin();
    response = flBool(true);
  } else if (strcmp(method, "closePlugin") == 0) {
    ClosePlugin();
    response = flBool(true);
  } else if (strcmp(method, "isVisible") == 0) {
    response = flBool(IsVisible());
  } else if (strcmp(method, "initialShow") == 0) {
    Geometry geometry(args);
    auto min_width = fl_value_get_int(fl_value_lookup_string(args, "min_width"));
    auto min_height = fl_value_get_int(fl_value_lookup_string(args, "min_height"));
    SetGeometry(geometry);
    SetGeometryHint(min_width, min_height);
    Show();
    response = flBool(true);
  } else if (strcmp(method, "show") == 0) {
    Show();
    response = flBool(true);
  } else if (strcmp(method, "hide") == 0) {
    Hide();
    response = flBool(true);
  } else if (strcmp(method, "isFocused") == 0) {
    response = flBool(IsFocused());
  } else if (strcmp(method, "focus") == 0) {
    Focus();
    response = flBool(true);
  } else if (strcmp(method, "setOpacity") == 0) {
    double opacity = fl_value_get_float(fl_value_lookup_string(args, "opacity"));
    SetOpacity(opacity);
    response = flBool(true);
  } else if (strcmp(method, "getGeometry") == 0) {
    Geometry geometry;
    GetGeometry(geometry);
    response = geometry.ToFlResponse();
  } else if (strcmp(method, "setGeometry") == 0) {
    Geometry geometry(args);
    SetGeometry(geometry);
    response = flBool(true);
  } else {
    response = flNotImplemented();
  }

  fl_method_call_respond(method_call, response, nullptr);
}

bool RNWindow::HandleGtkSignal(GdkEvent* event) {
  RNWindow* self = RNWindow::instance;
  if (self == nullptr) {
    return false;
  }

  switch (event->type) {
    case GDK_FOCUS_CHANGE:
      return self->OnFocusChange(reinterpret_cast<GdkEventFocus*>(event));
    case GDK_CONFIGURE:
      return self->OnConfigure(reinterpret_cast<GdkEventConfigure*>(event));
    case GDK_DELETE:
      return self->OnDelete();
    default:
      return false;
  }
}
