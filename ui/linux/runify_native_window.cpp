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

RNWindow::RNWindow(GtkWindow* gtk_window)
    : m_gtk_window(gtk_window)
    , m_gdk_window(gtk_widget_get_window(GTK_WIDGET(gtk_window))) {

    GdkDisplay* display = gdk_display_get_default();
    GdkMonitor* monitor = gdk_display_get_monitor_at_window(display, m_gdk_window);
    GdkRectangle frame;
    gdk_monitor_get_geometry(monitor, &frame);

    m_monitor_hwidth_px = frame.width / 2;
    m_monitor_hheight_px = frame.height / 2;
    m_window_width_ppm = static_cast<float>(frame.width) / static_cast<float>(gdk_monitor_get_width_mm(monitor));
    m_window_height_ppm = static_cast<float>(frame.height) / static_cast<float>(gdk_monitor_get_height_mm(monitor));
}

void RNWindow::Init(const Geometry& g) {
  SetGeometry(g);
  SetGeometryHint(480, 640);
}

bool RNWindow::IsVisible() {
  return gtk_widget_is_visible(GTK_WIDGET(m_gtk_window));
}

void RNWindow::Show() {
  gtk_widget_show(GTK_WIDGET(m_gtk_window));
  gtk_window_present(m_gtk_window);
}

void RNWindow::Hide() {
  gtk_widget_hide(GTK_WIDGET(m_gtk_window));
}

void RNWindow::Close() {
  m_enable_close = true;
  gtk_window_close(m_gtk_window);
}

bool RNWindow::IsFocused() {
  return gtk_window_is_active(m_gtk_window);
}

void RNWindow::Focus() {
  gtk_window_present(m_gtk_window);
}

void RNWindow::SetOpacity(double opacity) {
    gtk_widget_set_opacity(GTK_WIDGET(m_gtk_window), opacity);
}

void RNWindow::GetGeometry(Geometry& g) {
    gint x_px, y_px, width_px, height_px;
    gtk_window_get_position(m_gtk_window, &x_px, &y_px);
    gtk_window_get_size(m_gtk_window, &width_px, &height_px);

    // Result is in millimeter
    g.x = static_cast<float>(x_px + width_px / 2 - m_monitor_hwidth_px) / m_window_width_ppm;
    g.y = static_cast<float>(y_px + height_px / 2 - m_monitor_hheight_px) / m_window_height_ppm;
    g.width = static_cast<float>(width_px) / m_window_width_ppm;
    g.height = static_cast<float>(height_px) / m_window_height_ppm;
}

void RNWindow::SetGeometry(const Geometry& g) {
    gint width_px = static_cast<gint>(g.width * m_window_width_ppm + 0.5);
    gint height_px = static_cast<gint>(g.height * m_window_height_ppm + 0.5);
    gint x_px = static_cast<gint>((g.x - g.width * 0.5) * m_window_width_ppm + 0.5) + m_monitor_hwidth_px;
    gint y_px = static_cast<gint>((g.y - g.height * 0.5) * m_window_height_ppm + 0.5) + m_monitor_hheight_px;

    gdk_window_move_resize(m_gdk_window, x_px, y_px, width_px, height_px);
}

void RNWindow::SetGeometryHint(int min_width, int min_height) {
    GdkGeometry geometry;
    geometry.min_width = min_width;
    geometry.min_height = min_height;
    geometry.max_width = G_MAXINT;
    geometry.max_height = G_MAXINT;

    gdk_window_set_geometry_hints(m_gdk_window, &geometry, static_cast<GdkWindowHints>(GDK_HINT_MIN_SIZE | GDK_HINT_MAX_SIZE));
}

void RNWindow::HandleMethodCall(FlMethodCall* method_call) {
  g_autoptr(FlMethodResponse) response = nullptr;

  const gchar* method = fl_method_call_get_name(method_call);
  FlValue* args = fl_method_call_get_args(method_call);

  if (strcmp(method, "init") == 0) {
    Geometry geometry(args);
    Init(geometry);
    response = flBool(true);
  } else if (strcmp(method, "isVisible") == 0) {
    response = flBool(IsVisible());
  } else if (strcmp(method, "show") == 0) {
    Show();
    response = flBool(true);
  } else if (strcmp(method, "hide") == 0) {
    Hide();
    response = flBool(true);
  } else if (strcmp(method, "close") == 0) {
    Close();
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
