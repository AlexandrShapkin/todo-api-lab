import os
import requests
from datetime import datetime, timedelta, timezone

BASE = os.getenv("API_BASE_URL", "http://localhost:8080/v1")

def test_create_task(auth_headers):
    due = (datetime.now(timezone.utc) + timedelta(days=1)).replace(microsecond=0).isoformat().replace("+00:00", "Z")
    r = requests.post(f"{BASE}/tasks", headers=auth_headers, json={
        "title": "Test Task",
        "description": "For API test",
        "dueTime": due,
        "completed": False
    })
    assert r.status_code == 201, f"Unexpected status {r.status_code}: {r.text}"

    try:
        task = r.json()
    except ValueError:
        assert False, f"Response is not JSON: {r.text}"

    os.environ["TASK_ID"] = task["id"]

def test_get_task(auth_headers):
    task_id = os.getenv("TASK_ID")
    r = requests.get(f"{BASE}/tasks/{task_id}", headers=auth_headers)
    assert r.status_code == 200, f"Unexpected status {r.status_code}: {r.text}"

    try:
        task = r.json()
    except ValueError:
        assert False, f"Response is not JSON: {r.text}"

    assert task["title"] == "Test Task", f"Unexpected title: {task}"

def test_patch_task(auth_headers):
    task_id = os.getenv("TASK_ID")
    r = requests.patch(f"{BASE}/tasks/{task_id}", headers=auth_headers, json={
        "completed": True
    })
    assert r.status_code == 200, f"Unexpected status {r.status_code}: {r.text}"

    try:
        task = r.json()
    except ValueError:
        assert False, f"Response is not JSON: {r.text}"

    assert task["completed"] is True, f"Unexpected completed field: {task}"

def test_delete_task(auth_headers):
    task_id = os.getenv("TASK_ID")
    r = requests.delete(f"{BASE}/tasks/{task_id}", headers=auth_headers)
    assert r.status_code in (200, 204), f"Unexpected status {r.status_code}: {r.text}"
