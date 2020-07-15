from flask import Flask
app = Flask(__name__)


@app.route('/')
def index():
    return "Locus load test main page."


@app.route('/hello')
def hello_world():
    return 'Hello, World!'


@app.route('/item')
def items():
    return 'This is items list'

if __name__ == "__main__":
    app.run(host='0.0.0.0', port=8080, debug=True)