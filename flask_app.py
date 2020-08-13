from flask import Flask, request

app = Flask(__name__)

@app.route('/')
def default_index():
   return request.remote_addr

if __name__ == "__main__":
   app.run(debug=False, port=5000, host='0.0.0.0')
