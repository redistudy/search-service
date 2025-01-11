from pydantic import BaseModel

# 인텐트 검색 Request
class IntentRequest(BaseModel):
    text: str
    
# 인텐트 검색 Response
class IntentResponse(BaseModel):
    intent: str
    confidence: float

# 텍스트 검색 Request
class SearchRequest(BaseModel):
    text: str

class SearchResponse(BaseModel):
    vector: list[float]