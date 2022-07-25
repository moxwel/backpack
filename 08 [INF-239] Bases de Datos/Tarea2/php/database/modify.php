<?php

session_start();
require "db.php";

$id_cuenta = $_SESSION['id'];

// La comprobacion de datos solo se hara si se ha enviado un
// formulario desde la pagina correspondiente.
if (isset($_POST['env'])) {
  $nomb = $_POST['nombre'];
  $email = $_POST['correo'];
  $telefono = $_POST['telef'];
  $pass = $_POST['pass'];
  $pass2 = $_POST['pass_conf'];

  if ($pass == $pass2) {
      $pwd = password_hash($pass, PASSWORD_BCRYPT);

      $query = "UPDATE cuenta SET nombre='$nomb', correo='$email', telefono='$telefono', contrasenna='$pwd' WHERE id_cuenta='$id_cuenta'";
      $res = mysqli_query($conn, $query);

      if ($res) {
        session_destroy();
        session_start();
        $_SESSION['mensaje'] = "Cuenta modificada con exito. Vuelve a iniciar sesion.";
        header('Location: /sansapp/php/login.php');
      } else {
        $_SESSION['mensaje'] = "El correo ya esta en uso.";
        header('Location: /sansapp/php/modificar.php');
      }
  } else {
      $_SESSION['mensaje'] = "Las contraseñas no coinciden.";
      header('Location: /sansapp/php/modificar.php');
  }
} else {
  $_SESSION['mensaje'] = "Usa el formulario por favor.";
  header('Location: /sansapp/php/signup.php');
}
