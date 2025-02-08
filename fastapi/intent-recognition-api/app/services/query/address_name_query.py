from typing import Dict, List


class AddressNameQuery:
    def __init__(self):
        pass
    def generate_query(self, query_embedding : List[float]) -> Dict[str, str]:
        return {
            "_source": ["title", "address", "location"],
            "query": {
            "script_score": {
                "query": {
                    "match_all": {}
                },
                "script": {
                    "source": "\
                        double titleWeight = params.title_weight;\
                        double addressWeight = params.address_weight;\
                        return titleWeight * cosineSimilarity(params.query_vector, 'title_vector') + \
                        addressWeight * cosineSimilarity(params.query_vector, 'address_vector');\
                    ",
                "params": {
                    "query_vector": [],
                    "title_weight": 1.0, 
                    "address_weight": 1.0
                }
                }
            }
            }
        }


address_name_query = AddressNameQuery()

def generate_query(intent: str) -> Dict[str, str]:
    return address_name_query.generate_query(intent)