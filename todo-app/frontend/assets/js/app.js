const API_URL = "http://localhost:8080/api/todos";

const list = document.getElementById("todo-list");

const form = document.getElementById("todo-form");
const fileInput = document.getElementById("todo-file");
const input = document.getElementById("todo-input");

// Fetch & render todos
async function fetchTodos() {
    const res = await fetch(API_URL);
    const todos = await res.json();

    list.innerHTML = "";

    todos.forEach(todo => {
        console.log("todo", todo);

        const li = document.createElement("li");
        li.innerHTML = `
      <span>${todo.title ?? "-"}</span>
      <button class="delete-btn" onclick="deleteTodo(${todo.id})">Delete</button>
    `;
        list.appendChild(li);
    });
}

// Add new todo
form.addEventListener("submit", async (e) => {
    e.preventDefault();

    const title = input.value.trim();
    if (!title) return;

    let pdfBase64 = "";

    if (fileInput.files.length > 0) {
        const file = fileInput.files[0];
        pdfBase64 = await toBase64(file);
    }

    await fetch(API_URL, {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            title: title,
            file: pdfBase64
        })
    });

    input.value = "";
    fileInput.value = "";
    fetchTodos();
});

// Delete todo
async function deleteTodo(id) {
    await fetch(`${API_URL}/${id}`, {
        method: "DELETE"
    });

    fetchTodos();
}

// to base 64
function toBase64(file) {
    return new Promise((resolve, reject) => {
        const reader = new FileReader();

        reader.readAsDataURL(file); // base64
        reader.onload = () => {
            // hasilnya: data:application/pdf;base64,xxxx
            resolve(reader.result);
        };
        reader.onerror = error => reject(error);
    });
}


// Initial load
fetchTodos();
