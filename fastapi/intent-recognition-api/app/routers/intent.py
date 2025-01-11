from fastapi import APIRouter, HTTPException
from app.models.intent import IntentRequest, IntentResponse
from app.services.intent_classifier import classify_intent


async def recognize_intent(request: IntentRequest):
    try:
        intent = classify_intent(request.text)
        print(intent)
        return IntentResponse(intent=intent["intent"], confidence=intent["confidence"])
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))
    
