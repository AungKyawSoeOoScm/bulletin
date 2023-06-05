var constraints = {
  username: {
    presence: {
      allowEmpty: false,
      message: "^Username is required",
    },
    length: {
      minimum: 4,
      maximum: 20,
      tooShort: "^Username is too short (minimum %{count} characters)",
      tooLong: "^Username is too long (maximum %{count} characters)",
    },
  },
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
var constraints2 = {
  username: {
    presence: {
      allowEmpty: false,
      message: "^Username is required",
    },
    length: {
      minimum: 4,
      maximum: 20,
      tooShort: "^Username is too short (minimum %{count} characters)",
      tooLong: "^Username is too long (maximum %{count} characters)",
    },
  },
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
var usernameInput = document.getElementById("username");
var emailInput = document.getElementById("email");
var passwordInput = document.getElementById("password");
var confirmInput = document.getElementById("confirmPassword");
var photoInput = document.getElementById("photo");
var createUserBtn = document.getElementById("createUserBtn");
var updateUserBtn = document.getElementById("updateUserBtn");

var nameError = document.getElementById("userError");
var emailError = document.getElementById("emailError");
var passwordError = document.getElementById("passwordError");
var cpasswordError = document.getElementById("confirmPasswordError");
profileError = document.getElementById("profileError");

emailInput.addEventListener("input", validateForm);
usernameInput.addEventListener("input", validateForm);
if (passwordInput) {
  passwordInput.addEventListener("input", validateForm);
}
if (confirmInput) {
  confirmInput.addEventListener("input", validateForm);
}

photoInput.addEventListener("input", validateForm);

function validateForm() {
  if (passwordInput && confirmInput) {
    var validationResult1 = validate(
      {
        username: usernameInput.value,
        email: emailInput.value,
        password: passwordInput.value,
        confirmPassword: confirmInput.value,
        photo: photoInput.files[0],
      },
      constraints
    );
  } else {
    var validationResult2 = validate(
      {
        username: usernameInput.value,
        email: emailInput.value,
        photo: photoInput.files[0],
      },
      constraints2
    );
  }

  // Clear previous error messages
  if (emailError) {
    emailError.textContent = "";
  }
  if (nameError) {
    nameError.textContent = "";
  }
  if (passwordError) {
    passwordError.textContent = "";
  }
  if (cpasswordError) {
    cpasswordError.textContent = "";
  }
  if (profileError) {
    profileError.textContent = "";
  }

  if (validationResult1) {
    var userErrors = validationResult1.username;
    var emailErrors = validationResult1.email;

    var passwordErrors = validationResult1.password;

    var confirmPasswordErrors = validationResult1.confirmPassword;

    var photoErrors = validationResult1.photo;

    if (userErrors && userErrors.length > 0) {
      nameError.textContent = userErrors[0];
    } else if (emailErrors && emailErrors.length > 0) {
      emailError.textContent = emailErrors[0];
    } else if (passwordInput && passwordErrors && passwordErrors.length > 0) {
      passwordError.textContent = passwordErrors[0];
    } else if (
      confirmInput &&
      confirmPasswordErrors &&
      confirmPasswordErrors.length > 0
    ) {
      cpasswordError.textContent = confirmPasswordErrors[0];
    } else if (photoErrors && photoErrors.length > 0) {
      profileError.textContent = photoErrors[0];
    }
  } else if (validationResult2) {
    var userErrors = validationResult2.username;
    var emailErrors = validationResult2.email;

    var photoErrors = validationResult2.photo;

    if (userErrors && userErrors.length > 0) {
      nameError.textContent = userErrors[0];
    } else if (emailErrors && emailErrors.length > 0) {
      emailError.textContent = emailErrors[0];
    } else if (photoErrors && photoErrors.length > 0) {
      profileError.textContent = photoErrors[0];
    }
  } else {
    nameError.textContent = "";
    emailError.textContent = "";
    if (passwordInput && confirmInput) {
      passwordError.textContent = "";

      cpasswordError.textContent = "";
    }

    profileError.textContent = "";
  }

  var selectedFile = photoInput.files[0];
  var isValidFile = validateFile(selectedFile);

  if (!isValidFile) {
    profileError.textContent = "Invalid file format";
  }

  if (createUserBtn) {
    createUserBtn.disabled = !!validationResult1 || !isValidFile;
  }
  if (updateUserBtn) {
    if (passwordInput && confirmInput) {
      updateUserBtn.disabled = !!validationResult1 || !isValidFile;
    } else {
      updateUserBtn.disabled = !!validationResult2 || !isValidFile;
    }
  }
}

function validateFile(file) {
  if (!file) {
    return true;
  }

  var allowedExtensions = ["jpeg", "jpg", "png", "gif"];
  var fileExtension = file.name.split(".").pop().toLowerCase();
  console.log(fileExtension);
  if (!allowedExtensions.includes(fileExtension)) {
    return false;
  }
  return true;
}
