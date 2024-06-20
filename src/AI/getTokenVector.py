from flask import Flask, request, jsonify
import gensim
import numpy as np

app = Flask(__name__)


def load_tencent_embedding(file_path):
    """
    Load Tencent AI Lab Embedding from the given file path.
    """
    model = gensim.models.KeyedVectors.load_word2vec_format(file_path, binary=False)
    return model


def get_word_vector(model, word):
    """
    Get the vector for a given word from the loaded model.
    """
    try:
        vector = model[word]
        return vector
    except KeyError:
        return "Word not in vocabulary"


# 文件路径，请根据你的实际路径进行修改
file_path = './tencent-ailab-embedding-zh-d100-v0.2.0-s.txt'

# 加载模型
print("Loading Tencent AI Lab Embedding model...")
model = load_tencent_embedding(file_path)
print("Model loaded.")


@app.route('/get_vector', methods=['POST'])
def get_vector():
    data = request.json
    word = data.get('word')
    if not word:
        return jsonify({'error': 'No word provided'}), 400

    vector = get_word_vector(model, word)
    if isinstance(vector, str):
        return jsonify({'error': vector}), 400
    else:
        return jsonify({'word': word, 'vector': vector.tolist()})


if __name__ == '__main__':
    app.run(port=8083)
