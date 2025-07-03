import subprocess
import json

class DotenvReader:
    def __init__(self, binary_path="./universal-dotenv"):
        self.binary_path = binary_path

    def read(self,state):
        try:
            output = subprocess.check_output([self.binary_path,state], text=True)
            try:
                return json.loads(output)
            except json.JSONDecodeError:
                return output.strip()  # fallback: return raw text
        except subprocess.CalledProcessError as e:
            raise RuntimeError(f"Failed to run dotenv-reader: {e}")
reader = DotenvReader()
env = reader.read("true")

print(env)
