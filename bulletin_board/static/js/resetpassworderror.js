// Define the validation rules
var constraints = {
  password: {
    presence: {
      allowEmpty: false,
      message: "^Password is empty",
    },
    length: {
      minimum: 6,
      maximum: 20,
      tooShort: "^Password is too short (minimum %{count} characters)",
      tooLong: "^Password is too long (maximum %{count} characters)",
    },
  },
  confirmPassword: {
    presence: {
      allowEmpty: false,
      message: "^Confirm password is empty",
    },
    equality: {
      attribute: "password",
      message: "^Password and password confirmation do not match",
    },
  },
};

var passwordInput = document.getElementById("password");
var confirmInput = document.getElementById("cpassword");
var resetPasswordBtn = document.getElementById("resetPasswordBtn");

var passwordError = document.getElementById("passwordError");
var confirmpasswordError = document.getElementById("confirmPasswordError");

// Attach input event listeners to the email and password inputs

passwordInput.addEventListener("input", validateForm);
confirmInput.addEventListener("input", validateForm);

function validateForm() {
  var validationResult = validate(
    {
      password: passwordInput.value,
      confirmPassword: confirmInput.value,
    },
    constraints
  );

  // Clear previous error messages
  if (passwordError) {
    passwordError.textContent = "";
  }
  if (confirmpasswordError) {
    confirmpasswordError.textContent = "";
  }

  if (validationResult) {
    var passwordErrors = validationResult.password;
    var confirmPasswordErrors = validationResult.confirmPassword;
    if (passwordErrors && passwordErrors.length > 0) {
      passwordError.textContent = passwordErrors[0];
    } else if (confirmPasswordErrors && confirmPasswordErrors.length > 0) {
      confirmpasswordError.textContent = confirmPasswordErrors[0];
    }
  } else {
    // Clear the error message when the password and confirm password match
    document.getElementById("confirmPasswordError").textContent = "";
  }

  resetPasswordBtn.disabled = !!validationResult;
}
