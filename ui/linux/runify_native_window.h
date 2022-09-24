#pragma once

#include <flutter_linux/flutter_linux.h>


struct Geometry;
class RNWindow {
public:
    RNWindow(GtkWindow* gtk_window);
    ~RNWindow() = default;

public:
    void Init(const Geometry& g);
    bool IsVisible();
    void Show();
    void Hide();
    void Close();
    bool IsFocused();
    void Focus();
    void SetOpacity(double opacity);
    void GetGeometry(Geometry& g);
    void SetGeometry(const Geometry& g);
    void SetGeometryHint(int min_width, int min_height);
    void HandleMethodCall(FlMethodCall* method_call);

public:
    bool EnableClose() const {
        return m_enable_close;
    }

private:
    GtkWindow* m_gtk_window = nullptr;
    GdkWindow* m_gdk_window = nullptr;

    int m_monitor_hwidth_px;
    int m_monitor_hheight_px;
    float m_window_width_ppm; // pixel per mm
    float m_window_height_ppm; // pixel per mm
    bool m_enable_close = false;
};
