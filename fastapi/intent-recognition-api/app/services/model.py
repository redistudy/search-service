from transformers import pipeline

class Model:
    def __init__(self):
        self.classifier = pipeline(
            task="zero-shot-classification",
            model="joeddav/xlm-roberta-large-xnli",
            device=-1
        )
        self.vectorizer = pipeline(
            task="feature-extraction",
            model="joeddav/xlm-roberta-large-xnli",
            device=-1
        )
    def load_model(self):
        return self.classifier
    
    def load_vector_model(self):
        return self.vectorizer


model=Model()

def load_model():
    return model.load_model()

def load_vector_model():
    return model.load_vector_model()