function showConfirmation() {
  document.getElementById("overlay").style.display = "block";
  document.getElementById("confirmTitle").textContent =
    document.getElementById("title").value;
  document.getElementById("confirmDescription").textContent =
    document.getElementById("description").value;
  document.getElementById("confirmStatus").textContent =
    document.getElementById("status").value;
}

function confirmCreate() {
  document.getElementById("createForm").submit();
}

function confirmUpdate() {
  document.getElementById("updateForm").submit();
}

function cancelCreate() {
  document.getElementById("overlay").style.display = "none";
}

var titleInput = document.getElementById("title");
titleInput.addEventListener("input", function () {
  var titleError = document.getElementById("titleError");
  if (titleInput.value.trim() !== "") {
    titleError.innerHTML = "";
  } else {
    titleError.innerHTML = "Title field is required";
  }
});

// Description input event listener
var descriptionInput = document.getElementById("description");
descriptionInput.addEventListener("input", function () {
  var descriptionError = document.getElementById("descriptionError");
  if (descriptionInput.value.trim() !== "") {
    descriptionError.innerHTML = "";
  } else {
    descriptionError.innerHTML = "Description field is required";
  }
});
