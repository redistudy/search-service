import json
from fastapi import APIRouter
from app.models.intent import IntentRequest, IntentResponse, SearchRequest, SearchResponse, SearchQueryRequest, SearchQueryResponse
from app.services.intent_classifier import generate_query_by_intent, classify_intent
from app.services.query_vectorizer import vectorized_text

router = APIRouter()

@router.post("/recognize-intent", response_model=IntentResponse)
async def recognize_intent(request: IntentRequest):
    intent = classify_intent(request.text)
    return intent

@router.post("/vectorized-text", response_model=SearchResponse)
async def vectorize(request: SearchRequest):
    embeded = vectorized_text(request.text)
    return embeded

@router.post("/query", response_model=SearchQueryResponse)
async def get_query(request: SearchQueryRequest):
    intent = classify_intent(request.text)
    print(intent)
    query = generate_query_by_intent(intent['intent'])
    return {'script' : query}