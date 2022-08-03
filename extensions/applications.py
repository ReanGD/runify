import os
import json
import tempfile
from pathlib import Path
from cairosvg import svg2png
from xdg.IconTheme import getIconPath
from xdg.DesktopEntry import DesktopEntry
from xdg.BaseDirectory import xdg_data_dirs


cache_file = "cache.json"

if os.path.exists(cache_file):
    with open(cache_file) as f:
        print(f.read())
    exit(0)

tmpdir = os.path.join(tempfile.gettempdir(), "icons")
if not os.path.exists(tmpdir):
    os.mkdir(tmpdir)


def get_icon_path(iconName):
    origin = getIconPath(iconName, 64)
    if origin is None:
        return origin

    origin = Path(origin)
    if origin.suffix != ".svg":
        return origin.as_posix()

    png_path = Path(tmpdir, origin.name)
    png_path = png_path.with_suffix(".png")
    if not png_path.exists():
        svg2png(url=origin.as_posix(), write_to=png_path.as_posix())

    return png_path.as_posix()

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
            icon_path = get_icon_path(obj.getIcon())
            apps[obj.getName()] = {"name": obj.getName(), "app": full_path, "icon": icon_path}

# for i in range(1000):
#     name = f"application {i} name"
#     apps[name] = {"name": name, "app": name}

text = json.dumps([it for it in apps.values()], ensure_ascii=False)

with open(cache_file, "w") as f:
    f.write(text)

print(text)
