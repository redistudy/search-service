from transformers import pipeline
from typing import Dict, List
from app.services.model import model
from app.services.query.query_factory import QueryFactory

class IntentClassifier:
    def __init__(self):
        self.classifier = model.load_model()
        self.candiate_labels = [
            "address",
            "poi_name",
            "poi_name_address"
        ]

    def classify_intent(self, input_data: str) -> Dict[str, List[str]]:
        try:
            result = self.classifier(input_data, self.candiate_labels) 
            # 예측된 인텐트와 확률 반환 labels : 인텐트, confidence : 확률
            predicated_intent = result["labels"][0]
            confidence = result["scores"][0]
            return {
                "intent": predicated_intent,
                "confidence": float(confidence)
            }
        except Exception as e:
            print(f"Error in classify_intent: {e}")
            return {
                "intent": "unknown",
                "confidence": 0.0
            }
    def generate_query_by_intent(self, intent: str) -> Dict[str, str]:
        query_generator = QueryFactory.create_query(intent)
        return query_generator.generate_query(intent)
    
intent_classifier = IntentClassifier()

def classify_intent(input_data: str) -> Dict[str, List[str]]:
    return intent_classifier.classify_intent(input_data)

def generate_query_by_intent(intent: str) -> Dict[str, str]:
    query = intent_classifier.generate_query_by_intent(intent)
    print('query:', query)
    return {'query' : query}