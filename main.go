package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/message", messageHandler)
	http.HandleFunc("/submit_contact", submitContactHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
	fmt.Println("Starting server on 8080")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `
<!DOCTYPE html>
<html>

<head>
    <title>Solar Rayz Got Shade</title>
    <link href="https://cdn.jsdelivr.net/npm/tailwindcss@2.2.19/dist/tailwind.min.css" rel="stylesheet">
    <script src="https://unpkg.com/htmx.org@1.9.6" integrity="sha384-FhXw7b6AlE/jyjlZH5iHa/tTe9EpJ1Y55RjcgPbjeWMskSxZt1v9qkxLJWNJaGni" crossorigin="anonymous"></script>
</head>

<body class="bg-gray-200 min-h-screen flex flex-col items-center p-4 space-y-8">
    <!-- Hero Image Section -->
    <div class="relative h-96 w-full max-w-4xl mb-8 bg-cover bg-center" style="background-image: url('static/solar-panel.jpg')">
        <div class="absolute inset-0 bg-black opacity-50"></div>
        <div class="absolute inset-0 flex flex-col justify-center items-center p-8 space-y-4">
            <h1 class="text-4xl font-bold text-white">Solar Rays Solar Panels</h1>
            <p class="text-lg text-center text-white">
                Harness the power of the sun with Solar Rays! Our high-quality solar panels are designed for maximum efficiency and longevity. Join the green energy revolution today and reduce your carbon footprint while saving on energy bills.
            </p>
        </div>
    </div>

    <!-- Contact Us Section -->
    <div class="bg-white p-8 w-full max-w-xl rounded shadow-md space-y-4">
        <h2 class="text-2xl font-bold">Contact Us</h2>
        <p>Got questions about our solar panels? Reach out to us below!</p>
        <form action="/submit_contact" method="POST">
            <div class="space-y-4">
                <!-- Name field -->
                <div>
                    <label for="name" class="block text-sm font-medium text-gray-600">Full Name</label>
                    <input type="text" id="name" name="name" required class="mt-1 p-2 w-full border rounded-md">
                </div>

                <!-- Email field -->
                <div>
                    <label for="email" class="block text-sm font-medium text-gray-600">Email Address</label>
                    <input type="email" id="email" name="email" required class="mt-1 p-2 w-full border rounded-md">
                </div>

                <!-- Phone number field -->
                <div>
                    <label for="phone" class="block text-sm font-medium text-gray-600">Phone Number</label>
                    <input type="tel" id="phone" name="phone" class="mt-1 p-2 w-full border rounded-md">
                </div>

                <!-- Website Field -->
                <div>
                    <label for="website" class="block text-sm font-medium text-gray-600">Your Company Website</label>
                    <input type="text" id="website" name="website" required class="mt-1 p-2 w-full border rounded-md">
                </div>

                <!-- Message textarea -->
                <div>
                    <label for="message" class="block text-sm font-medium text-gray-600">Message</label>
                    <textarea id="message" name="message" rows="4" required class="mt-1 p-2 w-full border rounded-md"></textarea>
                </div>

                <!-- Submit button -->
                <div>
                    <button type="submit" class="w-full bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600 focus:outline-none focus:border-blue-700 focus:ring focus:ring-blue-200">Submit</button>
                </div>
            </div>
        </form>
    </div>

    <!-- Chat Bot Section -->
    <div class="bg-white p-8 w-full max-w-xl rounded shadow-md space-y-4">
        <h1 class="text-2xl">Chat with our Bot</h1>
        <div id="chatBox" class="border p-4 h-56 overflow-y-scroll">
            <div class="mb-2"><strong>Bot:</strong> Hello! Type a message to start chatting.</div>
        </div>
        <input type="text" id="userMessage" placeholder="Type your message..." class="border p-2 w-full rounded" hx-get="/message" hx-params="message: this.value" hx-trigger="keyup changed delay:500ms" hx-target="#chatBox" hx-swap="beforeend" hx-postdata="this.value=''">
    </div>

    <script>
        document.getElementById("userMessage").addEventListener("htmx:afterOnLoad", function() {
            let chatBox = document.getElementById("chatBox");
            chatBox.scrollTop = chatBox.scrollHeight;
        });
    </script>
</body>

</html>



    `)
}

func submitContactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	// Extract form values
	name := r.FormValue("name")
	email := r.FormValue("email")
	phone := r.FormValue("phone")
	message := r.FormValue("message")

	// For demonstration purposes, just print out the received form data.
	// In a real-world scenario, you might save this to a database, send an email, etc.
	fmt.Printf("Received contact form:\nName: %s\nEmail: %s\nPhone: %s\nMessage: %s\n", name, email, phone, message)

	// Send a response back to the user
	fmt.Fprintf(w, "Thank you for contacting us, %s! We have received your message and will get back to you shortly.", name)
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	// Simulate a simple bot response based on received message
	message := r.URL.Query().Get("message")

	botResponse := ""
	switch strings.ToLower(message) {
	case "hello":
		botResponse = "Hello! How can I assist you today?"
	case "how are you":
		botResponse = "I'm just a bot, so I don't have feelings, but thanks for asking! How can I help?"
	default:
		botResponse = "I'm not sure how to answer that. Please ask another question."
	}

	fmt.Fprint(w, botResponse)
}
