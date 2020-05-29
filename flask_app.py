from flask import Flask, request

app = Flask(__name__)

@app.route('/')
def default_index():
   return request.remote_addr
