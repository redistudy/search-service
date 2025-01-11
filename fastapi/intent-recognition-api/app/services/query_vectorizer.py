from typing import Dict, List
from fastapi import HTTPException
from app.services.model import model 
import numpy as np

class QueryVectorizer:
    def __init__(self):
        self.classifier = model.load_vector_model()
    
    # query를 벡터로 변환
    def vectorized_text(self, text: str) -> Dict[str, List[float]]:
        if not text:
            raise HTTPException(status_code=400, detail="Query is empty.")

        # query를 벡터로 변환
        embeddings = self.classifier(text)
        vector = np.mean(embeddings[0], axis = 0).tolist()
        print(len(vector))
        # 결과 반환 (JSON)
        return {
            "vector": vector
        }

query_vectorizer = QueryVectorizer()

# vectorized_text 함수를 호출하여 query를 벡터로 변환
def vectorized_text(text: str) -> Dict[str, List[float]]:
    return query_vectorizer.vectorized_text(text)