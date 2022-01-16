<?php

session_start();
require "db.php";

// La comprobacion de datos solo se hara si se ha enviado un
// formulario desde la pagina correspondiente.
if (isset($_POST['env'])) {
  $nomb = $_POST['nombre'];
  $rol = $_POST['rol'];
  $email = $_POST['correo'];
  $fecha_nac = $_POST['fecha_nac'];
  $pass = $_POST['pass'];
  $pass2 = $_POST['pass_conf'];

  if ($pass == $pass2) {
      $pwd = password_hash($pass, PASSWORD_BCRYPT);

      $query = "INSERT INTO cuenta(rol_usm, correo, contrasenna, nombre, fecha_nacimiento) VALUES('$rol', '$email', '$pwd', '$nomb', '$fecha_nac')";
      $res = mysqli_query($conn, $query);

      if ($res) {
        $_SESSION['mensaje'] = "Cuenta creada con exito. Inicia sesion ahora.";
          header('Location: /sansapp/php/login.php');
      } else {
        $_SESSION['mensaje'] = "El correo o el rol ya estan en uso.";
          header('Location: /sansapp/php/signup.php');
      }
  } else {
      $_SESSION['mensaje'] = "Las contraseñas no coinciden.";
      header('Location: /sansapp/php/signup.php');
  }
} else {
  $_SESSION['mensaje'] = "Usa el formulario por favor.";
  header('Location: /sansapp/php/signup.php');
}
