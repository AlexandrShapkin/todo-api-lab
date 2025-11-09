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
    assert r.status_code == 201, f"Unexpected status {r.status_code}: {r.text}"

    body = r.json()
    for key in ["username", "userId", "accessToken", "refreshToken"]:
        assert key in body, f"Missing {key} in response"
    os.environ["ACCESS_TOKEN"] = body["accessToken"]
    os.environ["REFRESH_TOKEN"] = body["refreshToken"]

def test_login():
    r = requests.post(f"{BASE}/auth/login", json={
        "username": USERNAME,
        "password": PASSWORD
    })
    assert r.status_code == 200, f"Unexpected status {r.status_code}: {r.text}"
    body = r.json()
    for key in ["username", "userId", "accessToken", "refreshToken"]:
        assert key in body, f"Missing {key} in response"
    os.environ["ACCESS_TOKEN"] = body["accessToken"]
    os.environ["REFRESH_TOKEN"] = body["refreshToken"]

def test_refresh():
    r = requests.post(f"{BASE}/auth/refresh", json={
        "refreshToken": os.getenv("REFRESH_TOKEN")
    })
    assert r.status_code == 200, f"Unexpected status {r.status_code}: {r.text}"
    body = r.json()
    for key in ["username", "userId", "accessToken", "refreshToken"]:
        assert key in body, f"Missing {key} in response"
    os.environ["ACCESS_TOKEN"] = body["accessToken"]
    os.environ["REFRESH_TOKEN"] = body["refreshToken"]

def test_me():
    token = os.getenv("ACCESS_TOKEN")
    r = requests.get(f"{BASE}/auth/me", headers={
        "Authorization": f"Bearer {token}"
    })
    assert r.status_code == 200, f"Unexpected status {r.status_code}: {r.text}"
    body = r.json()
    assert body["username"] == USERNAME
    assert "userId" in body