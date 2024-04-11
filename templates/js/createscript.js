document.getElementById("mobileDetailForm").addEventListener("submit", function(event) {
    event.preventDefault();
    const formData = new FormData(this);
    fetch("http://localhost:8080/addmobile", {
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
        alert("Mobile detail added successfully!");
        console.log("Server response:", data);
        window.location.assign('listview.html');
        
    })
    .catch(error => {
        console.error('There was a problem with the fetch operation:', error);
        alert("Failed to add mobile detail. Please try again.");
    });
});