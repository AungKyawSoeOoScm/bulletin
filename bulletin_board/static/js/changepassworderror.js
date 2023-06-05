// Define the validation rules
var constraints = {
  currentPassword: {
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

var currentpasswordInput = document.getElementById("cpassword");
var passwordInput = document.getElementById("password");
var confirmInput = document.getElementById("ncpassword");
var changePasswordBtn = document.getElementById("changePasswordBtn");

var currentpasswordError = document.getElementById("currentPasswordError");
var passwordError = document.getElementById("passwordError");
var confirmpasswordError = document.getElementById("confirmPasswordError");

// Attach input event listeners to the email and password inputs

currentpasswordInput.addEventListener("input", validateForm);
passwordInput.addEventListener("input", validateForm);
confirmInput.addEventListener("input", validateForm);

function validateForm() {
  var validationResult = validate(
    {
      currentPassword: currentpasswordInput.value,
      password: passwordInput.value,
      confirmPassword: confirmInput.value,
    },
    constraints
  );

  // Clear previous error messages
  if (currentpasswordError) {
    currentpasswordError.textContent = "";
  }
  if (passwordError) {
    passwordError.textContent = "";
  }
  if (confirmpasswordError) {
    confirmpasswordError.textContent = "";
  }

  if (validationResult) {
    var currentPasswordErrors = validationResult.currentPassword;
    var passwordErrors = validationResult.password;
    var confirmPasswordErrors = validationResult.confirmPassword;
    if (currentPasswordErrors && currentPasswordErrors.length > 0) {
      currentpasswordError.textContent = currentPasswordErrors[0];
    } else if (passwordErrors && passwordErrors.length > 0) {
      passwordError.textContent = passwordErrors[0];
    } else if (confirmPasswordErrors && confirmPasswordErrors.length > 0) {
      confirmpasswordError.textContent = confirmPasswordErrors[0];
    }
  } else {
    // Clear the error message when the password and confirm password match
    document.getElementById("confirmPasswordError").textContent = "";
  }

  changePasswordBtn.disabled = !!validationResult;
}
