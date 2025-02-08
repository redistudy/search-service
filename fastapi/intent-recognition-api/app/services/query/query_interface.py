from typing import Dict, List
from abc import ABC, abstractmethod

class QueryInterface(ABC):
    @abstractmethod
    def generate_query(self, query_embedding: List[float]) -> Dict[str, str]:
        pass
    