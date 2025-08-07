# ✂️ Go URL Shortener (No Frameworks)

A simple, fast URL shortener built with **pure Go** and **Redis**, featuring:

- No external web frameworks
- Minimal HTML frontend
- JSON API support
- Redis for storage
- Single binary deployment

---

## 🚀 Features

- Shorten long URLs to short codes
- Redirect short URLs to original destinations
- Redis for key-value storage
- Works with both browser & API
- Fully self-contained (`main.go`)

---

## 📦 Requirements

- Go 1.21+  
- Redis running locally (port `6379`)

---

## 🛠 Setup & Run

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/go-url-shortener.git
cd go-url-shortener
2. Initialize Go Module
bash
Copy
Edit
go mod init url-shortener
go get github.com/go-redis/redis/v8
go get golang.org/x/net
3. Start Redis
bash
Copy
Edit
redis-server
If Redis is not installed:

bash
Copy
Edit
sudo apt update
sudo apt install redis
4. Run the Application
bash
Copy
Edit
go run main.go
Visit: http://localhost:8080

🌐 Usage
🔗 From Browser (HTML Form)
Go to http://localhost:8080

Enter a long URL

Get a shortened URL like:

bash
Copy
Edit
http://localhost:8080/r/abc123
🧪 From API (POST JSON)
bash
Copy
Edit
curl -X POST http://localhost:8080/shorten \
  -H "Content-Type: application/json" \
  -d '{"url": "https://example.com"}'
Response:

json
Copy
Edit
{ "short_url": "http://localhost:8080/r/abc123" }
🗄 Redis Storage
Keys are stored with prefix: short:<code>

Example:

bash
Copy
Edit
redis-cli
> KEYS short:*
> GET short:abc123
📁 Project Structure
css
Copy
Edit
.
├── main.go       # Main logic (HTML, API, Redis, redirect)
└── README.md     # Project documentation
🧹 Reset Redis (Optional)
⚠️ Destroys all data.

bash
Copy
Edit
redis-cli FLUSHALL
📌 License
MIT License — use freely for personal or commercial projects.

✨ Future Ideas
Add analytics for link clicks

Expiring short links

Auth + admin dashboard

Custom aliases

Made with ❤️ using Go and Redis
