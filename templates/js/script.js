document.addEventListener("DOMContentLoaded", function() {
    var currentURL = window.location.href;
var urlParts = currentURL.split('?');
if (urlParts.length > 1) {
    var queryString = urlParts[1];
    var params = queryString.split('&');
    var id;
    params.forEach(function(param) {
        var paramParts = param.split('=');
        if (paramParts[0] === 'id') {
            id = paramParts[1];
        }
    });
    if (!id) {
        console.error("Missing 'id' parameter");
    } else {
    var secondAPIURL = "http://localhost:8080/getmobile/" + id;
    fetch(secondAPIURL) 
    .then(response => {
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        return response.json();
    })
    .then(data => {
        displayStudentInfo(data);
    })
    .catch(error => {
        console.error('There was a problem with the fetch operation:', error);
    });
}} else {
    console.error("URL does not contain a query string");
}}
);

function displayStudentInfo(student) {
    document.getElementById('head').innerHTML = student.name;
    const studentInfoDiv = document.getElementById('studentInfo');
    const studentInfoHTML = `
    <div class="row">
    <div class="col-md-4">
        <a href="file:///home/nidhinsajeev/Desktop/mobileapp/templates/student.html?id=${student.id}">
            <img src="${student.image}" class="img-fluid rounded" alt="Student Image">
        </a>
    </div>
    <div class="col-md-8">
        <div class="student-details">
            <p><strong>ID:</strong> ${student.id}</p>
            <p><strong>Name:</strong> ${student.name}</p>
            <p><strong>Specifications:</strong> ${student.specs}</p>
            <p><strong>Price:</strong> ${student.price}</p>
        </div>
    </div>
</div>
`;


    studentInfoDiv.innerHTML = studentInfoHTML;
}