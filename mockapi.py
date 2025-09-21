from flask import Flask, jsonify
import random

app = Flask(__name__)

@app.route("/temp/<device_id>")
def temp(device_id):
    return jsonify({"value": round(20 + random.random()*10, 2)})

@app.route("/hum/<device_id>")
def hum(device_id):
    return jsonify({"value": round(40 + random.random()*20, 2)})

@app.route("/air/<device_id>")
def air(device_id):
    return jsonify({"value": round(0 + random.random()*100, 2)})

if __name__ == "__main__":
    app.run(port=8081)
