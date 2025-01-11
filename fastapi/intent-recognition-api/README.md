# intent-recognition-api/intent-recognition-api/README.md

# Intent Recognition API

This project is a FastAPI application that provides an API for intent recognition. It allows users to classify intents based on input data.

## Project Structure

```
ntent-recognition-api/
├── app/
│   ├── __init__.py
│   ├── main.py
│   ├── routers/
│   │   ├── intent.py
│   |   ├── query.py
│   |   └── router.py
│   ├── models/
│   │   └── intent.py      
│   └── services/
│       ├── model.py         
│       ├── intent_classifier.py 
│       └── query_vectorizer.py 
├── tests/
│   ├── __init__.py
│   ├── test_intent.py      
│   └── test_vectorizer.py 
├── requirements.txt       
└── README.md               

## Setup Instructions

1. Clone the repository:
   ```
   git clone <repository-url>
   cd intent-recognition-api
   ```

2. Install the required dependencies:
   ```
   pip install -r requirements.txt
   ```

3. Run the FastAPI application:
   ```
   uvicorn app.main:app --reload
   ```

## Usage

Once the application is running, you can access the API at `http://127.0.0.1:8000`. The API documentation is available at `http://127.0.0.1:8000/docs`.

## Endpoints

- **POST /intent**: Classifies the intent based on the provided input data.

## Testing

To run the tests, use the following command:
```
pytest tests/test_intent.py
```

## License

This project is licensed under the MIT License.