def test_intent_classification(client):
    response = client.post("/intent", json={"text": "What is the weather like?"})
    assert response.status_code == 200
    assert "intent" in response.json()

def test_intent_classification_invalid_input(client):
    response = client.post("/intent", json={"text": ""})
    assert response.status_code == 400
    assert "detail" in response.json()