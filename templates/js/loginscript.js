document.getElementById("loginForm").addEventListener("submit", function(event) {
    event.preventDefault();

    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;
    const formData = new FormData();
    formData.append('username', username);
    formData.append('password', password);

    fetch("http://localhost:8080/login", {
        method: "POST",
        body: formData
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        return response.text();
    })
    .then(data => {
        if (data.includes("successful")) {
            window.alert("Login successful!");
            window.location.assign('listview.html');
        } else {
            window.alert(data);
        }
    })
    .catch(error => {
        console.error('There was a problem with the fetch operation:', error);
    });
});
