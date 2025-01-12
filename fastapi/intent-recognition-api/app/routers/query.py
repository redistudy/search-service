from fastapi import APIRouter, HTTPException
from app.services.query_vectorizer import vectorized_text
from app.services.intent_classifier import intent_classifier

# convert query to embedding vector router
async def vectorized_text(text: str):
    try:
        vector = vectorized_text(text)
        return vector
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))

# get query router
async def get_query(text : str):
    try:
        intent = intent_classifier.classify_intent(text)
        query = intent_classifier.generate_query_by_intent(intent)
        return query
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))