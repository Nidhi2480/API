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
    const studentInfoDiv = document.getElementById('mobileInfo');
    const studentInfoHTML = `
    <div class="row">
    <div class="col-md-4">
        <a href="viewmobile.html?id=${mobile.id}">
            <img src="${mobile.image}" class="img-fluid rounded" alt="mobile Image">
        </a>
    </div>
    <div class="col-md-8">
        <div class="mobile-details border border-secondary rounded p-4 ">
            <p><strong>Name:</strong> ${mobile.name}</p></div><br>
            <div class="mobile-details border border-secondary rounded p-4 ">
            <p><strong>Specifications:</strong> ${mobile.specs}</p><br></div><br>
            <div class="mobile-details border border-secondary rounded p-4 ">
            <p><strong>Rs.</strong> ${mobile.price}/-</p>
            </div>
    </div>
</div>
`;


    studentInfoDiv.innerHTML = studentInfoHTML;
}