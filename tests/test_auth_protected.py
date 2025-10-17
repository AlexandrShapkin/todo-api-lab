import os
import requests
from datetime import datetime, timedelta, timezone

BASE = os.getenv("API_BASE_URL", "http://localhost:8080/v1")

INVALID_TOKEN = "Bearer invalid.token.value"


def test_me_with_invalid_token():
    r = requests.post(f"{BASE}/auth/me", headers={
        "Authorization": INVALID_TOKEN
    })
    assert r.status_code in (401, 403), f"Expected 401/403, got {r.status_code}: {r.text}"


def test_me_without_token():
    r = requests.post(f"{BASE}/auth/me")
    assert r.status_code in (401, 403), f"Expected 401/403, got {r.status_code}: {r.text}"


def test_create_task_with_invalid_token():
    due = (datetime.now(timezone.utc) + timedelta(days=1)).replace(microsecond=0).isoformat().replace("+00:00", "Z")
    r = requests.post(f"{BASE}/tasks", headers={"Authorization": INVALID_TOKEN}, json={
        "title": "Should Fail",
        "description": "Testing invalid token",
        "dueTime": due,
        "completed": False
    })
    assert r.status_code in (401, 403), f"Expected 401/403, got {r.status_code}: {r.text}"


def test_get_tasks_with_invalid_token():
    r = requests.get(f"{BASE}/tasks", headers={"Authorization": INVALID_TOKEN})
    assert r.status_code in (401, 403), f"Expected 401/403, got {r.status_code}: {r.text}"


def test_delete_task_with_invalid_token():
    fake_task_id = "00000000-0000-0000-0000-000000000000"
    r = requests.delete(f"{BASE}/tasks/{fake_task_id}", headers={"Authorization": INVALID_TOKEN})
    assert r.status_code in (401, 403), f"Expected 401/403, got {r.status_code}: {r.text}"
