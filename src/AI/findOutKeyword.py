from flask import Flask, request, jsonify
from textrank.TextRank import textRank



app = Flask(__name__)

@app.route('/extract', methods=['POST'])
def extract_keywords_and_summary():
    data = request.get_json()
    text = data.get('text', '')

    if not text:
        return jsonify({"error": "Text not provided"}), 400

    # Initialize TextRank with the provided text
    T = textRank.TextRank(text, pr_config={'alpha': 0.85, 'max_iter': 100})

    # Extract top 10 keywords
    keywords = T.get_n_keywords(10)

    # Format the response
    response = {"keywords": [{"keyword": kw[0], "weight": kw[1]} for kw in keywords]}

    return jsonify(response)

if __name__ == '__main__':
    app.run(port=5000)
