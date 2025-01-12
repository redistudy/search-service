from typing import Dict, List
from app.services.query.query_interface import QueryInterface

class AddressQuery(QueryInterface):
    def __init__(self):
        pass
    def generate_query(self, query_embedding : List[float]) -> str:
        return '''{
            "_source": ["title", "address", "location"],
            "query": {
            "script_score": {
                "query": {
                "match_all": {}  # 모든 문서에서 스코어 기반 필터링
                },
                "script": {
                "source": """
                    double titleWeight = params.title_weight;
                    double addressWeight = params.address_weight;
                    return titleWeight * cosineSimilarity(params.query_vector, 'title_vector') + 
                       addressWeight * cosineSimilarity(params.query_vector, 'address_vector');
                """,
                "params": {
                    "query_vector": query_embedding,
                    "title_weight": 0.0,  # 기본 가중치 값
                    "address_weight": 1.0  # 기본 가중치 값
                }
                }
            }
            }
        }'''


address_query = AddressQuery()

def generate_query(intent: str) -> Dict[str, str]:
    return address_query.generate_query(intent)