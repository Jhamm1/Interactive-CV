# Product service

from flask import Flask
from flask_restful import Resource, Api

app = Flask(__name__)
api = Api(app)

class Product(Resource):
    def get (self):
        return {
            'products': 'Software enigneer, focus on test automation',
            'product': 'PhD Researcher in HCI and Healthcare Technology- scholarship funded by EPSRC, Lecturer'
        }
api.add_resource(Product, '/')

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=80, debug=True)
    