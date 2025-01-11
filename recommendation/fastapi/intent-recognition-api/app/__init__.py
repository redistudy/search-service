# Intent Recognition API Initialization

from fastapi import FastAPI
from .routers import intent
from app.routers.router import router

app = FastAPI()

app.include_router(router)