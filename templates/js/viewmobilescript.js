document.addEventListener("DOMContentLoaded", function() {
    const urlParams = new URLSearchParams(window.location.search);
    const productId = urlParams.get('id');
    var secondAPIURL = "http://localhost:8080/getmobile/" + productId;
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
}
);

function displayStudentInfo(mobile) {
    document.getElementById('head').innerHTML = mobile.name;
    const studentInfoDiv = document.getElementById('studentInfo');
    const studentInfoHTML = `
    <div class="row">
    <div class="col-md-4">
        <a href="file:///home/nidhinsajeev/Desktop/mobileapp/templates/student.html?id=${mobile.id}">
            <img src="${mobile.image}" class="img-fluid rounded" alt="Student Image">
        </a>
    </div>
    <div class="col-md-8">
        <div class="student-details">
            <p><strong>Name:</strong> ${mobile.name}</p>
            <p><strong>Specifications:</strong> ${mobile.specs}</p><br><hr>
            <p><strong>Rs.</strong> ${mobile.price}</p>
        </div>
    </div>
</div>
`;


    studentInfoDiv.innerHTML = studentInfoHTML;
}