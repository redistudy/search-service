from fastapi import APIRouter
from app.models.intent import IntentRequest, IntentResponse, SearchRequest, SearchResponse
from app.services.intent_classifier import intent_classifier
from app.services.query_vectorizer import query_vectorizer

router = APIRouter()

@router.post("/recognize-intent", response_model=IntentResponse)
async def recognize_intent(request: IntentRequest):
    return intent_classifier.classify_intent(request.text)


@router.post("/vectorized-text", response_model=SearchResponse)
async def vectorized_text(request: SearchRequest):
    return query_vectorizer.vectorized_text(request.text)