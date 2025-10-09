import os
import requests

BASE = os.getenv("API_BASE_URL", "http://localhost:8080/v1")
USERNAME = os.getenv("USERNAME", "user")
PASSWORD = os.getenv("PASSWORD", "pass")

def test_register():
    r = requests.post(f"{BASE}/auth/register", json={
        "username": USERNAME,
        "password": PASSWORD
    })
    assert r.status_code in (200, 201), f"Unexpected status {r.status_code}: {r.text}"
    
    try:
        body = r.json()
    except ValueError:
        assert False, f"Response is not JSON: {r.text}"
    
    assert "accessToken" in body, f"No accessToken in response: {body}"

def test_login():
    r = requests.post(f"{BASE}/auth/login", json={
        "username": USERNAME,
        "password": PASSWORD
    })
    assert r.status_code == 200, f"Unexpected status {r.status_code}: {r.text}"

    try:
        body = r.json()
    except ValueError:
        assert False, f"Response is not JSON: {r.text}"

    assert "accessToken" in body, f"No accessToken in response: {body}"
    os.environ["ACCESS_TOKEN"] = body["accessToken"]

def test_me():
    token = os.getenv("ACCESS_TOKEN")
    r = requests.post(f"{BASE}/auth/me", headers={
        "Authorization": f"Bearer {token}"
    })    

    assert r.status_code == 200, f"Unexpected status {r.status_code}: {r.text}"

    try:
        body = r.json()
    except ValueError:
        assert False, f"Response is not JSON: {r.text}"

    assert body["username"] == USERNAME, f"Unexpected username: {body}"