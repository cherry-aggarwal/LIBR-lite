# LIBR-lite

What you can do:

- 📨 Submit a message
- 📤 Fetch messages by timestamp

Moderation is decided using **majority voting** (2 out of 3).


## ⚙️ Setup Instructions

###  Install PostgreSQL

###  Configure Environment Variables

Create a .env file in the root folder:

change .env_sample filename to .env 
set up your credentials in .env


###  Run the Server

You should see the server running at:
http://localhost:3000


## Request-Response

1)📨 POST / <br />
Send a message for moderation.\

Request Body:\
{
  "content": "This is a test message"
}

Sample Response:\
{
  "id": "generated-uuid",
  "timestamp": 1744219507,
  "status": "approved"
}


2)📥 GET /{timestamp}<br />
Fetch messages by a specific timestamp.\

Body:\
1744219507
