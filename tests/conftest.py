import os
import requests
import pytest

BASE = os.getenv("API_BASE_URL", "http://localhost:8080/v1")
USERNAME = os.getenv("USERNAME", "user")
PASSWORD = os.getenv("PASSWORD", "pass")

@pytest.fixture(scope="session")
def access_token():
    r = requests.post(f"{BASE}/auth/login", json={
        "username": USERNAME,
        "password": PASSWORD
    })
    if r.status_code != 200:
        r = requests.post(f"{BASE}/auth/register", json={
            "username": USERNAME,
            "password": PASSWORD
        })
    body = r.json()
    return body["accessToken"]

@pytest.fixture
def auth_headers(access_token):
    return {"Authorization": f"Bearer {access_token}"}
