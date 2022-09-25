#pragma once

#include <flutter_linux/flutter_linux.h>


struct Geometry;
class RNWindow {
public:
    RNWindow(GtkWindow* gtk_window, FlMethodChannel* channel);
    ~RNWindow();

public:
    void InitPlugin(const Geometry& g);
    void ClosePlugin();
    bool IsVisible() const;
    void Show() const;
    void Hide() const;
    bool IsFocused() const;
    void Focus() const;
    void SetOpacity(double opacity) const;
    void GetGeometry(Geometry& g) const;
    void SetGeometry(const Geometry& g) const;
    void SetGeometryHint(int min_width, int min_height) const;

private:
    bool OnDelete() const;
    bool OnFocusChange(GdkEventFocus* event) const;
    bool OnConfigure(GdkEventConfigure* event) const;

public:
    void HandleMethodCall(FlMethodCall* method_call);
    static bool HandleGtkSignal(GdkEvent* event);

private:
    static RNWindow* instance;

    gulong m_delete_handler = 0;
    gulong m_focus_in_handler = 0;
    gulong m_focus_out_handler = 0;
    gulong m_configure_handler = 0;

    GtkWindow* m_gtk_window = nullptr;
    GdkWindow* m_gdk_window = nullptr;
    FlMethodChannel* m_channel = nullptr;

    int m_monitor_hwidth_px;
    int m_monitor_hheight_px;
    float m_window_width_ppm; // pixel per mm
    float m_window_height_ppm; // pixel per mm
};
