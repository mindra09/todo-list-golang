# Run Project 
go run todo-app/cmd/server/main 


# 🧩 Collaborative To-Do List (AI Assisted)

This document explains the **design architecture** for a simple **Collaborative To-Do / Ticket Board** application built with **Go (Gorilla Mux)** and a basic **HTML/CSS/JavaScript frontend**.

The goal of this README is to make **junior developers clearly understand how to implement the system**, while showing **good system design practice**.

---

## 📌 Overview

This application allows users to:

* Add, view, and delete todos (tickets)
* Collaborate through a shared board
* Automatically predict todo category using an **AI service**
* Refine AI predictions and keep refined data for future training

The backend uses **in-memory storage (slice)** to keep the implementation simple.

---

## 🏗 High Level Architecture

```
Frontend (HTML / JS)
       |
       | REST API (JSON)
       v
Go Backend (Gorilla Mux)
 ├── HTTP Handler
 ├── Business Logic (Usecase)
 ├── In-Memory Storage
 ├── AI Integration Service
       |
       v
External AI Service
```

### Explanation

* **Frontend** communicates with backend using REST APIs
* **Backend** handles business logic and AI integration
* **AI Service** is a third-party system for category prediction

---

## 🗂 ERD – Database Model (Conceptual)

> Even though this app does not use a database, designing the data model helps future scalability.

### Entity: `Todo`

```
+-------------------------+
|          Todo           |
+-------------------------+
| id (int, PK)            |
| text (string)           |
| completed (boolean)     |
| ai_category (string)    |
| category (string)       |
| refined (boolean)       |
| created_at (timestamp)  |
+-------------------------+
```

### Field Description

| Field       | Description                         |
| ----------- | ----------------------------------- |
| id          | Unique identifier                   |
| text        | Todo or ticket description          |
| completed   | Completion status                   |
| ai_category | Category predicted by AI            |
| category    | Final category (can be refined)     |
| refined     | Indicates if user refined AI result |
| created_at  | Creation timestamp                  |

---

## 🔄 Communication Flow

### 1️⃣ Create Todo + AI Prediction

```
User
 │
 │ Submit todo + optional PDF
 ▼
Frontend
 │ POST /todos
 ▼
Backend
 │ Call AI Service
 ▼
AI Service
 │ Return category
 ▼
Backend
 │ Save todo in memory
 ▼
Frontend (JSON Response)
```

### Explanation

1. User creates a todo and optionally uploads a PDF
2. Frontend converts PDF to base64
3. Backend forwards base64 to AI service
4. AI predicts category
5. Backend saves todo and returns result

---

### 2️⃣ Get All Todos

```
Frontend
 │ GET api/todos
 ▼
Backend
 │ Read from in-memory slice
 ▼
Frontend
```

---

### 3️⃣ Delete Todo

```
Frontend
 │ DELETE api/todos/{id}
 ▼
Backend
 │ Remove todo from slice
 ▼
Frontend
```

---

## 🧠 AI Integration Design

### AI Request

```json
{
  "file": "base64_stringify_pdf"
}
```

### AI Response

```json
{
  "data": {
    "category": "Bug"
  },
  "status": "success",
  "message": "ok"
}
```

### Notes

* File is **optional**
* If file is missing or AI fails, backend uses a default category

---

## 🔁 Human-in-the-Loop (Refinement)

```
AI Prediction
     |
     v
User Review
     |
     v
Refined Category
```

* AI provides initial category
* User may refine it
* Refined result is marked and stored for training purposes

---

## 🧱 Backend Component Diagram

```
+-------------------+
|   HTTP Handler    |
+---------+---------+
          |
          v
+---------+---------+
|     Usecase       |
| (Business Logic)  |
+---------+---------+
          |
          v
+---------+---------+
|   Repository      |
| (In-Memory Data)  |
+-------------------+

          |
          v
+-------------------+
|   AI Service      |
| (External API)    |
+-------------------+
```

---

## 🤝 Collaboration Concept

```
+-------------------+
|   Shared Board    |
+---------+---------+
          |
          v
+-------------------+
|    Todo Items     |
+-------------------+
```

* All users interact with the same board
* Changes are immediately reflected for everyone

---

## ⚠️ Error Handling Strategy

| Scenario        | Behavior          |
| --------------- | ----------------- |
| Todo not found  | 404 response      |
| Invalid request | 400 response      |
| AI service down | Fallback category |
| Missing file    | Skip AI call      |

---

## 🚀 Future Improvements

* WebSocket for real-time collaboration
* Database (PostgreSQL / MySQL)
* Redis for caching
* User authentication & authorization
* Async worker for AI requests


