// Define the validation rules
document.addEventListener("DOMContentLoaded", function () {
  var constraints = {
    email: {
      presence: {
        allowEmpty: false,
        message: "^Email is required",
      },
      email: true,
      length: {
        maximum: 50,
        tooLong: "^Email is too long (maximum %{count} characters)",
      },
    },
  };

  // Get the email and password input elements
  var emailInput = document.getElementById("email");
  var forgetPasswordBtn = document.getElementById("forgetPasswordBtn");

  var emailError = document.getElementById("emailError");

  // Attach input event listeners to the email and password inputs
  emailInput.addEventListener("input", validateLoginForm);

  // Function to validate the form and enable/disable the login button
  function validateLoginForm() {
    // Perform validation
    var validationResult = validate(
      {
        email: emailInput.value,
      },
      constraints
    );

    // Clear previous error messages
    if (emailError) {
      emailError.textContent = "";
    }

    // Display validation errors, if any
    if (validationResult) {
      var emailErrors = validationResult.email;
      if (emailErrors && emailErrors.length > 0) {
        emailError.textContent = emailErrors[0];
        // document.getElementById('passwordError').style.display = 'none'; // Hide the password error message
      }
    }

    // Enable/disable the login button
    forgetPasswordBtn.disabled = !!validationResult;
  }
});
