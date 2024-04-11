document.getElementById("loginForm").addEventListener("submit", function(event) {
<<<<<<< HEAD
    event.preventDefault(); 
    const formData = new FormData(this); 
=======
    event.preventDefault();

    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;
    const formData = new FormData();
    formData.append('username', username);
    formData.append('password', password);

>>>>>>> 5b80a44f5c6aa687719752b3d40f36b88def21aa
    fetch("http://localhost:8080/login", {
        method: "POST",
        body: formData
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
<<<<<<< HEAD
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
=======
        return response.text();
    })
    .then(data => {
        if (data.includes("successful")) {
            window.alert("Login successful!");
            window.location.assign('listview.html');
        } else {
            window.alert(data);
>>>>>>> 5b80a44f5c6aa687719752b3d40f36b88def21aa
        }
    })
    .catch(error => {
        console.error('There was a problem with the fetch operation:', error);
<<<<<<< HEAD
        window.alert("An error occurred. Please try again later.");
    });
});
=======
    });
});
>>>>>>> 5b80a44f5c6aa687719752b3d40f36b88def21aa
