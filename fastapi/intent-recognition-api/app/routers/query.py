from fastapi import APIRouter, HTTPException
from app.services.query_vectorizer import vectorized_text

async def vectorized_text(text: str):
    try:
        vector = vectorized_text(text)
        return vector
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))