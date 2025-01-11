from fastapi import FastAPI
from app.routers.router import router

app = FastAPI()

app.include_router(router=router)

@app.get("/")
def read_root():
    return {"message": "Welcome to the Intent Recognition API"}