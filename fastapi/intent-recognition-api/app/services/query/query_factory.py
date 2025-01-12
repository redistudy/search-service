from enum import Enum
from app.services.query.query_interface import QueryInterface
from app.services.query.name_query import NameQuery
from app.services.query.address_query import AddressQuery
from app.services.query.address_name_query import AddressNameQuery

class QueryType(Enum):
    NAME = "name"
    ADDRESS = "address"
    ADDRESS_NAME = "poi_name_address"

class QueryFactory:
    @staticmethod
    def create_query(query_type: str) -> QueryInterface:
        query_map = {
            QueryType.NAME.value: NameQuery(),
            QueryType.ADDRESS.value: AddressQuery(),
            QueryType.ADDRESS_NAME.value: AddressNameQuery()
        }
        return query_map.get(query_type, NameQuery())
    
query_factory = QueryFactory()