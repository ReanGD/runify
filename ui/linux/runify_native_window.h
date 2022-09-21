#pragma once

#include <flutter_linux/flutter_linux.h>


class RNWindow {
public:
    RNWindow(GtkWindow* gtk_window);
    ~RNWindow() = default;

public:
    FlMethodResponse* IsVisible();
    FlMethodResponse* Show();
    FlMethodResponse* Hide();
    FlMethodResponse* IsFocused();
    FlMethodResponse* Focus();
    FlMethodResponse* SetOpacity(FlValue* args);
    FlMethodResponse* GetGeometry();
    FlMethodResponse* SetGeometry(FlValue* args);

private:
    GtkWindow* m_gtk_window = nullptr;
    GdkWindow* m_gdk_window = nullptr;

    int m_monitor_hwidth_px;
    int m_monitor_hheight_px;
    float m_window_width_ppm; // pixel per mm
    float m_window_height_ppm; // pixel per mm
};
