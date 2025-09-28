# LogHarbor 🚢

**LogHarbor** is a lightweight log aggregation service built with Go and Gin. Designed for speed, modularity, and scalability, it enables developers to ingest, store, and query structured log data through clean RESTful endpoints.
---

## 🔧 Features

- Fast and minimalistic Go backend using Gin
- RESTful API for log ingestion (`POST /logs`)
- Log retrieval with structured JSON output (`GET /logs`)
- In-memory storage for rapid prototyping
- Modular architecture ready for MongoDB, PostgreSQL, or Kafka integration
- Clean project structure for scalability and maintainability

---

## 📁 Project Structure

```bash
LogHarbor/
├── main.go                  # Entry point
├── go.mod                   # Module definition
├── config/                  # Configuration files
├── models/                  # Log data structures
├── controllers/             # HTTP handlers
├── services/                # Business logic
├── storage/                 # In-memory storage
├── middleware/              # Optional logging/CORS
├── routes/                  # Route definitions
├── utils/                   # Helper functions
```

🚀 Getting Started
```
1. Clone the repository
bash
git clone https://github.com/yourusername/logharbor.git
cd logharbor
2. Install dependencies
bash
go mod tidy
3. Run the server
bash
go run main.go
```
Server will start on http://localhost:5000

📡 API Endpoints
POST /logs
Submit a log entry.
```
json
{
  "level": "INFO",
  "message": "User logged in",
  "source": "auth-service",
  "timestamp": "2025-09-26T19:30:00Z"
}
```
GET /logs
Retrieve all stored logs.

🧱 Built With
Go
Gin

📄 License
This project is licensed under the MIT License. See the LICENSE file for details.

✨ Author
Murad — Full-stack developer & system architect GitHub: @MuradIsazade777

🧠 Future Plans
1. MongoDB and Kafka integration

2. JWT authentication

3. Log filtering by level/source/timestamp

4. Swagger documentation

5. Dockerized deployment
