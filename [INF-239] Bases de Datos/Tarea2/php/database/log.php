<?php
session_start();
require "db.php";

// La comprobacion de datos solo se hara si se ha enviado un
// formulario desde la pagina correspondiente.
if (isset($_POST['env'])) {
  $email = $_POST['correo'];
  $pass = $_POST['pass'];

  // Extraccion de datos desde db
  $query = "SELECT * FROM cuenta WHERE correo='$email'";
  $res = mysqli_query($conn, $query);
  $fila = mysqli_fetch_array($res);

  // Si no hay un email registrado con el
  // correo dado, entonces devuelve un error.
  if ($fila) {
    $nomb = $fila['nombre'];
    $rol = $fila['rol_usm'];
    $id = $fila['id_cuenta'];
    $fecnam = $fila['fecha_nacimiento'];
    $tel = $fila['telefono'];
    $pwd = $fila['contrasenna'];

    if (password_verify($pass, $pwd)) {
      $_SESSION['logged'] = true;
      $_SESSION['email'] = $email;
      $_SESSION['id'] = $id;
      $_SESSION['nombre'] = $nomb;
      $_SESSION['rol_usm'] = $rol;
      $_SESSION['fech_nac'] = $fecnam;
      $_SESSION['telefono'] = $tel;
      header('Location: /sansapp/php/home.php');
    } else {
      $_SESSION['mensaje'] = "Credenciales incorrectas.";
      header('Location: /sansapp/php/login.php');
    }
  } else {
    $_SESSION['mensaje'] = "Credenciales incorrectas.";
    header('Location: /sansapp/php/login.php');
  }
} else {
  $_SESSION['mensaje'] = "Usa el formulario por favor.";
  header('Location: /sansapp/php/login.php');
}
