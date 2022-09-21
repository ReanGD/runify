#include "runify_native_window.h"


static FlMethodResponse* flBool(bool value) {
    g_autoptr(FlValue) result = fl_value_new_bool(value);
    return FL_METHOD_RESPONSE(fl_method_success_response_new(result));
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

FlMethodResponse* RNWindow::IsVisible() {
  bool is_visible = gtk_widget_is_visible(GTK_WIDGET(m_gtk_window));

  return flBool(is_visible);
}

FlMethodResponse* RNWindow::Show() {
  gtk_widget_show(GTK_WIDGET(m_gtk_window));
  gtk_window_present(m_gtk_window);

  return flBool(true);
}

FlMethodResponse* RNWindow::Hide() {
  gtk_widget_hide(GTK_WIDGET(m_gtk_window));

  return flBool(true);
}

FlMethodResponse* RNWindow::IsFocused() {
  bool is_focused = gtk_window_is_active(m_gtk_window);

  return flBool(is_focused);
}

FlMethodResponse* RNWindow::Focus() {
  gtk_window_present(m_gtk_window);

  return flBool(true);
}


FlMethodResponse* RNWindow::SetOpacity(FlValue* args) {
    gdouble opacity = fl_value_get_float(fl_value_lookup_string(args, "opacity"));
    gtk_widget_set_opacity(GTK_WIDGET(m_gtk_window), opacity);

    return flBool(true);
}

FlMethodResponse* RNWindow::GetGeometry() {
    gint x_px, y_px, width_px, height_px;
    gtk_window_get_position(m_gtk_window, &x_px, &y_px);
    gtk_window_get_size(m_gtk_window, &width_px, &height_px);

    float width_mm = static_cast<float>(width_px) / m_window_width_ppm;
    float height_mm = static_cast<float>(height_px) / m_window_height_ppm;

    float x_mm = static_cast<float>(x_px + width_px / 2 - m_monitor_hwidth_px) / m_window_width_ppm;
    float y_mm = static_cast<float>(y_px + height_px / 2 - m_monitor_hheight_px) / m_window_height_ppm;

    g_autoptr(FlValue) result = fl_value_new_map();
    fl_value_set_string_take(result, "x", fl_value_new_float(x_mm));
    fl_value_set_string_take(result, "y", fl_value_new_float(y_mm));
    fl_value_set_string_take(result, "width", fl_value_new_float(width_mm));
    fl_value_set_string_take(result, "height", fl_value_new_float(height_mm));

    return FL_METHOD_RESPONSE(fl_method_success_response_new(result));
}

FlMethodResponse* RNWindow::SetGeometry(FlValue* args) {
    float x_mm = fl_value_get_float(fl_value_lookup_string(args, "x"));
    float y_mm = fl_value_get_float(fl_value_lookup_string(args, "y"));
    float width_mm = fl_value_get_float(fl_value_lookup_string(args, "width"));
    float height_mm = fl_value_get_float(fl_value_lookup_string(args, "height"));

    gint width_px = static_cast<gint>(width_mm * m_window_width_ppm + 0.5);
    gint height_px = static_cast<gint>(height_mm * m_window_height_ppm + 0.5);
    gint x_px = static_cast<gint>((x_mm - width_mm * 0.5) * m_window_width_ppm + 0.5) + m_monitor_hwidth_px;
    gint y_px = static_cast<gint>((y_mm - height_mm * 0.5) * m_window_height_ppm + 0.5) + m_monitor_hheight_px;

    gdk_window_move_resize(m_gdk_window, x_px, y_px, width_px, height_px);

    return flBool(true);
}
