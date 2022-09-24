#pragma once

#include <flutter_linux/flutter_linux.h>


struct Geometry;
class RNWindow {
public:
    RNWindow(GtkWindow* gtk_window, FlMethodChannel* channel);
    ~RNWindow();

public:
    void InitWindow(const Geometry& g) const;
    bool IsVisible() const;
    void Show() const;
    void Hide() const;
    void Close();
    bool IsFocused() const;
    void Focus() const;
    void SetOpacity(double opacity) const;
    void GetGeometry(Geometry& g) const;
    void SetGeometry(const Geometry& g) const;
    void SetGeometryHint(int min_width, int min_height) const;

public:
    bool EnableClose() const { return m_enable_close; }
    void EmitEvent(const char* event_name) const;
    void HandleMethodCall(FlMethodCall* method_call);

public:
    static RNWindow* instance;

private:
    GtkWindow* m_gtk_window = nullptr;
    GdkWindow* m_gdk_window = nullptr;
    FlMethodChannel* m_channel = nullptr;

    int m_monitor_hwidth_px;
    int m_monitor_hheight_px;
    float m_window_width_ppm; // pixel per mm
    float m_window_height_ppm; // pixel per mm
    bool m_enable_close = false;
};
