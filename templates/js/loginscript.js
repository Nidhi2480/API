document.getElementById("loginForm").addEventListener("submit", function(event) {
    event.preventDefault(); 
    const formData = new FormData(this); 
    fetch("http://localhost:8080/login", {
        method: "POST",
        body: formData
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        return response.json();
    })
    .then(data => {
        if (data.role === "admin") {
            sessionStorage.setItem('sessionID', data.role);
            window.alert("Login successful!");
            window.location.assign('listview.html');
        } else if (data.role === "user") {
            window.alert("Login successful!");
            sessionStorage.setItem('sessionID', data.role);
            window.location.assign('listview.html');
        } else {
            window.alert("Username or password incorrect");
        }
    })
    .catch(error => {
        console.error('There was a problem with the fetch operation:', error);
        window.alert("An error occurred. Please try again later.");
    });
});