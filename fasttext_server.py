# fasttext_server.py
from flask import Flask, request, jsonify
import fasttext

app = Flask(__name__)
model = fasttext.load_model('cc.en.300.bin')  # Завантажте модель FastText

@app.route('/vectorize', methods=['POST'])
def vectorize():
    word = request.json.get('word')
    vector = model.get_word_vector(word).tolist()
    return jsonify(vector=vector)

if __name__ == '__main__':
    app.run(port=5000)
