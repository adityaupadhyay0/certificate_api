Certificate Management API
==========================

This is a Go-based API for managing certificates, built using the Gin framework.  
It allows users to create, update, retrieve, and upload certificates in bulk via CSV.

Features:
- Retrieve a certificate by ID  
- Create a new certificate  
- Get all certificates  
- Update an existing certificate  
- Upload certificates via CSV file  

Project Structure:
------------------
/certificate-api
│── /uploads            -> Directory for uploaded CSV files
│── /routes             -> API route handlers
│── /models             -> Data models
│── /controllers        -> Business logic for handling requests
│── /utils              -> Utility functions
│── main.go             -> Entry point for the application
│── go.mod              -> Go module file
│── README.txt          -> Project documentation

Installation:
-------------
1. Clone the repository:
   git clone https://github.com/your-username/certificate-api.git
   cd certificate-api

2. Install dependencies:
   go mod tidy

3. Run the server:
   go run main.go

API Endpoints:
--------------
Method   | Endpoint                 | Description
---------|--------------------------|---------------------------------
GET      | /certificates/:id        | Get certificate by ID
POST     | /certificates            | Create a new certificate
GET      | /certificates            | Get all certificates
PUT      | /certificates/:id        | Update a certificate
POST     | /certificates/upload     | Upload certificates via CSV

Uploading Certificates via CSV:
-------------------------------
To upload a CSV file, send a POST request to `/certificates/upload` with a file.  
The CSV format should be:

Name,Course,IssuedTo,IssueDate,ExpiryDate,Issuer  
John Doe,Go Programming,John Doe,2024-01-01,2025-01-01,Go Academy  

-----------------------------------
Happy coding! 🚀
