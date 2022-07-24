import os
import json
from xdg.DesktopEntry import DesktopEntry
from xdg.BaseDirectory import xdg_data_dirs


apps = {}
for root_dir in xdg_data_dirs[::-1]:
    app_dir = os.path.join(root_dir, "applications")
    if not os.path.exists(app_dir):
        continue

    for root, dirnames, filenames in os.walk(app_dir):
        for filename in filenames:
            if not filename.endswith(".desktop"):
                continue

            full_path = os.path.join(root, filename)
            obj = DesktopEntry(full_path)
            if obj.getNoDisplay() or obj.getHidden():
                continue

            apps[obj.getName()] = {"name": obj.getName(), "app": full_path}

print(json.dumps([it for it in apps.values()], ensure_ascii=False))
