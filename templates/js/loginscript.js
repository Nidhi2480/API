document.getElementById("loginForm").addEventListener("submit", function(event) {
    event.preventDefault(); 
    const formData = new FormData(this); 
    fetch("http://localhost:8080/login", {
        method: "POST",
        body: formData
    })
    .then(response => {
        if (!response.ok) {
            window.alert('invalid Username or Password');
            window.location.assign('login.html');
        }
        return response.json();
    })
    .then(data => {
        console.log(data)
        if (data.token){
            console.log(data.token)
            sessionStorage.setItem('sessionID', data.token);
            sessionStorage.setItem('sessionROLE', data.role);
            window.alert("Login successful!");
            window.location.assign('listview.html');
        } else {
            msg=data.message;
            window.alert(msg);
        }
    })
    .catch(error => {
        console.error('There was a problem with the fetch operation:', error);
        window.alert("An error occurred. Please try again later.");
    });
});
