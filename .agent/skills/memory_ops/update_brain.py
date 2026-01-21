import json
import sys

def update_brain(key, value):
    try:
        with open('project_brain.json', 'r+') as f:
            data = json.load(f)
            keys = key.split('.')
            ref = data
            for k in keys[:-1]:
                ref = ref[k]
            ref[keys[-1]] = value
            f.seek(0)
            json.dump(data, f, indent=2)
            f.truncate()
    except Exception as e:
        print(f"Error updating brain: {e}")

if __name__ == "__main__":
    if len(sys.argv) > 2:
        update_brain(sys.argv[1], sys.argv[2])
