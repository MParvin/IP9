from flask import Flask, request

app = Flask(__name__)

browserList = ['chrome','edge','firefox','konqueror','mozilla','msie','netscape','opera','safari']
@app.route('/')
def default_index():
    if request.user_agent.browser not in browserList:
       return request.remote_addr+"\n"
    else:
       return request.remote_addr

if __name__ == "__main__":
   app.run(debug=False, port=5000, host='0.0.0.0')
