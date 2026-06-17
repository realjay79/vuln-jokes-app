# The Jokes App 🃏

A deliberately vulnerable (and then fixed) React + Go application, built to demonstrate a complete Secure SDLC process. This repository contains the **remediated version** of the app, as documented in the [companion Medium series](https://jayriniv.medium.com/securing-the-dumbest-app-ive-built-on-purpose-cd4495b47d5b).

## Architecture

- **Frontend**: React (port 3000)
- **Backend**: Go (port 8080)
- **External API**: [Official Joke API](https://official-joke-api.appspot.com/)

## Prerequisites

- **Go** (version 1.19 or later)
- **Node.js** (version 16.x or later) and **npm** / **yarn**

## Quick Start

1.  **Clone the repository**
    ```bash
    git clone https://github.com/realjay79/vuln-jokes-app.git
    cd vuln-jokes-app

**2. Run the Backend (Go)**

bash
cd backend
go mod download
go run main.go
The API server will start on http://localhost:8080.

**3. Run the Frontend (React)**
Open a new terminal window.

bash
cd frontend
npm install
npm start
The React development server will start on http://localhost:3000.

Open the App
Navigate to http://localhost:3000 in your browser. You should see the Jokes App interface.

**API Endpoints**
Method	Endpoint	Description
GET	/joke?id={joke_id}	Fetches a joke by ID from the Official Joke API
Example Request:

```
GET http://localhost:8080/joke?id=123

```
Example Response: 
json
{
  "id": 123,
  "type": "general",
  "setup": "Why did the scarecrow win an award?",
  "punchline": "Because he was outstanding in his field."
}



**Security Notes**
This repository contains the fixed version of the application. The vulnerabilities documented in the blog series have been remediated.

Known vulnerability (now fixed): The original code allowed arbitrary URL injection via the id parameter, which could enable SSRF attacks. The fixed version validates and sanitizes the input.

**Development & Testing**
**Running Scans (SAST/SCA)**
To replicate the security scans discussed in the blog series:

**Semgrep (SAST/SCA):**

**In the root directory**
semgrep --config p/default --config p/security-audit ./backend ./frontend
(Note: Using local custom rules will provide more thorough results, as noted in the series.)

**OWASP ZAP (DAST):**
Ensure the application is running on localhost:3000 and localhost:8080.

Open ZAP and set the target URL to http://localhost:3000.
